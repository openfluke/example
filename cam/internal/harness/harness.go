// Package harness is tiny shared glue for cam examples that *prove* features.
package harness

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/openfluke/welvet/core"
	"github.com/openfluke/welvet/layers/dense"
	"github.com/openfluke/welvet/layers/parallel"
	"github.com/openfluke/welvet/quant"
	"github.com/openfluke/welvet/weights"
)

// DenseTwin builds 2× Dense cams merged by combine.
func DenseTwin(dim, out int, combine parallel.CombineMode) (*parallel.Layer, error) {
	return parallel.Hemispheres(dim, out, 2, combine, core.ActivationLinear, core.DTypeFloat32, quant.FormatNone)
}

// DenseN builds n Dense cams.
func DenseN(dim, out, n int, combine parallel.CombineMode) (*parallel.Layer, error) {
	return parallel.Hemispheres(dim, out, n, combine, core.ActivationLinear, core.DTypeFloat32, quant.FormatNone)
}

// FromBranches wraps arbitrary same-width Ops as cams.
func FromBranches(dim int, combine parallel.CombineMode, branches ...any) (*parallel.Layer, error) {
	return parallel.HemispheresFrom(parallel.Config{
		Dim: dim, Branches: len(branches), Combine: combine, OutFeat: 0,
	}, branches, nil)
}

// ToyXY returns flat [1,dim] input and [1,out] target.
func ToyXY(dim, out int) (x, y *core.Tensor[float32]) {
	x = core.NewTensor[float32](1, dim)
	y = core.NewTensor[float32](1, out)
	for i := range x.Data {
		x.Data[i] = float32(i+1) * 0.05
	}
	for i := range y.Data {
		y.Data[i] = 0.5
	}
	return x, y
}

// TrainN runs TrainMSE steps; returns last loss.
func TrainN(para *parallel.Layer, x, y *core.Tensor[float32], n int, lr float64) (float64, error) {
	var last float64
	for i := 0; i < n; i++ {
		loss, err := parallel.TrainMSE(para, x, y, parallel.ModeNormalBP, lr)
		if err != nil {
			return last, err
		}
		last = loss
		para.Remember(x, y)
	}
	return last, nil
}

// TrainTowardPost builds a target matching Forward(post) shape.
func TrainTowardPost(para *parallel.Layer, x *core.Tensor[float32], n int, lr float64, fill float32) (float64, error) {
	_, post, err := parallel.Forward(para, x)
	if err != nil {
		return 0, err
	}
	y := core.NewTensor[float32](post.Shape...)
	for i := range y.Data {
		y.Data[i] = fill
	}
	return TrainN(para, x, y, n, lr)
}

// MustDense is dense.New or panic.
func MustDense(in, out int) *dense.Layer {
	d, err := dense.New(in, out, core.ActivationLinear, core.DTypeFloat32)
	if err != nil {
		panic(err)
	}
	return d
}

// DenseWeights copies cam i primary Dense weights (nil if not Dense).
func DenseWeights(para *parallel.Layer, i int) []float32 {
	d, ok := para.DenseBranch(i)
	if !ok || d == nil || d.Weights == nil {
		return nil
	}
	w, ok := d.Weights.MasterF32()
	if !ok || w == nil {
		return nil
	}
	return append([]float32(nil), w...)
}

// SetDenseWeights overwrites cam i Dense master weights.
func SetDenseWeights(para *parallel.Layer, i int, w []float32) error {
	d, ok := para.DenseBranch(i)
	if !ok {
		return fmt.Errorf("cam %d not dense", i)
	}
	return d.Weights.SetFromF32(w)
}

// DivergeCams forces cam0≠cam1 so sync/DNA/cos have a signal.
// cam0 ← identity-ish, cam1 ← zeros (or scaled noise).
func DivergeCams(para *parallel.Layer) error {
	d0, ok0 := para.DenseBranch(0)
	d1, ok1 := para.DenseBranch(1)
	if !ok0 || !ok1 {
		return fmt.Errorf("DivergeCams needs Dense cams")
	}
	w0, ok := d0.Weights.MasterF32()
	if !ok {
		return fmt.Errorf("no master f32")
	}
	rows, cols := d0.Weights.Rows, d0.Weights.Cols
	a := make([]float32, len(w0))
	b := make([]float32, len(w0))
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			i := r*cols + c
			if r == c%rows || (rows == 1 && c == 0) {
				a[i] = 1
			} else {
				a[i] = float32(r+c) * 0.01
			}
			b[i] = float32(c-r) * 0.07 // clearly different
		}
	}
	if err := d0.Weights.SetFromF32(a); err != nil {
		return err
	}
	return d1.Weights.SetFromF32(b)
}

// DivergeTri forces three Dense cams to distinct weights.
func DivergeTri(para *parallel.Layer) error {
	for i := 0; i < 3; i++ {
		d, ok := para.DenseBranch(i)
		if !ok {
			return fmt.Errorf("need dense cam %d", i)
		}
		w, ok := d.Weights.MasterF32()
		if !ok {
			return fmt.Errorf("master")
		}
		out := make([]float32, len(w))
		for j := range out {
			out[j] = float32(i+1)*0.2 + float32(j)*0.01
		}
		if err := d.Weights.SetFromF32(out); err != nil {
			return err
		}
	}
	return nil
}

// WeightEq reports max abs diff.
func WeightMaxDiff(a, b []float32) float64 {
	if len(a) != len(b) {
		return math.Inf(1)
	}
	var m float64
	for i := range a {
		d := math.Abs(float64(a[i] - b[i]))
		if d > m {
			m = d
		}
	}
	return m
}

// Cos01 is StoreCosine of cam0 vs cam1.
func Cos01(para *parallel.Layer) float64 {
	d0, ok0 := para.DenseBranch(0)
	d1, ok1 := para.DenseBranch(1)
	if ok0 && ok1 {
		c, err := weights.StoreCosine(d0.Weights, d1.Weights)
		if err == nil {
			return c
		}
	}
	m := para.RefreshMetrics()
	if len(m.BranchCosines) > 1 {
		return m.BranchCosines[1]
	}
	return 0
}

// CosIJ cosine between dense cams i and j.
func CosIJ(para *parallel.Layer, i, j int) float64 {
	di, oki := para.DenseBranch(i)
	dj, okj := para.DenseBranch(j)
	if !oki || !okj {
		return 0
	}
	c, err := weights.StoreCosine(di.Weights, dj.Weights)
	if err != nil {
		return 0
	}
	return c
}

// Banner prints a section header.
func Banner(title string) { fmt.Printf("\n=== %s ===\n", title) }

// OK prints a passing proof line.
func OK(name string, parts ...any) {
	var b strings.Builder
	b.WriteString("  PASS  ")
	b.WriteString(name)
	for _, p := range parts {
		b.WriteString(fmt.Sprintf("  %v", p))
	}
	fmt.Println(b.String())
}

// Require fails the process if cond is false.
func Require(name string, cond bool, detail string) {
	if cond {
		OK(name, detail)
		return
	}
	fmt.Printf("  FAIL  %s  %s\n", name, detail)
	os.Exit(1)
}

// Requiref formats detail.
func Requiref(name string, cond bool, format string, args ...any) {
	Require(name, cond, fmt.Sprintf(format, args...))
}

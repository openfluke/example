package main

import (
	"fmt"
	"math"

	"github.com/openfluke/example/cam/internal/harness"
	"github.com/openfluke/welvet/core"
	"github.com/openfluke/welvet/layers/dense"
	"github.com/openfluke/welvet/layers/parallel"
)

func main() {
	harness.Banner("02_combine — prove merge modes")
	x, _ := harness.ToyXY(8, 4)

	proveAvgAddWidth(x)
	proveConcatWidth(x)
	proveMaxRouting(x)
	proveSparseK(x)
	proveDisagree(x)
	proveFilter(x)
	fmt.Println("\nall Combine proofs passed")
}

func proveAvgAddWidth(x *core.Tensor[float32]) {
	for _, c := range []parallel.CombineMode{parallel.CombineAvg, parallel.CombineAdd} {
		para, _ := harness.DenseTwin(8, 4, c)
		_, post, err := parallel.Forward(para, x)
		if err != nil {
			panic(err)
		}
		harness.Requiref(string(c)+"_width", len(post.Shape) == 2 && post.Shape[1] == 4,
			"outShape=%v want [1 4]", post.Shape)
	}
}

func proveConcatWidth(x *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineConcat)
	_, post, err := parallel.Forward(para, x)
	if err != nil {
		panic(err)
	}
	harness.Requiref("concat_width", post.Shape[1] == 8, "outShape=%v want feat=8", post.Shape)
}

func proveMaxRouting(x *core.Tensor[float32]) {
	// Force cam0 >> cam1 on output dim 0 so max must pick cam0 for that unit.
	para, _ := harness.DenseTwin(8, 4, parallel.CombineMax)
	_ = harness.DivergeCams(para)
	// Make cam0 row0 all large positive projections
	w0 := harness.DenseWeights(para, 0)
	for i := range w0 {
		w0[i] = 2
	}
	w1 := harness.DenseWeights(para, 1)
	for i := range w1 {
		w1[i] = -2
	}
	_ = harness.SetDenseWeights(para, 0, w0)
	_ = harness.SetDenseWeights(para, 1, w1)

	_, post, err := parallel.Forward(para, x)
	if err != nil {
		panic(err)
	}
	// With Linear act, cam0 posts should dominate → combined ≈ cam0
	_, o0, _ := forwardBranch(para, 0, x)
	ok := true
	for j := range post.Data {
		if math.Abs(float64(post.Data[j]-o0.Data[j])) > 1e-4 {
			ok = false
		}
	}
	harness.Requiref("max_routes_to_cam0", ok, "combined should match cam0 when cam0≫cam1")
}

func proveSparseK(x *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineSparseK)
	para.Cfg.SparseK = 1
	w0 := harness.DenseWeights(para, 0)
	w1 := harness.DenseWeights(para, 1)
	for i := range w0 {
		w0[i] = 3
		w1[i] = 0.01
	}
	_ = harness.SetDenseWeights(para, 0, w0)
	_ = harness.SetDenseWeights(para, 1, w1)
	_, post, err := parallel.Forward(para, x)
	if err != nil {
		panic(err)
	}
	_, o0, _ := forwardBranch(para, 0, x)
	ok := true
	for j := range post.Data {
		if math.Abs(float64(post.Data[j]-o0.Data[j])) > 1e-3 {
			ok = false
		}
	}
	harness.Requiref("sparsek_keeps_strong", ok, "SparseK=1 should ≈ strongest cam (cam0)")
}

func proveDisagree(x *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineDisagree)
	para.Cfg.DisagreeBeta = 1
	_ = harness.DivergeCams(para)
	_, post, err := parallel.Forward(para, x)
	if err != nil {
		panic(err)
	}
	_, a, _ := forwardBranch(para, 0, x)
	_, b, _ := forwardBranch(para, 1, x)
	ok := true
	for j := range post.Data {
		mean := 0.5 * (float64(a.Data[j]) + float64(b.Data[j]))
		want := mean + 1.0*(float64(a.Data[j])-float64(b.Data[j])) // β=1
		if math.Abs(float64(post.Data[j])-want) > 1e-4 {
			ok = false
		}
	}
	harness.Requiref("disagree_formula", ok, "y = mean + β(a-b) holds")
}

func proveFilter(x *core.Tensor[float32]) {
	a := harness.MustDense(8, 4)
	b := harness.MustDense(8, 4)
	gate, err := dense.New(8, 2, core.ActivationLinear, core.DTypeFloat32)
	if err != nil {
		panic(err)
	}
	para, err := parallel.HemispheresFrom(parallel.Config{
		Dim: 8, OutFeat: 4, Branches: 2, Combine: parallel.CombineFilter,
	}, []any{a, b}, gate)
	if err != nil {
		panic(err)
	}
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeNormalBP)
	y := core.NewTensor[float32](1, 4)
	for i := range y.Data {
		y.Data[i] = 0.5
	}
	gw, _ := gate.Weights.MasterF32()
	before := append([]float32(nil), gw...)
	_, err = harness.TrainN(para, x, y, 20, 0.1)
	if err != nil {
		panic(err)
	}
	gw2, _ := gate.Weights.MasterF32()
	dGate := harness.WeightMaxDiff(before, gw2)
	harness.Requiref("filter_gate_trains", dGate > 1e-5, "MoE gate Δ=%.4g", dGate)
}

// forwardBranch re-forwards one Dense cam (avg path helper).
func forwardBranch(para *parallel.Layer, i int, x *core.Tensor[float32]) (pre, post *core.Tensor[float32], err error) {
	d, ok := para.DenseBranch(i)
	if !ok {
		return nil, nil, fmt.Errorf("not dense")
	}
	return dense.Forward(d, x)
}

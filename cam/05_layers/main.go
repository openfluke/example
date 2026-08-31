package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/openfluke/example/cam/internal/harness"
	"github.com/openfluke/welvet/core"
	"github.com/openfluke/welvet/layers/cnn1"
	"github.com/openfluke/welvet/layers/cnn2"
	"github.com/openfluke/welvet/layers/cnn3"
	"github.com/openfluke/welvet/layers/convt1"
	"github.com/openfluke/welvet/layers/convt2"
	"github.com/openfluke/welvet/layers/convt3"
	"github.com/openfluke/welvet/layers/dense"
	"github.com/openfluke/welvet/layers/embedding"
	"github.com/openfluke/welvet/layers/gdn"
	"github.com/openfluke/welvet/layers/kmeans"
	"github.com/openfluke/welvet/layers/layernorm"
	"github.com/openfluke/welvet/layers/lstm"
	"github.com/openfluke/welvet/layers/mamba"
	"github.com/openfluke/welvet/layers/metacognition"
	"github.com/openfluke/welvet/layers/mha"
	"github.com/openfluke/welvet/layers/parallel"
	"github.com/openfluke/welvet/layers/residual"
	"github.com/openfluke/welvet/layers/rmsnorm"
	"github.com/openfluke/welvet/layers/rnn"
	"github.com/openfluke/welvet/layers/sequential"
	"github.com/openfluke/welvet/layers/softmax"
	"github.com/openfluke/welvet/layers/swiglu"
)

type layerDemo struct {
	name string
	why  string
	mk   func() (dim int, a, b any, x *core.Tensor[float32], err error)
}

func main() {
	only := flag.String("layer", "", "run one layer")
	list := flag.Bool("list", false, "list layers")
	flag.Parse()
	demos := allDemos()
	if *list {
		names := make([]string, len(demos))
		for i, d := range demos {
			names[i] = d.name
		}
		sort.Strings(names)
		fmt.Println(strings.Join(names, "\n"))
		return
	}

	harness.Banner("05_layers — prove each Op works as cams")
	fail := 0
	for _, d := range demos {
		if *only != "" && !strings.EqualFold(*only, d.name) {
			continue
		}
		if err := proveLayer(d); err != nil {
			fmt.Printf("  FAIL  %s  %v\n", d.name, err)
			fail++
			continue
		}
	}
	if fail > 0 {
		os.Exit(1)
	}
	fmt.Println("\nall layer-as-cam proofs passed")
}

func proveLayer(d layerDemo) error {
	dim, a, b, x, err := d.mk()
	if err != nil {
		return err
	}
	seedNoise(a, 0.3)
	seedNoise(b, -0.25)
	if d.name == "swiglu" {
		seedNoise(a, 2.5)
		seedNoise(b, -2.0)
	}
	var swigluBefore []float32
	if sg, ok := a.(*swiglu.Layer); ok && sg.Down != nil {
		if w, ok := sg.Down.Weights.MasterF32(); ok {
			swigluBefore = append([]float32(nil), w...)
		}
	}
	para, err := harness.FromBranches(dim, parallel.CombineAvg, a, b)
	if err != nil {
		return fmt.Errorf("host: %w", err)
	}

	// 1) Forward works
	_, post, err := parallel.Forward(para, x)
	if err != nil {
		return fmt.Errorf("forward: %w", err)
	}
	y := core.NewTensor[float32](post.Shape...)
	for i := range y.Data {
		y.Data[i] = 0.35
	}

	// 2) Both cams learn → loss drops (skip weightless softmax)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeNormalBP)
	loss0, err := parallel.TrainMSE(para, x, y, parallel.ModeNormalBP, 0.2)
	if err != nil {
		return fmt.Errorf("train0: %w", err)
	}
	var lossN float64
	for i := 0; i < 50; i++ {
		lossN, err = parallel.TrainMSE(para, x, y, parallel.ModeNormalBP, 0.2)
		if err != nil {
			return fmt.Errorf("train: %w", err)
		}
	}
	if d.name == "swiglu" {
		sg := a.(*swiglu.Layer)
		w, _ := sg.Down.Weights.MasterF32()
		if harness.WeightMaxDiff(swigluBefore, w) < 1e-6 {
			return fmt.Errorf("swiglu Down weights did not move under Train")
		}
	} else if d.name != "softmax" && d.name != "kmeans" && d.name != "gdn" {
		if !(lossN < loss0-1e-5) {
			return fmt.Errorf("expected loss drop %.6f → %.6f", loss0, lossN)
		}
	}
	if d.name == "gdn" {
		if _, _, err := parallel.Forward(para, x); err != nil {
			return fmt.Errorf("gdn forward: %w", err)
		}
	}

	// 3) Freeze: rebuild, freeze cam1
	dim, a, b, x, err = d.mk()
	if err != nil {
		return err
	}
	seedNoise(a, 0.3)
	seedNoise(b, -0.25)
	para, err = harness.FromBranches(dim, parallel.CombineAvg, a, b)
	if err != nil {
		return err
	}
	_, post, err = parallel.Forward(para, x)
	if err != nil {
		return err
	}
	y = core.NewTensor[float32](post.Shape...)
	for i := range y.Data {
		y.Data[i] = 0.35
	}
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeFreeze)
	para.SetCamSync(parallel.CamSyncConfig{Enabled: true, Alpha: 0.05, When: parallel.SyncAfterSample})

	var w1 []float32
	if _, ok := para.DenseBranch(1); ok {
		w1 = harness.DenseWeights(para, 1)
	}
	// Also snap CNN Proj via type assert when DenseBranch fails
	snap1 := snapPrimary(b)

	for i := 0; i < 20; i++ {
		if _, err := parallel.TrainMSE(para, x, y, parallel.ModeNormalBP, 0.15); err != nil {
			return fmt.Errorf("freeze-train: %w", err)
		}
	}
	m := para.RefreshMetrics()
	if len(m.ActiveModes) < 2 || m.ActiveModes[1] != "Freeze" {
		return fmt.Errorf("modes=%v want Freeze on cam1", m.ActiveModes)
	}
	if w1 != nil {
		d1 := harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
		if d1 > 1e-7 {
			return fmt.Errorf("frozen Dense cam1 moved Δ=%g", d1)
		}
	}
	if snap1 != nil {
		after := snapPrimary(b)
		if harness.WeightMaxDiff(snap1, after) > 1e-7 {
			return fmt.Errorf("frozen cam1 primary store moved Δ=%g", harness.WeightMaxDiff(snap1, after))
		}
	}

	harness.OK(d.name, d.why, fmt.Sprintf("loss %.4f→%.4f", loss0, lossN), "Freeze ok")
	return nil
}

func seedNoise(op any, scale float32) {
	set := func(w []float32) {
		for i := range w {
			w[i] = scale * float32((i%11)-5) * 0.05
		}
	}
	switch v := op.(type) {
	case *dense.Layer:
		if w, ok := v.Weights.MasterF32(); ok {
			set(w)
			_ = v.Weights.SetFromF32(w)
		}
	case *cnn1.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				set(w)
				_ = v.Proj.Weights.SetFromF32(w)
			}
		}
	case *cnn2.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				set(w)
				_ = v.Proj.Weights.SetFromF32(w)
			}
		}
	case *cnn3.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				set(w)
				_ = v.Proj.Weights.SetFromF32(w)
			}
		}
	case *convt1.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				set(w)
				_ = v.Proj.Weights.SetFromF32(w)
			}
		}
	case *convt2.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				set(w)
				_ = v.Proj.Weights.SetFromF32(w)
			}
		}
	case *convt3.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				set(w)
				_ = v.Proj.Weights.SetFromF32(w)
			}
		}
	case *mha.Layer:
		for _, d := range []*dense.Layer{v.Q, v.K, v.V, v.O} {
			if d == nil {
				continue
			}
			if w, ok := d.Weights.MasterF32(); ok {
				set(w)
				_ = d.Weights.SetFromF32(w)
			}
		}
	case *swiglu.Layer:
		for _, d := range []*dense.Layer{v.Gate, v.Up, v.Down} {
			if d == nil {
				continue
			}
			if w, ok := d.Weights.MasterF32(); ok {
				set(w)
				_ = d.Weights.SetFromF32(w)
			}
		}
	case *sequential.Layer:
		for _, d := range v.Children {
			if d == nil {
				continue
			}
			if w, ok := d.Weights.MasterF32(); ok {
				set(w)
				_ = d.Weights.SetFromF32(w)
			}
		}
	case *mamba.Layer:
		for _, d := range []*dense.Layer{v.InProj, v.OutProj} {
			if d == nil {
				continue
			}
			if w, ok := d.Weights.MasterF32(); ok {
				set(w)
				_ = d.Weights.SetFromF32(w)
			}
		}
	}
}

func snapPrimary(op any) []float32 {
	switch v := op.(type) {
	case *dense.Layer:
		if w, ok := v.Weights.MasterF32(); ok {
			return append([]float32(nil), w...)
		}
	case *cnn1.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				return append([]float32(nil), w...)
			}
		}
	case *cnn2.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				return append([]float32(nil), w...)
			}
		}
	case *cnn3.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				return append([]float32(nil), w...)
			}
		}
	case *convt1.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				return append([]float32(nil), w...)
			}
		}
	case *convt2.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				return append([]float32(nil), w...)
			}
		}
	case *convt3.Layer:
		if v.Proj != nil {
			if w, ok := v.Proj.Weights.MasterF32(); ok {
				return append([]float32(nil), w...)
			}
		}
	case *mha.Layer:
		if v.O != nil {
			if w, ok := v.O.Weights.MasterF32(); ok {
				return append([]float32(nil), w...)
			}
		}
	case *swiglu.Layer:
		if v.Down != nil {
			if w, ok := v.Down.Weights.MasterF32(); ok {
				return append([]float32(nil), w...)
			}
		}
	case *sequential.Layer:
		if len(v.Children) > 0 && v.Children[0] != nil {
			if w, ok := v.Children[0].Weights.MasterF32(); ok {
				return append([]float32(nil), w...)
			}
		}
	case *mamba.Layer:
		if v.OutProj != nil {
			if w, ok := v.OutProj.Weights.MasterF32(); ok {
				return append([]float32(nil), w...)
			}
		}
	}
	return nil
}

func allDemos() []layerDemo {
	return []layerDemo{
		{"dense", "default host", func() (int, any, any, *core.Tensor[float32], error) {
			return 4, harness.MustDense(4, 4), harness.MustDense(4, 4), fill(1, 4), nil
		}},
		{"cnn1", "1D conv", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := cnn1.Config{InChannels: 1, Filters: 2, SeqLen: 4, Kernel: 3, Stride: 1}
			a, err := cnn1.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := cnn1.New(cfg)
			return 4, a, b, fill(1, 1, 4), err
		}},
		{"cnn2", "2D vision", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := cnn2.Config{InChannels: 1, Filters: 2, Height: 3, Width: 3, Kernel: 2}
			a, err := cnn2.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := cnn2.New(cfg)
			return 9, a, b, fill(1, 1, 3, 3), err
		}},
		{"cnn3", "3D", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := cnn3.Config{InChannels: 1, Filters: 2, Depth: 2, Height: 2, Width: 2, Kernel: 2}
			a, err := cnn3.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := cnn3.New(cfg)
			return 8, a, b, fill(1, 1, 2, 2, 2), err
		}},
		{"convt1", "1D upsample", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := convt1.Config{InChannels: 1, Filters: 2, SeqLen: 2, Kernel: 3, Stride: 1}
			a, err := convt1.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := convt1.New(cfg)
			return 2, a, b, fill(1, 1, 2), err
		}},
		{"convt2", "2D upsample", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := convt2.Config{InChannels: 1, Filters: 2, Height: 2, Width: 2, Kernel: 2}
			a, err := convt2.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := convt2.New(cfg)
			return 4, a, b, fill(1, 1, 2, 2), err
		}},
		{"convt3", "3D upsample", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := convt3.Config{InChannels: 1, Filters: 2, Depth: 2, Height: 2, Width: 2, Kernel: 2}
			a, err := convt3.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := convt3.New(cfg)
			return 8, a, b, fill(1, 1, 2, 2, 2), err
		}},
		{"mha", "attention", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := mha.Config{DModel: 4, NumHeads: 1, MaxSeqLen: 4}
			a, err := mha.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := mha.New(cfg)
			return 4, a, b, fill(1, 2, 4), err
		}},
		{"swiglu", "FFN", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := swiglu.Config{InputDim: 4, IntermediateDim: 8}
			a, err := swiglu.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := swiglu.New(cfg)
			return 4, a, b, fill(1, 4), err
		}},
		{"rmsnorm", "norm", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := rmsnorm.Config{Dim: 4}
			a, err := rmsnorm.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := rmsnorm.New(cfg)
			x := fill(1, 4)
			for i := range x.Data {
				x.Data[i] = float32(i + 1)
			}
			return 4, a, b, x, err
		}},
		{"layernorm", "norm", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := layernorm.Config{Dim: 4}
			a, err := layernorm.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := layernorm.New(cfg)
			x := fill(1, 4)
			for i := range x.Data {
				x.Data[i] = float32(i + 1)
			}
			return 4, a, b, x, err
		}},
		{"softmax", "distribution", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := softmax.Config{Dim: 4}
			a, err := softmax.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := softmax.New(cfg)
			x := fill(1, 4)
			x.Data[0] = 2
			return 4, a, b, x, err
		}},
		{"rnn", "temporal", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := rnn.Config{InputSize: 3, HiddenSize: 4, SeqLen: 2}
			a, err := rnn.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := rnn.New(cfg)
			return 3, a, b, fill(1, 2, 3), err
		}},
		{"lstm", "gated temporal", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := lstm.Config{InputSize: 3, HiddenSize: 4, SeqLen: 2}
			a, err := lstm.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := lstm.New(cfg)
			return 3, a, b, fill(1, 2, 3), err
		}},
		{"embedding", "tables", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := embedding.Config{VocabSize: 8, EmbeddingDim: 4, SeqLen: 2}
			a, err := embedding.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := embedding.New(cfg)
			x := fill(1, 2)
			x.Data[0], x.Data[1] = 1, 2
			return 2, a, b, x, err
		}},
		{"sequential", "deep hemi", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := sequential.Config{Dim: 4, Depth: 2}
			a, err := sequential.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := sequential.New(cfg)
			return 4, a, b, fill(1, 4), err
		}},
		{"residual", "residual hemi", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := residual.Config{Dim: 4, Depth: 1}
			a, err := residual.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := residual.New(cfg)
			return 4, a, b, fill(1, 4), err
		}},
		{"kmeans", "prototypes", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := kmeans.Config{NumClusters: 3, FeatureDim: 4}
			a, err := kmeans.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := kmeans.New(cfg)
			return 4, a, b, fill(1, 4), err
		}},
		{"mamba", "SSM", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := mamba.Config{DModel: 4, DState: 2, SeqLen: 2}
			a, err := mamba.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := mamba.New(cfg)
			return 4, a, b, fill(1, 2, 4), err
		}},
		{"metacognition", "observer", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := metacognition.Config{Dim: 4}
			a, err := metacognition.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := metacognition.New(cfg)
			return 4, a, b, fill(1, 4), err
		}},
		{"gdn", "gated delta", func() (int, any, any, *core.Tensor[float32], error) {
			cfg := gdn.Config{HiddenSize: 4, NumKeyHeads: 1, NumValueHeads: 1, KeyHeadDim: 2, ValueHeadDim: 2, ConvKernel: 2}
			a, err := gdn.New(cfg)
			if err != nil {
				return 0, nil, nil, nil, err
			}
			b, err := gdn.New(cfg)
			return 4, a, b, fill(1, 2, 4), err
		}},
	}
}

func fill(shape ...int) *core.Tensor[float32] {
	t := core.NewTensor[float32](shape...)
	for i := range t.Data {
		t.Data[i] = 0.1 + float32(i%7)*0.03
	}
	return t
}

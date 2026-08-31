package main

import (
	"fmt"

	"github.com/openfluke/example/cam/internal/harness"
	"github.com/openfluke/welvet/core"
	"github.com/openfluke/welvet/layers/dense"
	"github.com/openfluke/welvet/layers/parallel"
	"github.com/openfluke/welvet/quant"
)

func main() {
	harness.Banner("03_camsync — prove sync pulls cos up")
	x, y := harness.ToyXY(8, 4)

	proveAlphaPull(x, y)
	proveOneWay(x, y)
	proveGroups(x, y)
	proveCross(x, y)
	fmt.Println("\nall CamSync proofs passed")
}

func proveAlphaPull(x, y *core.Tensor[float32]) {
	// Train with NO sync → diverge further, measure cos.
	// Then SyncNow at α=1 → cos must jump toward 1.
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeTween) // different updates → more diverge
	_, _ = harness.TrainN(para, x, y, 15, 0.1)
	cosBefore := harness.Cos01(para)

	para.SetCamSync(parallel.CamSyncConfig{Enabled: true, Alpha: 1.0, When: parallel.SyncManual})
	if err := para.SyncNow(); err != nil {
		panic(err)
	}
	cosAfter := harness.Cos01(para)
	harness.Requiref("α=1_hard_sync", cosAfter > 0.999 && cosAfter > cosBefore+0.01,
		"cos %.4f → %.4f (hard sync should ≈1 and beat pre)", cosBefore, cosAfter)

	// Soft α: diverge again, soft pull should raise cos but not fully equalize weights.
	para2, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para2)
	cos0 := harness.Cos01(para2)
	para2.SetCamSync(parallel.CamSyncConfig{Enabled: true, Alpha: 0.25, When: parallel.SyncManual})
	_ = para2.SyncNow()
	cos1 := harness.Cos01(para2)
	diffStill := harness.WeightMaxDiff(harness.DenseWeights(para2, 0), harness.DenseWeights(para2, 1))
	harness.Requiref("α=0.25_soft_pull", cos1 > cos0 && diffStill > 1e-3,
		"cos %.4f→%.4f, still distinct ΔW=%.3g", cos0, cos1, diffStill)
}

func proveOneWay(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	w1Before := harness.DenseWeights(para, 1)
	cos0 := harness.Cos01(para)
	para.SetCamSync(parallel.CamSyncConfig{
		Enabled: true, Alpha: 1, When: parallel.SyncManual,
		BranchAlpha: []float64{1, 0}, // write cam0 only
	})
	_ = para.SyncNow()
	d1 := harness.WeightMaxDiff(w1Before, harness.DenseWeights(para, 1))
	cos1 := harness.Cos01(para)
	harness.Requiref("BranchAlpha_one_way", d1 < 1e-7 && cos1 > cos0,
		"cam1 unchanged Δ=%.4g; cos %.4f→%.4f (cam0 pulled to mean)", d1, cos0, cos1)
}

func proveGroups(x, y *core.Tensor[float32]) {
	tri, _ := harness.DenseN(8, 4, 3, parallel.CombineAvg)
	// Orthogonal-ish cams so cos02 starts low.
	for i := 0; i < 3; i++ {
		d, _ := tri.DenseBranch(i)
		w, _ := d.Weights.MasterF32()
		out := make([]float32, len(w))
		for j := range out {
			out[j] = 0
		}
		// one-hot-ish rows per cam
		cols := d.Weights.Cols
		for r := 0; r < d.Weights.Rows; r++ {
			out[r*cols+(i%cols)] = float32(1 + i)
		}
		_ = d.Weights.SetFromF32(out)
	}
	w2 := harness.DenseWeights(tri, 2)
	diff02Before := harness.WeightMaxDiff(harness.DenseWeights(tri, 0), harness.DenseWeights(tri, 2))
	tri.SetCamSync(parallel.CamSyncConfig{
		Enabled: true, Alpha: 1, When: parallel.SyncManual,
		Groups: [][]int{{0, 1}},
	})
	_ = tri.SyncNow()
	d2 := harness.WeightMaxDiff(w2, harness.DenseWeights(tri, 2))
	diff01 := harness.WeightMaxDiff(harness.DenseWeights(tri, 0), harness.DenseWeights(tri, 1))
	diff02 := harness.WeightMaxDiff(harness.DenseWeights(tri, 0), harness.DenseWeights(tri, 2))
	harness.Requiref("Groups_exclude_cam2",
		d2 < 1e-7 && diff01 < 1e-5 && diff02 > 0.5 && diff02Before > 0.5,
		"cam2 frozen Δ=%.4g; cam0≡cam1 Δ=%.4g; cam0≠cam2 Δ=%.3g (was %.3g)",
		d2, diff01, diff02, diff02Before)
}

func proveCross(x, y *core.Tensor[float32]) {
	const H = 8
	stem, _ := dense.NewConfigured[float32](8, H, core.ActivationLinear, core.DTypeFloat32, quant.FormatNone, nil)
	paraA, _ := parallel.Hemispheres(H, H, 2, parallel.CombineAdd, core.ActivationLinear, core.DTypeFloat32, quant.FormatNone)
	head, _ := dense.NewConfigured[float32](H, 4, core.ActivationLinear, core.DTypeFloat32, quant.FormatNone, nil)
	stack, err := parallel.Sandwich(stem, paraA, head)
	if err != nil {
		panic(err)
	}
	_ = harness.DivergeCams(paraA)
	cos0 := harness.Cos01(paraA)
	paraA.SetBranchModes(parallel.ModeNormalBP, parallel.ModeNormalBP)
	stack.SetCamSync(parallel.CamSyncConfig{
		Enabled: true, Alpha: 1.0, When: parallel.SyncManual,
		Cross: []parallel.SyncPair{{
			A: parallel.SyncEndpoint{StackIdx: 1, Branch: 0, Store: 0},
			B: parallel.SyncEndpoint{StackIdx: 1, Branch: 1, Store: 0},
		}},
	})
	if err := stack.SyncNow(); err != nil {
		panic(err)
	}
	cos1 := harness.Cos01(paraA)
	harness.Requiref("Cross_pair", cos1 > 0.999 && cos1 > cos0,
		"Cross hard-sync cos %.4f→%.4f", cos0, cos1)

	// Also one TrainStackMSE sample to prove training+path still works
	xx, yy := harness.ToyXY(8, 4)
	_, err = parallel.TrainStackMSE(stack, xx, yy, parallel.ModeNormalBP, 0.05)
	harness.Requiref("Cross_stack_trains", err == nil, "TrainStackMSE err=%v", err)
}

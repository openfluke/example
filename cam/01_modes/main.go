package main

import (
	"fmt"

	"github.com/openfluke/example/cam/internal/harness"
	"github.com/openfluke/welvet/core"
	"github.com/openfluke/welvet/layers/parallel"
)

func main() {
	harness.Banner("01_modes — prove BranchModes")
	x, y := harness.ToyXY(8, 4)

	proveBothBP(x, y)
	proveFreeze(x, y)
	proveShadow(x, y)
	proveAdversarial(x, y)
	proveMemory(x, y)
	proveTween(x, y)
	fmt.Println("\nall BranchMode proofs passed")
}

func proveBothBP(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeNormalBP)
	w0, w1 := harness.DenseWeights(para, 0), harness.DenseWeights(para, 1)
	_, _ = harness.TrainN(para, x, y, 20, 0.1)
	d0 := harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0))
	d1 := harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
	harness.Requiref("both_BP", d0 > 1e-4 && d1 > 1e-4, "both cams moved Δ=%.4g / %.4g", d0, d1)
}

func proveFreeze(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeFreeze)
	w0, w1 := harness.DenseWeights(para, 0), harness.DenseWeights(para, 1)
	_, _ = harness.TrainN(para, x, y, 20, 0.1)
	d0 := harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0))
	d1 := harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
	harness.Requiref("Freeze", d0 > 1e-4 && d1 < 1e-7, "cam0 moved Δ=%.4g, frozen cam1 Δ=%.4g (want ~0)", d0, d1)
}

func proveShadow(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeShadow)
	para.SetCamKit(parallel.CamKit{ShadowCoef: 1.0})
	teacherBefore := harness.DenseWeights(para, 1)
	studentBefore := harness.DenseWeights(para, 0)
	_, _ = harness.TrainN(para, x, y, 25, 0.08)
	dT := harness.WeightMaxDiff(teacherBefore, harness.DenseWeights(para, 1))
	dS := harness.WeightMaxDiff(studentBefore, harness.DenseWeights(para, 0))
	harness.Requiref("Shadow", dT < 1e-7 && dS > 1e-4, "teacher frozen Δ=%.4g, student moved Δ=%.4g", dT, dS)
}

func proveAdversarial(x, y *core.Tensor[float32]) {
	// Compare loss: BP∥BP should beat BP∥Adv on same init path.
	mk := func(modes ...parallel.TrainMode) float64 {
		para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
		// same diverge every time
		_ = harness.DivergeCams(para)
		para.SetBranchModes(modes...)
		loss, _ := harness.TrainN(para, x, y, 30, 0.08)
		return loss
	}
	lossBP := mk(parallel.ModeNormalBP, parallel.ModeNormalBP)
	lossAdv := mk(parallel.ModeNormalBP, parallel.ModeAdversarial)
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeAdversarial)
	w0, w1 := harness.DenseWeights(para, 0), harness.DenseWeights(para, 1)
	_, _ = harness.TrainN(para, x, y, 20, 0.08)
	d0 := harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0))
	d1 := harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
	harness.Requiref("Adversarial", d0 > 1e-4 && d1 > 1e-4 && lossAdv > lossBP*0.95,
		"both move Δ=%.3g/%.3g; Adv loss %.4f ≥ BP loss %.4f (fight)", d0, d1, lossAdv, lossBP)
}

func proveMemory(x, y *core.Tensor[float32]) {
	// High thresh → asleep (no cam1 update). Low thresh → awake (cam1 updates).
	asleep := func() float64 {
		para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
		_ = harness.DivergeCams(para)
		para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeMemory)
		para.SetCamKit(parallel.CamKit{SurpriseThresh: 1e9})
		w1 := harness.DenseWeights(para, 1)
		_, _ = harness.TrainN(para, x, y, 15, 0.1)
		return harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
	}
	awake := func() float64 {
		para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
		_ = harness.DivergeCams(para)
		para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeMemory)
		para.SetCamKit(parallel.CamKit{SurpriseThresh: 1e-12}) // always surprised
		w1 := harness.DenseWeights(para, 1)
		_, _ = harness.TrainN(para, x, y, 15, 0.1)
		return harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
	}
	dSleep, dWake := asleep(), awake()
	harness.Requiref("Memory", dSleep < 1e-7 && dWake > 1e-4,
		"asleep Δ=%.4g (~0), awake Δ=%.4g (>0)", dSleep, dWake)
}

func proveTween(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeTween)
	w1 := harness.DenseWeights(para, 1)
	_, _ = harness.TrainN(para, x, y, 20, 0.08)
	d1 := harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
	harness.Requiref("Tween", d1 > 1e-5, "Tween cam moved Δ=%.4g", d1)
}

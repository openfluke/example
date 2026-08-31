package main

import (
	"fmt"

	"github.com/openfluke/example/cam/internal/harness"
	"github.com/openfluke/welvet/core"
	"github.com/openfluke/welvet/layers/parallel"
)

func main() {
	harness.Banner("06_recipes — prove composed behaviors")
	x, y := harness.ToyXY(8, 4)

	proveTeacher(x, y)
	proveSleep(x, y)
	proveDebate(x, y)
	proveSurprise(x, y)
	proveDream(x, y)
	proveConcatCross()
	fmt.Println("\nall recipe proofs passed")
}

func proveTeacher(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeShadow)
	para.SetCamKit(parallel.CamKit{ShadowCoef: 1})
	para.SetCamSync(parallel.CamSyncConfig{
		Enabled: true, Alpha: 0.5, When: parallel.SyncAfterSample,
		BranchAlpha: []float64{1, 0},
	})
	t0 := harness.DenseWeights(para, 1)
	s0 := harness.DenseWeights(para, 0)
	_, _ = harness.TrainN(para, x, y, 20, 0.08)
	harness.Requiref("teacher_student",
		harness.WeightMaxDiff(t0, harness.DenseWeights(para, 1)) < 1e-7 &&
			harness.WeightMaxDiff(s0, harness.DenseWeights(para, 0)) > 1e-4,
		"teacher frozen, student moved")
}

func proveSleep(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetRotateSchedule([][]parallel.TrainMode{
		{parallel.ModeNormalBP, parallel.ModeFreeze},
		{parallel.ModeFreeze, parallel.ModeNormalBP},
	}, 5)
	w0, w1 := harness.DenseWeights(para, 0), harness.DenseWeights(para, 1)
	_, _ = harness.TrainN(para, x, y, 5, 0.1)
	d0, d1 := harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0)), harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
	harness.Requiref("sleep_slot0", d0 > 1e-4 && d1 < 1e-7, "only cam0 Δ=%.3g/%.3g", d0, d1)
	w0, w1 = harness.DenseWeights(para, 0), harness.DenseWeights(para, 1)
	_, _ = harness.TrainN(para, x, y, 5, 0.1)
	d0, d1 = harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0)), harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
	harness.Requiref("sleep_slot1", d0 < 1e-7 && d1 > 1e-4, "only cam1 Δ=%.3g/%.3g", d0, d1)
}

func proveDebate(x, y *core.Tensor[float32]) {
	lossBP := func() float64 {
		p, _ := harness.DenseTwin(8, 4, parallel.CombineDisagree)
		_ = harness.DivergeCams(p)
		p.Cfg.DisagreeBeta = 0.75
		p.SetBranchModes(parallel.ModeNormalBP, parallel.ModeNormalBP)
		l, _ := harness.TrainN(p, x, y, 25, 0.06)
		return l
	}()
	para, _ := harness.DenseTwin(8, 4, parallel.CombineDisagree)
	para.Cfg.DisagreeBeta = 0.75
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeAdversarial)
	w0, w1 := harness.DenseWeights(para, 0), harness.DenseWeights(para, 1)
	loss, _ := harness.TrainN(para, x, y, 25, 0.06)
	d0 := harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0))
	d1 := harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
	harness.Requiref("debate", d0 > 1e-4 && d1 > 1e-4 && loss >= lossBP*0.9,
		"both move; debate loss %.4f vs coop %.4f", loss, lossBP)
}

func proveSurprise(x, y *core.Tensor[float32]) {
	asleep := func() float64 {
		p, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
		_ = harness.DivergeCams(p)
		p.SetBranchModes(parallel.ModeNormalBP, parallel.ModeMemory)
		p.SetCamKit(parallel.CamKit{SurpriseThresh: 1e9})
		w := harness.DenseWeights(p, 1)
		_, _ = harness.TrainN(p, x, y, 12, 0.1)
		return harness.WeightMaxDiff(w, harness.DenseWeights(p, 1))
	}()
	awake := func() float64 {
		p, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
		_ = harness.DivergeCams(p)
		p.SetBranchModes(parallel.ModeNormalBP, parallel.ModeMemory)
		p.SetCamKit(parallel.CamKit{SurpriseThresh: 1e-12})
		w := harness.DenseWeights(p, 1)
		_, _ = harness.TrainN(p, x, y, 12, 0.1)
		return harness.WeightMaxDiff(w, harness.DenseWeights(p, 1))
	}()
	harness.Requiref("surprise_memory", asleep < 1e-7 && awake > 1e-4,
		"asleep Δ=%.4g awake Δ=%.4g", asleep, awake)
}

func proveDream(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeFreeze)
	para.SetCamKit(parallel.CamKit{Dream: &parallel.DreamBuffer{Cap: 64}})
	_, _ = harness.TrainN(para, x, y, 10, 0.08)
	w := harness.DenseWeights(para, 0)
	avg, err := para.DreamPulse(5, parallel.ModeNormalBP, 0.05)
	harness.Requiref("dream_consolidate", err == nil && avg > 0 &&
		harness.WeightMaxDiff(w, harness.DenseWeights(para, 0)) > 1e-5,
		"avg=%.4f weights moved on replay", avg)
}

func proveConcatCross() {
	dA := harness.MustDense(8, 3)
	dB := harness.MustDense(8, 5)
	mix, err := harness.FromBranches(8, parallel.CombineConcat, dA, dB)
	if err != nil {
		panic(err)
	}
	mix.SetBranchModes(parallel.ModeNormalBP, parallel.ModeTween)
	x, _ := harness.ToyXY(8, 8)
	_, post, err := parallel.Forward(mix, x)
	if err != nil {
		panic(err)
	}
	harness.Requiref("cross_modal_concat", post.Shape[1] == 8, "feat=%d want 3+5=8", post.Shape[1])
	y := core.NewTensor[float32](1, 8)
	for i := range y.Data {
		y.Data[i] = 0.4
	}
	wA, _ := dA.Weights.MasterF32()
	before := append([]float32(nil), wA...)
	_, _ = harness.TrainN(mix, x, y, 15, 0.08)
	wA2, _ := dA.Weights.MasterF32()
	harness.Requiref("cross_modal_trains", harness.WeightMaxDiff(before, wA2) > 1e-5, "cam0 moved under concat")
}

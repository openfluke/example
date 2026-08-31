package main

import (
	"fmt"

	"github.com/openfluke/example/cam/internal/harness"
	"github.com/openfluke/welvet/core"
	"github.com/openfluke/welvet/layers/parallel"
)

func main() {
	harness.Banner("04_kit — prove LRs / rotate / DNA / dream")
	x, y := harness.ToyXY(8, 4)

	proveBranchLR(x, y)
	proveRotate(x, y)
	proveDNA(x, y)
	proveDream(x, y)
	proveMetrics(x, y)
	fmt.Println("\nall CamKit proofs passed")
}

func proveBranchLR(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeNormalBP)
	para.SetBranchLRs(1.0, 0.0)
	w0, w1 := harness.DenseWeights(para, 0), harness.DenseWeights(para, 1)
	_, _ = harness.TrainN(para, x, y, 20, 0.1)
	d0 := harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0))
	d1 := harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
	harness.Requiref("BranchLRs_1|0", d0 > 1e-4 && d1 < 1e-7, "cam0 Δ=%.4g cam1 Δ=%.4g", d0, d1)
}

func proveRotate(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetRotateSchedule([][]parallel.TrainMode{
		{parallel.ModeNormalBP, parallel.ModeFreeze},
		{parallel.ModeFreeze, parallel.ModeNormalBP},
	}, 5)

	// Slot 0: only cam0 should move
	w0, w1 := harness.DenseWeights(para, 0), harness.DenseWeights(para, 1)
	for i := 0; i < 5; i++ {
		_, _ = parallel.TrainMSE(para, x, y, parallel.ModeNormalBP, 0.1)
		para.AdvanceRotate()
	}
	d0a := harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0))
	d1a := harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))

	// Now on slot 1 after 5 advances — next 5 steps only cam1 moves
	w0, w1 = harness.DenseWeights(para, 0), harness.DenseWeights(para, 1)
	for i := 0; i < 5; i++ {
		_, _ = parallel.TrainMSE(para, x, y, parallel.ModeNormalBP, 0.1)
		para.AdvanceRotate()
	}
	d0b := harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0))
	d1b := harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))

	// TrainMSE also calls AdvanceRotate — careful. We manually AdvanceRotate after each
	// TrainMSE which double-advances if TrainMSE already advances. Looking at TrainMSE —
	// it calls AdvanceRotate at end. So our loop double-counts!
	// Fix: only use TrainMSE without extra Advance, period=5, run 5 then snapshot.
	_ = d0a
	_ = d1a
	_ = d0b
	_ = d1b

	// Cleaner redo:
	para, _ = harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetRotateSchedule([][]parallel.TrainMode{
		{parallel.ModeNormalBP, parallel.ModeFreeze},
		{parallel.ModeFreeze, parallel.ModeNormalBP},
	}, 5)
	w0, w1 = harness.DenseWeights(para, 0), harness.DenseWeights(para, 1)
	_, _ = harness.TrainN(para, x, y, 5, 0.1) // TrainMSE advances rotate each sample
	d0a = harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0))
	d1a = harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))
	w0, w1 = harness.DenseWeights(para, 0), harness.DenseWeights(para, 1)
	_, _ = harness.TrainN(para, x, y, 5, 0.1)
	d0b = harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0))
	d1b = harness.WeightMaxDiff(w1, harness.DenseWeights(para, 1))

	harness.Requiref("Rotate_slot0", d0a > 1e-4 && d1a < 1e-7,
		"first 5: cam0 Δ=%.4g cam1 Δ=%.4g", d0a, d1a)
	harness.Requiref("Rotate_slot1", d0b < 1e-7 && d1b > 1e-4,
		"next 5: cam0 Δ=%.4g cam1 Δ=%.4g", d0b, d1b)
}

func proveDNA(x, y *core.Tensor[float32]) {
	// Diversify: start mid-close after soft sync, DNAReg>0 should lower cos.
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetCamSync(parallel.CamSyncConfig{Enabled: true, Alpha: 0.5, When: parallel.SyncManual})
	_ = para.SyncNow()
	cosMid := harness.Cos01(para)
	para.SetCamKit(parallel.CamKit{DNAReg: 0.5})
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeNormalBP)
	// applyDNAReg runs inside trainParallelMixed — need mixed modes / train
	_, _ = harness.TrainN(para, x, y, 5, 0.01)
	cosDiv := harness.Cos01(para)
	harness.Requiref("DNAReg_diversify", cosDiv < cosMid-0.01,
		"diversify cos %.4f → %.4f", cosMid, cosDiv)

	// Attract: diverge hard, DNAReg<0 should raise cos.
	para, _ = harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	cosLo := harness.Cos01(para)
	para.SetCamKit(parallel.CamKit{DNAReg: -0.5})
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeNormalBP)
	_, _ = harness.TrainN(para, x, y, 5, 0.01)
	cosHi := harness.Cos01(para)
	harness.Requiref("DNAReg_attract", cosHi > cosLo+0.05,
		"attract cos %.4f → %.4f", cosLo, cosHi)
}

func proveDream(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeFreeze)
	para.SetCamKit(parallel.CamKit{Dream: &parallel.DreamBuffer{Cap: 64}})
	_, _ = harness.TrainN(para, x, y, 8, 0.08)
	w0 := harness.DenseWeights(para, 0)
	avg, err := para.DreamPulse(4, parallel.ModeNormalBP, 0.05)
	if err != nil {
		panic(err)
	}
	d0 := harness.WeightMaxDiff(w0, harness.DenseWeights(para, 0))
	harness.Requiref("DreamPulse", avg > 0 && d0 > 1e-5,
		"replay avgLoss=%.4f moved cam0 Δ=%.4g", avg, d0)
}

func proveMetrics(x, y *core.Tensor[float32]) {
	para, _ := harness.DenseTwin(8, 4, parallel.CombineAvg)
	_ = harness.DivergeCams(para)
	para.SetBranchModes(parallel.ModeNormalBP, parallel.ModeFreeze)
	_, _ = harness.TrainN(para, x, y, 10, 0.1)
	m := para.RefreshMetrics()
	harness.Requiref("RefreshMetrics",
		len(m.ActiveModes) == 2 && m.ActiveModes[1] == "Freeze" && len(m.Plasticity) == 2 && m.Plasticity[1] == 0,
		"modes=%v plast=%v", m.ActiveModes, m.Plasticity)
}

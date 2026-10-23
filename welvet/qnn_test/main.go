// qnn_test — Go CPU vs Hexagon NPU output deviation (same weights / inputs).
//
//	$env:PATH = "c:\git\welvet\native\qnn\windows_arm64;$env:PATH"
//	cd example/welvet && go run ./qnn_test
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"

	"github.com/openfluke/welvet/core"
	"github.com/openfluke/welvet/layers/dense"
	"github.com/openfluke/welvet/layers/layernorm"
	"github.com/openfluke/welvet/layers/mha"
	"github.com/openfluke/welvet/layers/rmsnorm"
	"github.com/openfluke/welvet/layers/softmax"
	"github.com/openfluke/welvet/layers/swiglu"
	"github.com/openfluke/welvet/nativeqnn"
	"github.com/openfluke/welvet/quant"
)

type sample struct {
	I        int
	Go, NPU  float32
	AbsDelta float64
}

type diff struct {
	N          int
	ExactPct   float64 // bit-identical elements
	DiffPct    float64 // 100 - ExactPct
	MaxAbs     float64
	MeanAbs    float64
	RMSE       float64
	MeanRelPct float64 // mean(|Δ| / max(|go|,|npu|,eps)) * 100
	MaxRelPct  float64
	Cosine     float64
	PctGT1e3   float64 // % elems with |Δ| > 1e-3
	PctGT1e2   float64
	PctGT1e1   float64
	Head       []sample // first few elems
	Worst      sample   // max |Δ|
}

func compare(goOut, npuOut []float32) diff {
	n := len(goOut)
	if len(npuOut) < n {
		n = len(npuOut)
	}
	d := diff{N: n, Cosine: 1}
	if n == 0 {
		return d
	}
	const eps = 1e-8
	var exact int
	var sumAbs, sumSq, sumRel, dot, ng, nq float64
	var maxAbs, maxRel float64
	var maxIdx int
	var gt3, gt2, gt1 int
	headN := 6
	if headN > n {
		headN = n
	}
	d.Head = make([]sample, headN)
	for i := 0; i < n; i++ {
		a, b := float64(goOut[i]), float64(npuOut[i])
		delta := math.Abs(a - b)
		if i < headN {
			d.Head[i] = sample{I: i, Go: goOut[i], NPU: npuOut[i], AbsDelta: delta}
		}
		if delta == 0 {
			exact++
		}
		sumAbs += delta
		sumSq += delta * delta
		if delta > maxAbs {
			maxAbs = delta
			maxIdx = i
		}
		den := math.Max(math.Max(math.Abs(a), math.Abs(b)), eps)
		rel := delta / den
		sumRel += rel
		if rel > maxRel {
			maxRel = rel
		}
		if delta > 1e-3 {
			gt3++
		}
		if delta > 1e-2 {
			gt2++
		}
		if delta > 1e-1 {
			gt1++
		}
		dot += a * b
		ng += a * a
		nq += b * b
	}
	d.ExactPct = 100 * float64(exact) / float64(n)
	d.DiffPct = 100 - d.ExactPct
	d.MaxAbs = maxAbs
	d.MeanAbs = sumAbs / float64(n)
	d.RMSE = math.Sqrt(sumSq / float64(n))
	d.MeanRelPct = 100 * sumRel / float64(n)
	d.MaxRelPct = 100 * maxRel
	d.PctGT1e3 = 100 * float64(gt3) / float64(n)
	d.PctGT1e2 = 100 * float64(gt2) / float64(n)
	d.PctGT1e1 = 100 * float64(gt1) / float64(n)
	d.Worst = sample{I: maxIdx, Go: goOut[maxIdx], NPU: npuOut[maxIdx], AbsDelta: maxAbs}
	if ng > 0 && nq > 0 {
		d.Cosine = dot / (math.Sqrt(ng) * math.Sqrt(nq))
	}
	return d
}

func fillRand(dst []float32, seed int64, scale float32) {
	r := rand.New(rand.NewSource(seed))
	for i := range dst {
		dst[i] = (r.Float32()*2 - 1) * scale
	}
}

func ones(n int) []float32 {
	out := make([]float32, n)
	for i := range out {
		out[i] = 1
	}
	return out
}

func printRow(name, geom, note string, d diff, err error) {
	if err != nil {
		fmt.Printf("%-10s %-28s  FAIL %v\n", name, geom, err)
		return
	}
	fmt.Printf("%-10s %-28s  differ=%6.2f%%  meanRel=%7.3f%%  maxRel=%8.2f%%  max|Δ|=%-10.3g RMSE=%-10.3g cos=%7.4f  |Δ|>1e-3=%5.1f%%  >1e-2=%5.1f%%\n",
		name, geom, d.DiffPct, d.MeanRelPct, d.MaxRelPct, d.MaxAbs, d.RMSE, d.Cosine, d.PctGT1e3, d.PctGT1e2)
	if note != "" {
		fmt.Printf("           note: %s\n", note)
	}
	fmt.Printf("           samples (idx | go | npu | |Δ|):\n")
	for _, s := range d.Head {
		fmt.Printf("             [%4d]  % .6g  % .6g  % .3g\n", s.I, s.Go, s.NPU, s.AbsDelta)
	}
	w := d.Worst
	fmt.Printf("           worst      [%4d]  go=% .6g  npu=% .6g  |Δ|=%.3g\n", w.I, w.Go, w.NPU, w.AbsDelta)
}

func main() {
	if !nativeqnn.Available() {
		fmt.Println("SKIP: need windows/arm64 + libwelvet_qnn on PATH (+ QAIRT)")
		os.Exit(0)
	}
	sess, err := nativeqnn.DefaultSession()
	if err != nil {
		fmt.Println("open HTP:", err)
		os.Exit(1)
	}

	fmt.Println("qnn_test — Go CPU (tiled) vs NPU  ·  welvet.qnn", nativeqnn.Version())
	fmt.Println("differ%  = % of elements not bit-identical to Go (HTP often rounds → ~100% even when close)")
	fmt.Println("meanRel% = mean( |go−npu| / max(|go|,|npu|,1e-8) ) × 100   ← main “how different” number")
	fmt.Println("cos      = cosine similarity (≈1 = same behaviour)")
	fmt.Println(strings.Repeat("─", 120))

	runDense(sess)
	runSoftmax(sess)
	runRMS(sess)
	runLN(sess)
	runCnn2(sess)
	runEmbed(sess)
	runResidual(sess)
	runSwiGlu(sess)
	runMha(sess)

	fmt.Println(strings.Repeat("─", 120))
	fmt.Println("done")
}

func runDense(s *nativeqnn.Session) {
	const batch, in, out = 8, 128, 128
	geom := fmt.Sprintf("B=%d in=%d out=%d", batch, in, out)
	W := make([]float32, out*in)
	bias := make([]float32, out)
	x := make([]float32, batch*in)
	fillRand(W, 1, 0.05)
	fillRand(bias, 2, 0.01)
	fillRand(x, 3, 0.1)
	yQ := make([]float32, batch*out)

	if err := s.DensePrepareF32(batch, in, out, W, bias); err != nil {
		printRow("dense", geom, "", diff{}, err)
		return
	}
	if err := s.DenseExecuteF32(x, yQ); err != nil {
		printRow("dense", geom, "", diff{}, err)
		return
	}

	l, err := dense.NewConfigured(in, out, core.ActivationLinear, core.DTypeFloat32, quant.FormatNone, W)
	if err != nil {
		printRow("dense", geom, "", diff{}, err)
		return
	}
	bf := make([]float64, out)
	for i := range bias {
		bf[i] = float64(bias[i])
	}
	l.Weights.Bias = bf
	l.Exec.Backend = core.BackendCPUTiled
	tx := core.NewTensor[float32](batch, in)
	copy(tx.Data, x)
	_, post, err := dense.Forward(l, tx)
	if err != nil {
		printRow("dense", geom, "", diff{}, err)
		return
	}
	printRow("dense", geom, "", compare(post.Data, yQ), nil)
}

func runSoftmax(s *nativeqnn.Session) {
	const rows, cols = 16, 256
	geom := fmt.Sprintf("rows=%d cols=%d", rows, cols)
	x := make([]float32, rows*cols)
	fillRand(x, 4, 2)
	yQ := make([]float32, len(x))
	if err := s.SoftmaxForwardF32(rows, cols, x, yQ); err != nil {
		printRow("softmax", geom, "", diff{}, err)
		return
	}
	l, _ := softmax.New(softmax.Config{Dim: cols})
	l.Exec.Backend = core.BackendCPUTiled
	t := core.NewTensor[float32](rows, cols)
	copy(t.Data, x)
	_, post, err := softmax.Forward(l, t)
	if err != nil {
		printRow("softmax", geom, "", diff{}, err)
		return
	}
	printRow("softmax", geom, "", compare(post.Data, yQ), nil)
}

func runRMS(s *nativeqnn.Session) {
	const batch, seq, dim = 4, 32, 128
	geom := fmt.Sprintf("B=%d S=%d D=%d", batch, seq, dim)
	gamma := ones(dim)
	x := make([]float32, batch*seq*dim)
	fillRand(x, 5, 0.5)
	yQ := make([]float32, len(x))
	if err := s.RmsNormForwardF32(batch, seq, dim, 1e-6, gamma, x, yQ); err != nil {
		printRow("rmsnorm", geom, "", diff{}, err)
		return
	}
	l, err := rmsnorm.NewConfigured(rmsnorm.Config{Dim: dim, Eps: 1e-6}, core.DTypeFloat32, quant.FormatNone, gamma)
	if err != nil {
		printRow("rmsnorm", geom, "", diff{}, err)
		return
	}
	l.Exec.Backend = core.BackendCPUTiled
	t := core.NewTensor[float32](batch, seq, dim)
	copy(t.Data, x)
	_, post, err := rmsnorm.Forward(l, t)
	if err != nil {
		printRow("rmsnorm", geom, "", diff{}, err)
		return
	}
	printRow("rmsnorm", geom, "", compare(post.Data, yQ), nil)
}

func runLN(s *nativeqnn.Session) {
	const batch, seq, dim = 4, 32, 128
	geom := fmt.Sprintf("B=%d S=%d D=%d", batch, seq, dim)
	gamma, beta := ones(dim), make([]float32, dim)
	x := make([]float32, batch*seq*dim)
	fillRand(x, 6, 0.4)
	yQ := make([]float32, len(x))
	if err := s.LayerNormForwardF32(batch, seq, dim, 1e-6, gamma, beta, x, yQ); err != nil {
		printRow("layernorm", geom, "", diff{}, err)
		return
	}
	l, err := layernorm.NewConfigured(layernorm.Config{Dim: dim, Eps: 1e-6}, core.DTypeFloat32, quant.FormatNone, gamma, beta)
	if err != nil {
		printRow("layernorm", geom, "", diff{}, err)
		return
	}
	l.Exec.Backend = core.BackendCPUTiled
	t := core.NewTensor[float32](batch, seq, dim)
	copy(t.Data, x)
	_, post, err := layernorm.Forward(l, t)
	if err != nil {
		printRow("layernorm", geom, "", diff{}, err)
		return
	}
	printRow("layernorm", geom, "", compare(post.Data, yQ), nil)
}

func runCnn2(s *nativeqnn.Session) {
	// HTP path is 1×1 (W = cout*cin) — match Go GEMV-per-pixel.
	const n, h, w, cin, cout = 2, 16, 16, 8, 8
	geom := fmt.Sprintf("N=%d %dx%d cin=%d cout=%d k=1x1", n, h, w, cin, cout)
	W := make([]float32, cout*cin)
	bias := make([]float32, cout)
	x := make([]float32, n*h*w*cin)
	fillRand(W, 7, 0.05)
	fillRand(x, 8, 0.1)
	yQ := make([]float32, n*h*w*cout)
	if err := s.Cnn2ForwardF32(n, h, w, cin, cout, W, bias, x, yQ); err != nil {
		printRow("cnn2", geom, "", diff{}, err)
		return
	}
	yG := make([]float32, len(yQ))
	spatial := h * w
	for ni := 0; ni < n; ni++ {
		for sp := 0; sp < spatial; sp++ {
			xi := (ni*spatial + sp) * cin
			yi := (ni*spatial + sp) * cout
			for oc := 0; oc < cout; oc++ {
				var acc float32
				base := oc * cin
				for ic := 0; ic < cin; ic++ {
					acc += W[base+ic] * x[xi+ic]
				}
				yG[yi+oc] = acc // HTP ignores bias today
			}
		}
	}
	printRow("cnn2", geom, "Go ref = NHWC 1×1 GEMV (same as w2a)", compare(yG, yQ), nil)
}

func runEmbed(s *nativeqnn.Session) {
	const vocab, dim, nTok = 512, 64, 64
	geom := fmt.Sprintf("vocab=%d dim=%d n=%d", vocab, dim, nTok)
	table := make([]float32, vocab*dim)
	ids := make([]int32, nTok)
	fillRand(table, 18, 0.05)
	for i := range ids {
		ids[i] = int32(i % vocab)
	}
	yQ := make([]float32, nTok*dim)
	if err := s.EmbeddingForwardF32(vocab, dim, nTok, table, ids, yQ); err != nil {
		printRow("embedding", geom, "", diff{}, err)
		return
	}
	yG := make([]float32, nTok*dim)
	for i, id := range ids {
		off := int(id) * dim
		copy(yG[i*dim:(i+1)*dim], table[off:off+dim])
	}
	printRow("embedding", geom, "", compare(yG, yQ), nil)
}

func runResidual(s *nativeqnn.Session) {
	const n = 8192
	geom := fmt.Sprintf("n=%d", n)
	a, b := make([]float32, n), make([]float32, n)
	fillRand(a, 19, 0.5)
	fillRand(b, 20, 0.5)
	yQ := make([]float32, n)
	if err := s.ResidualForwardF32(n, a, b, yQ); err != nil {
		printRow("residual", geom, "", diff{}, err)
		return
	}
	yG := make([]float32, n)
	for i := 0; i < n; i++ {
		yG[i] = a[i] + b[i]
	}
	printRow("residual", geom, "", compare(yG, yQ), nil)
}

func runSwiGlu(s *nativeqnn.Session) {
	const batch, in, hidden = 2, 64, 128
	geom := fmt.Sprintf("B=%d in=%d hidden=%d", batch, in, hidden)
	wG := make([]float32, hidden*in)
	wU := make([]float32, hidden*in)
	wD := make([]float32, in*hidden)
	x := make([]float32, batch*in)
	fillRand(wG, 21, 0.02)
	fillRand(wU, 22, 0.02)
	fillRand(wD, 23, 0.02)
	fillRand(x, 24, 0.1)
	yQ := make([]float32, batch*in)
	if err := s.SwiGluForwardF32(batch, in, hidden, wG, wU, wD, x, yQ); err != nil {
		printRow("swiglu", geom, "", diff{}, err)
		return
	}
	l, err := swiglu.NewConfigured(swiglu.Config{InputDim: in, IntermediateDim: hidden}, core.DTypeFloat32, quant.FormatNone, wG, wU, wD)
	if err != nil {
		printRow("swiglu", geom, "", diff{}, err)
		return
	}
	l.Exec.Backend = core.BackendCPUTiled
	tx := core.NewTensor[float32](batch, in)
	copy(tx.Data, x)
	_, post, err := swiglu.Forward(l, tx)
	if err != nil {
		printRow("swiglu", geom, "", diff{}, err)
		return
	}
	// NPU compose uses sigmoid(gate), Go SwiGLU uses SiLU — show both.
	printRow("swiglu", geom, "vs Go SiLU (NPU uses sigmoid gate)", compare(post.Data, yQ), nil)
	ySig := swigluSigmoidRef(batch, in, hidden, wG, wU, wD, x)
	printRow("swiglu~", geom, "vs sigmoid-gate ref (NPU semantics)", compare(ySig, yQ), nil)
}

func swigluSigmoidRef(batch, in, hidden int, wG, wU, wD, x []float32) []float32 {
	// y = ((σ(xWg) ⊙ (xWu)) Wd)  — matches welvet.qnn SwiGlu::forward
	gate := make([]float32, batch*hidden)
	up := make([]float32, batch*hidden)
	h := make([]float32, batch*hidden)
	y := make([]float32, batch*in)
	for b := 0; b < batch; b++ {
		xb := x[b*in : (b+1)*in]
		for j := 0; j < hidden; j++ {
			var g, u float32
			for i := 0; i < in; i++ {
				g += xb[i] * wG[j*in+i]
				u += xb[i] * wU[j*in+i]
			}
			sig := float32(1 / (1 + math.Exp(float64(-g))))
			gate[b*hidden+j] = sig
			up[b*hidden+j] = u
			h[b*hidden+j] = sig * u
		}
		hb := h[b*hidden : (b+1)*hidden]
		for i := 0; i < in; i++ {
			var acc float32
			for j := 0; j < hidden; j++ {
				acc += hb[j] * wD[i*hidden+j]
			}
			y[b*in+i] = acc
		}
	}
	return y
}

func runMha(s *nativeqnn.Session) {
	const batch, seq, dim, heads = 1, 8, 32, 4
	geom := fmt.Sprintf("B=%d S=%d D=%d H=%d", batch, seq, dim, heads)
	wQ := make([]float32, dim*dim)
	wK := make([]float32, dim*dim)
	wV := make([]float32, dim*dim)
	wO := make([]float32, dim*dim)
	x := make([]float32, batch*seq*dim)
	for i := 0; i < dim; i++ {
		wQ[i*dim+i], wK[i*dim+i], wV[i*dim+i], wO[i*dim+i] = 1, 1, 1, 1
	}
	fillRand(x, 25, 0.05)
	yQ := make([]float32, len(x))
	if err := s.MhaForwardF32(batch, seq, dim, heads, wQ, wK, wV, wO, x, yQ); err != nil {
		printRow("mha", geom, "", diff{}, err)
		return
	}
	l, err := mha.NewConfigured(mha.Config{DModel: dim, NumHeads: heads, Causal: true}, core.DTypeFloat32, quant.FormatNone, wQ, wK, wV, wO)
	if err != nil {
		printRow("mha", geom, "", diff{}, err)
		return
	}
	l.Exec.Backend = core.BackendCPUTiled
	tx := core.NewTensor[float32](batch, seq, dim)
	copy(tx.Data, x)
	_, post, err := mha.Forward(l, tx)
	if err != nil {
		printRow("mha", geom, "", diff{}, err)
		return
	}
	printRow("mha", geom, "vs Go causal MHA (NPU compose may differ on mask/scale)", compare(post.Data, yQ), nil)
}

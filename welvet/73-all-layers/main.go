package main

import (
	"fmt"
	"os"

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
	"github.com/openfluke/welvet/layers/seqmix"
	"github.com/openfluke/welvet/layers/sequential"
	"github.com/openfluke/welvet/layers/softmax"
	"github.com/openfluke/welvet/layers/swiglu"
	"github.com/openfluke/welvet/quant"
	"github.com/openfluke/welvet/runtime/dispatch"
)

type demo struct {
	name, when, where, why string
	run                    func() error
}

func main() {
	demos := []demo{
		{"dense", "linear / FFN / proj", "layers/dense", "MatVec microkernel everything else builds on", runDense},
		{"cnn1", "1D seq / audio patches", "layers/cnn1", "im2col→Dense Proj", runCNN1},
		{"cnn2", "images / 2D maps", "layers/cnn2", "vision stem", runCNN2},
		{"cnn3", "volumes / voxels", "layers/cnn3", "3D conv", runCNN3},
		{"convt1", "1D upsample", "layers/convt1", "decoder / generative 1D", runConvT1},
		{"convt2", "2D upsample", "layers/convt2", "decoder / U-Net expand", runConvT2},
		{"convt3", "3D upsample", "layers/convt3", "volumetric decoder", runConvT3},
		{"mha", "token mixing via attention", "layers/mha", "transformers", runMHA},
		{"swiglu", "gated FFN", "layers/swiglu", "modern LLM MLP", runSwiGLU},
		{"rmsnorm", "stabilize activations", "layers/rmsnorm", "LLM-style norm", runRMS},
		{"layernorm", "stabilize activations", "layers/layernorm", "classic transformer norm", runLN},
		{"softmax", "distributions / heads", "layers/softmax", "classify / combine", runSoftmax},
		{"embedding", "token → vector", "layers/embedding", "lookup tables", runEmbed},
		{"rnn", "short temporal", "layers/rnn", "classic recurrent", runRNN},
		{"lstm", "gated temporal", "layers/lstm", "longer memory than RNN", runLSTM},
		{"sequential", "stack Dense depth", "layers/sequential", "deep hemisphere body", runSeq},
		{"residual", "skip + body", "layers/residual", "stable depth", runRes},
		{"parallel", "multi-cam / MoE", "layers/parallel", "combine cams + CamSync host", runPara},
		{"kmeans", "prototypes / codebook", "layers/kmeans", "cluster features", runKMeans},
		{"mamba", "selective SSM", "layers/mamba", "linear-time seq mix", runMamba},
		{"gdn", "gated delta / linear attn", "layers/gdn", "decode-friendly mixer", runGDN},
		{"metacognition", "observer / confidence", "layers/metacognition", "meta signals beside main path", runMeta},
		{"seqmix", "name the mixer kind", "layers/seqmix", "contract only — swap MHA/Mamba/GDN", runSeqMix},
		{"dispatch", "any Op → Forward", "runtime/dispatch", "grid walk without type switches in your code", runDispatch},
	}

	fail := 0
	for _, d := range demos {
		if err := d.run(); err != nil {
			fmt.Printf("FAIL %-14s %v\n", d.name, err)
			fail++
			continue
		}
		fmt.Printf("OK   %-14s when=%s | where=%s | why=%s\n", d.name, d.when, d.where, d.why)
	}
	fmt.Printf("\nall-layers: ok=%d fail=%d\n", len(demos)-fail, fail)
	if fail > 0 {
		os.Exit(1)
	}
}

func fill(shape ...int) *core.Tensor[float32] {
	t := core.NewTensor[float32](shape...)
	for i := range t.Data {
		t.Data[i] = 0.1 + float32(i%7)*0.03
	}
	return t
}

func runDense() error {
	l, err := dense.New(4, 4, core.ActivationReLU, core.DTypeFloat32)
	if err != nil {
		return err
	}
	_, _, err = dense.Forward(l, fill(1, 4))
	return err
}
func runCNN1() error {
	l, err := cnn1.New(cnn1.Config{InChannels: 1, Filters: 2, SeqLen: 4, Kernel: 3})
	if err != nil {
		return err
	}
	_, _, err = cnn1.Forward(l, fill(1, 1, 4))
	return err
}
func runCNN2() error {
	l, err := cnn2.New(cnn2.Config{InChannels: 1, Filters: 2, Height: 3, Width: 3, Kernel: 2})
	if err != nil {
		return err
	}
	_, _, err = cnn2.Forward(l, fill(1, 1, 3, 3))
	return err
}
func runCNN3() error {
	l, err := cnn3.New(cnn3.Config{InChannels: 1, Filters: 2, Depth: 2, Height: 2, Width: 2, Kernel: 2})
	if err != nil {
		return err
	}
	_, _, err = cnn3.Forward(l, fill(1, 1, 2, 2, 2))
	return err
}
func runConvT1() error {
	l, err := convt1.New(convt1.Config{InChannels: 1, Filters: 2, SeqLen: 2, Kernel: 3})
	if err != nil {
		return err
	}
	_, _, err = convt1.Forward(l, fill(1, 1, 2))
	return err
}
func runConvT2() error {
	l, err := convt2.New(convt2.Config{InChannels: 1, Filters: 2, Height: 2, Width: 2, Kernel: 2})
	if err != nil {
		return err
	}
	_, _, err = convt2.Forward(l, fill(1, 1, 2, 2))
	return err
}
func runConvT3() error {
	l, err := convt3.New(convt3.Config{InChannels: 1, Filters: 2, Depth: 2, Height: 2, Width: 2, Kernel: 2})
	if err != nil {
		return err
	}
	_, _, err = convt3.Forward(l, fill(1, 1, 2, 2, 2))
	return err
}
func runMHA() error {
	l, err := mha.New(mha.Config{DModel: 4, NumHeads: 1, MaxSeqLen: 4})
	if err != nil {
		return err
	}
	_, _, err = mha.Forward(l, fill(1, 2, 4))
	return err
}
func runSwiGLU() error {
	l, err := swiglu.New(swiglu.Config{InputDim: 4, IntermediateDim: 8})
	if err != nil {
		return err
	}
	_, _, err = swiglu.Forward(l, fill(1, 4))
	return err
}
func runRMS() error {
	l, err := rmsnorm.New(rmsnorm.Config{Dim: 4})
	if err != nil {
		return err
	}
	x := fill(1, 4)
	for i := range x.Data {
		x.Data[i] = float32(i + 1)
	}
	_, _, err = rmsnorm.Forward(l, x)
	return err
}
func runLN() error {
	l, err := layernorm.New(layernorm.Config{Dim: 4})
	if err != nil {
		return err
	}
	x := fill(1, 4)
	for i := range x.Data {
		x.Data[i] = float32(i + 1)
	}
	_, _, err = layernorm.Forward(l, x)
	return err
}
func runSoftmax() error {
	l, err := softmax.New(softmax.Config{Dim: 4})
	if err != nil {
		return err
	}
	_, _, err = softmax.Forward(l, fill(1, 4))
	return err
}
func runEmbed() error {
	l, err := embedding.New(embedding.Config{VocabSize: 8, EmbeddingDim: 4, SeqLen: 2})
	if err != nil {
		return err
	}
	x := fill(1, 2)
	x.Data[0], x.Data[1] = 1, 2
	_, _, err = embedding.Forward(l, x)
	return err
}
func runRNN() error {
	l, err := rnn.New(rnn.Config{InputSize: 3, HiddenSize: 4, SeqLen: 2})
	if err != nil {
		return err
	}
	_, _, err = rnn.Forward(l, fill(1, 2, 3))
	return err
}
func runLSTM() error {
	l, err := lstm.New(lstm.Config{InputSize: 3, HiddenSize: 4, SeqLen: 2})
	if err != nil {
		return err
	}
	_, _, err = lstm.Forward(l, fill(1, 2, 3))
	return err
}
func runSeq() error {
	l, err := sequential.New(sequential.Config{Dim: 4, Depth: 2})
	if err != nil {
		return err
	}
	_, _, err = sequential.Forward(l, fill(1, 4))
	return err
}
func runRes() error {
	l, err := residual.New(residual.Config{Dim: 4, Depth: 1})
	if err != nil {
		return err
	}
	_, _, err = residual.Forward(l, fill(1, 4))
	return err
}
func runPara() error {
	s, err := parallel.Bicameral(4, 8, 1, core.ActivationReLU, core.DTypeFloat32, quant.FormatNone)
	if err != nil {
		return err
	}
	_, _, err = parallel.ForwardStack(s, fill(1, 4))
	return err
}
func runKMeans() error {
	l, err := kmeans.New(kmeans.Config{NumClusters: 3, FeatureDim: 4})
	if err != nil {
		return err
	}
	_, _, err = kmeans.Forward(l, fill(1, 4))
	return err
}
func runMamba() error {
	l, err := mamba.New(mamba.Config{DModel: 4, DState: 2, SeqLen: 2})
	if err != nil {
		return err
	}
	_, _, err = mamba.Forward(l, fill(1, 2, 4))
	return err
}
func runGDN() error {
	l, err := gdn.New(gdn.Config{HiddenSize: 4, NumKeyHeads: 1, NumValueHeads: 1, KeyHeadDim: 2, ValueHeadDim: 2, ConvKernel: 2})
	if err != nil {
		return err
	}
	_, _, err = gdn.Forward(l, fill(1, 2, 4))
	return err
}
func runMeta() error {
	l, err := metacognition.New(metacognition.Config{Dim: 4})
	if err != nil {
		return err
	}
	_, _, err = metacognition.Forward(l, fill(1, 4))
	return err
}
func runSeqMix() error {
	c := seqmix.Contract{Kind: seqmix.KindSSM, DModel: 4, MaxT: 8}
	if c.Kind.String() != "ssm" {
		return fmt.Errorf("kind=%s", c.Kind)
	}
	return nil
}
func runDispatch() error {
	l, err := dense.New(4, 4, core.ActivationLinear, core.DTypeFloat32)
	if err != nil {
		return err
	}
	_, _, err = dispatch.Forward(l, fill(1, 4))
	return err
}

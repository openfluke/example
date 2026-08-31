package main

import (
	"fmt"
	"strings"

	"github.com/openfluke/welvet/core"
	"github.com/openfluke/welvet/layers/cnn2"
	"github.com/openfluke/welvet/layers/dense"
	"github.com/openfluke/welvet/runtime/dispatch"
)

func main() {
	// When: you hold an `any` Op (grid cell, Parallel branch, Sequential child)
	// Where: runtime/dispatch — single switch for Forward/Backward/Pack/ApplyGrad
	// Why: no silent Dense fallback; unknown Ops hard-error

	d, err := dense.New(4, 4, core.ActivationReLU, core.DTypeFloat32)
	must(err)
	x := core.NewTensor[float32](1, 4)
	for i := range x.Data {
		x.Data[i] = 0.25
	}
	pre, post, err := dispatch.Forward(d, x)
	must(err)
	gIn, gW, err := dispatch.Backward(d, post, x, pre)
	must(err)
	must(dispatch.ApplyGradSGD(d, gW, 1e-2))
	fmt.Println("dispatch dense", post.Shape, "gIn", len(gIn.Data), "gW", len(gW.Data))

	c, err := cnn2.New(cnn2.Config{InChannels: 1, Filters: 2, Height: 4, Width: 4, Kernel: 3})
	must(err)
	img := core.NewTensor[float32](1, 1, 4, 4)
	_, y, err := dispatch.Forward(c, img)
	must(err)
	fmt.Println("dispatch cnn2", y.Shape)

	_, _, err = dispatch.Forward("not-an-op", x)
	if err == nil || !strings.Contains(err.Error(), "unsupported Op") {
		panic(fmt.Sprintf("want unsupported Op error, got %v", err))
	}
	fmt.Println("dispatch rejects unknown:", err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"

	"github.com/openfluke/welvet/core"
	"github.com/openfluke/welvet/layers/cnn1"
	"github.com/openfluke/welvet/layers/cnn2"
	"github.com/openfluke/welvet/layers/cnn3"
)

func main() {
	// cnn1 — 1D conv (seq / audio patches)
	c1, err := cnn1.New(cnn1.Config{InChannels: 1, Filters: 4, SeqLen: 16, Kernel: 3, Stride: 1})
	must(err)
	x1 := core.NewTensor[float32](1, 1, 16)
	for i := range x1.Data {
		x1.Data[i] = 0.1
	}
	_, y1, err := cnn1.Forward(c1, x1)
	must(err)
	fmt.Println("cnn1", y1.Shape)

	// cnn2 — 2D vision
	c2, err := cnn2.New(cnn2.Config{InChannels: 1, Filters: 4, Height: 8, Width: 8, Kernel: 3, Stride: 1})
	must(err)
	x2 := core.NewTensor[float32](1, 1, 8, 8)
	for i := range x2.Data {
		x2.Data[i] = 0.1
	}
	_, y2, err := cnn2.Forward(c2, x2)
	must(err)
	fmt.Println("cnn2", y2.Shape)

	// cnn3 — 3D / volumetric
	c3, err := cnn3.New(cnn3.Config{InChannels: 1, Filters: 2, Depth: 4, Height: 4, Width: 4, Kernel: 2, Stride: 1})
	must(err)
	x3 := core.NewTensor[float32](1, 1, 4, 4, 4)
	for i := range x3.Data {
		x3.Data[i] = 0.1
	}
	_, y3, err := cnn3.Forward(c3, x3)
	must(err)
	fmt.Println("cnn3", y3.Shape)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

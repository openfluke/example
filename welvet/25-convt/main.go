package main

import (
	"fmt"

	"github.com/openfluke/welvet/core"
	"github.com/openfluke/welvet/layers/convt1"
	"github.com/openfluke/welvet/layers/convt2"
	"github.com/openfluke/welvet/layers/convt3"
)

func main() {
	t1, err := convt1.New(convt1.Config{InChannels: 4, Filters: 2, SeqLen: 8, Kernel: 3, Stride: 2})
	must(err)
	x1 := core.NewTensor[float32](1, 4, 8)
	_, y1, err := convt1.Forward(t1, x1)
	must(err)
	fmt.Println("convt1", y1.Shape)

	t2, err := convt2.New(convt2.Config{InChannels: 2, Filters: 2, Height: 4, Width: 4, Kernel: 2, Stride: 1})
	must(err)
	x2 := core.NewTensor[float32](1, 2, 4, 4)
	_, y2, err := convt2.Forward(t2, x2)
	must(err)
	fmt.Println("convt2", y2.Shape)

	t3, err := convt3.New(convt3.Config{InChannels: 1, Filters: 2, Depth: 2, Height: 2, Width: 2, Kernel: 2, Stride: 1})
	must(err)
	x3 := core.NewTensor[float32](1, 1, 2, 2, 2)
	_, y3, err := convt3.Forward(t3, x3)
	must(err)
	fmt.Println("convt3", y3.Shape)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

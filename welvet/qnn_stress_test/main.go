// qnn_stress_test — hammer Hexagon HTP for ~30s with a live countdown.
//
//	$env:PATH = "c:\git\welvet\native\qnn\windows_arm64;$env:PATH"
//	cd example/welvet/qnn_stress_test && go run .
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/openfluke/welvet/nativeqnn"
)

const (
	duration = 30 * time.Second
	batch    = 32
	inDim    = 512
	outDim   = 512
)

func main() {
	if !nativeqnn.Available() {
		fmt.Println("NO NPU — need windows/arm64 + libwelvet_qnn on PATH (+ QAIRT)")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("  ╔══════════════════════════════════════════╗")
	fmt.Println("  ║   WELVET · HEXAGON HTP STRESS TEST       ║")
	fmt.Println("  ╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("  welvet.qnn  %s\n", nativeqnn.Version())
	fmt.Printf("  layers      %v\n", nativeqnn.SupportedLayers())
	fmt.Printf("  workload    Dense %dx%dx%d  prepare-once + hammer execute\n", batch, inDim, outDim)
	fmt.Printf("  duration    %s\n", duration)
	fmt.Println()

	sess, err := nativeqnn.DefaultSession()
	must(err)

	W := make([]float32, outDim*inDim)
	bias := make([]float32, outDim)
	x := make([]float32, batch*inDim)
	y := make([]float32, batch*outDim)
	fill(W, 1, 0.02)
	fill(bias, 2, 0.01)
	fill(x, 3, 0.1)

	fmt.Print("  preparing HTP graph … ")
	tPrep := time.Now()
	must(sess.DensePrepareF32(batch, inDim, outDim, W, bias))
	fmt.Printf("ok (%s)\n", time.Since(tPrep).Round(time.Millisecond))
	must(sess.DenseExecuteF32(x, y)) // warm

	fmt.Println()
	fmt.Println("  GO — stressing NPU. Ctrl+C to abort.")
	fmt.Println()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	deadline := time.Now().Add(duration)
	start := time.Now()
	var ops int64
	var lastPhaseUS float64
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	done := false
	for !done {
		select {
		case <-stop:
			fmt.Println()
			fmt.Println("  aborted.")
			done = true
		case <-ticker.C:
			// burn as many executes as we can until next tick
			burstEnd := time.Now().Add(90 * time.Millisecond)
			for time.Now().Before(burstEnd) && time.Now().Before(deadline) {
				if err := sess.DenseExecuteF32(x, y); err != nil {
					fmt.Printf("\n  FAIL exec: %v\n", err)
					os.Exit(1)
				}
				ops++
				ph := sess.LastPhase()
				lastPhaseUS = ph.ExecuteUS()
			}
			left := time.Until(deadline)
			if left < 0 {
				left = 0
				done = true
			}
			elapsed := time.Since(start)
			opsPerSec := float64(ops) / elapsed.Seconds()
			bar := progressBar(1-left.Seconds()/duration.Seconds(), 28)
			// \r rewrite countdown line
			fmt.Printf("\r  %s  %5.1fs left  │  ops=%-7d  %.0f exec/s  │  last DLL exec≈%.0fµs   ",
				bar, left.Seconds(), ops, opsPerSec, lastPhaseUS)
		}
	}

	elapsed := time.Since(start)
	fmt.Println()
	fmt.Println()
	fmt.Println("  ╔══════════════════════════════════════════╗")
	fmt.Println("  ║   YES — YOU HAVE AN NPU                  ║")
	fmt.Println("  ╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("  ran         %s\n", elapsed.Round(time.Millisecond))
	fmt.Printf("  Dense exec  %d\n", ops)
	fmt.Printf("  throughput  %.0f graphExecute/s\n", float64(ops)/elapsed.Seconds())
	fmt.Printf("  last DLL    ≈%.0f µs  (graphExecute only)\n", lastPhaseUS)
	fmt.Printf("  checksum    y[0]=%.6g  y[mid]=%.6g  y[last]=%.6g\n",
		y[0], y[len(y)/2], y[len(y)-1])
	fmt.Println()
	fmt.Println("  Hexagon HTP stayed busy the whole countdown. Nice.")
	fmt.Println()
}

func progressBar(frac float64, width int) string {
	if frac < 0 {
		frac = 0
	}
	if frac > 1 {
		frac = 1
	}
	filled := int(frac * float64(width))
	if filled > width {
		filled = width
	}
	b := make([]byte, width+2)
	b[0] = '['
	for i := 0; i < width; i++ {
		if i < filled {
			b[i+1] = '#'
		} else {
			b[i+1] = '-'
		}
	}
	b[width+1] = ']'
	return string(b)
}

func fill(dst []float32, seed int64, scale float32) {
	r := rand.New(rand.NewSource(seed))
	for i := range dst {
		dst[i] = (r.Float32()*2 - 1) * scale
	}
}

func must(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

# 67. TrainMode — 29 named updates

**Part:** IV · Runtime  
**Package:** `github.com/openfluke/welvet/layers/parallel`  
**Status:** ok — ✅ 29 modes

## When

Building or training a net that needs the `layers/parallel` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/parallel"`

```bash
cd 67-train-modes && source ../env.sh && go run .
```

## Why

Backprop is one update, not the only one. Credit assignment (broadcast gap, head proxy, sparse duty clock) has to be a named axis you can race — not a comment in a notebook. Cameral Mix also needs one TrainMode per hemisphere on the same loss.

## What

parallel.TrainMode: AllNamedTrainModes() = 29 (Inherit omitted). Stack-local Split/Alt plus Step* 1D pipe twins and Mesh* grid schedulers. TrainStackMSE / TrainStackCE honour BranchModes. Display names use Short() / ShortTrainMode. Rival metric is hard Acc vs StepBP; Lucy Score is Tput × Avail × Acc — do not mix those sentences.

## Sample output (captured)

```
named 31 linestep 11
stepfastproxy Step[T][S][FP] <nil>
[T]=Tween  [S]=Split  [FP]=FastProxy  [L]=Linear  [HP]=HeadProxy  [F]=Freeze  [Sh]=Shadow  [A]=Adv  [M]=Memory
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/67-train-modes`).

# 21. layers/rnn · lstm

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/lstm`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/lstm` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/lstm"`

```bash
cd 21-rnn-lstm && source ../env.sh && go run .
```

## Why

Sequence models before transformers still need vanilla RNN and LSTM with BPTT on the shared MatVec stack.

## What

IH/HH (and LSTM gates) via Dense; recurrence ALU host; device required for WebGPU path.

## Sample output (captured)

```
64 64
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/21-rnn-lstm`).

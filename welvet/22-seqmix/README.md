# 22. layers/seqmix — mixer contract

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/seqmix`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/seqmix` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/seqmix"`

```bash
cd 22-seqmix && source ../env.sh && go run .
```

## Why

Attention, SSM, linear attn, and conv mixers must not be accidental forks of mha. Naming the contract keeps packages honest.

## What

KindAttention | KindSSM | KindLinearAttn | KindConvMix and Contract{Kind,DModel,MaxT}. No compute here.

## Sample output (captured)

```
attention
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/22-seqmix`).

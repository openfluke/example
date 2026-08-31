# 12. layers/mha — attention

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/mha`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/mha` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/mha"`

```bash
cd 12-mha && source ../env.sh && go run .
```

## Why

Transformers need multi-head attention with masks, RoPE/ALiBi, GQA/MQA, and cross-attn — without forking MatVec for every projection.

## What

Q/K/V/O are Dense children. Presets: DecoderCausal, EncoderBidirectional, CrossAttention, … Attn/RoPE ALU is host today; on-device attn shaders are still open.

## Sample output (captured)

```
true 256 <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/12-mha`).

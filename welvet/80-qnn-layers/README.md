# 80. QNN — HTP layer tour (session ABI)

**Part:** XI · Hexagon NPU (welvet.qnn)  
**Package:** `github.com/openfluke/welvet/nativeqnn`  
**Status:** ok — ✅ 12 layer forwards

## When

Smoke-test every shipped Hexagon forward without running the full w2a matrix.

## Where

```bash
cd 80-qnn-layers && source ../env.sh && go run .
```

## Why

Same session ABI as `w2a/suites/go_v_qnn`, tiny shapes. CNN/ConvT use **1×1** weights (`W = cout*cin`) — matches HTP today.

## What

Dense (prepare+exec), Softmax, RMSNorm, LayerNorm, CNN2/3, ConvT2/3, Embedding, Residual, SwiGLU, MHA.

See [feature book §80](https://openfluke.github.io/welvet/chapters/80-qnn-layers.html). Large RESULTS: §81.

## Sample output (captured)

```
(run locally — needs windows/arm64 libwelvet_qnn)
```

## Source

Mirrored for the [Welvet feature book](https://openfluke.github.io/welvet/) (`openfluke.github.io` §80).

# qnn_test — Go vs NPU output deviation

Same weights + inputs on **Go CPU tiled** and **Hexagon HTP** (`nativeqnn`). Reports how different the tensors are, plus **sample element values**.

## Run

```powershell
$env:PATH = "c:\git\welvet\native\qnn\windows_arm64;$env:PATH"
cd c:\git\example\welvet
go run ./qnn_test
```

## Columns

| Field | Meaning |
|-------|---------|
| **differ%** | Share of elements **not** bit-identical (often ~100% from HTP rounding) |
| **meanRel%** | Mean relative error % — main “how different” number |
| **maxRel%** / **max\|Δ\|** / **RMSE** | Worst-case / absolute error |
| **cos** | Cosine similarity (≈1 = same behaviour) |
| **samples** | First 6 elems: `idx \| go \| npu \| \|Δ\|` |
| **worst** | Index of max `\|go−npu\|` with both values |

**Watch:** `swiglu` NPU uses **sigmoid** gate (Go uses **SiLU**) — row `swiglu~` compares against sigmoid-gate ref. `mha` / `cnn2` show real compose divergence.

## Captured run (2026-10-23 · welvet.qnn 1.0.0 · Snapdragon X)

### Summary

| Layer | meanRel% | cos | max\|Δ\| | Verdict |
|-------|----------|-----|---------|---------|
| dense | 0.203% | 1.0000 | 2.33e-05 | match (HTP rounding) |
| softmax | 0.117% | 1.0000 | 4e-05 | match |
| rmsnorm | 0.029% | 1.0000 | 0.00158 | match |
| layernorm | 0.033% | 1.0000 | 0.00134 | match |
| embedding | 0.018% | 1.0000 | 1.53e-05 | match |
| residual | 0.111% | 1.0000 | 0.000467 | match |
| **swiglu** (vs Go SiLU) | 99.8% | −0.06 | 0.00082 | **semantic mismatch** |
| **swiglu~** (vs sigmoid ref) | 0.218% | 1.0000 | 6.46e-07 | match NPU semantics |
| cnn2 | 100% | −0.01 | 0.0253 | **diverge** |
| mha | 55.4% | 0.69 | 0.0453 | **partial diverge** |

### Output samples (go vs npu)

```
dense      B=8 in=128 out=128            differ=100.00%  meanRel=  0.203%  max|Δ|=2.33e-05   cos= 1.0000
           samples (idx | go | npu | |Δ|):
             [   0]  -0.0257031  -0.0257123   9.21e-06
             [   1]  -0.0144296  -0.0144264   3.24e-06
             [   2]   0.00831063   0.00831797   7.34e-06
             [   3]   0.00191715   0.00191629   8.55e-07
             [   4]  -0.0368897  -0.0368992   9.45e-06
             [   5]   0.0157839   0.0157781   5.81e-06
           worst      [ 519]  go=-0.0738468  npu=-0.0738701  |Δ|=2.33e-05

softmax    rows=16 cols=256              differ=100.00%  meanRel=  0.117%  max|Δ|=4e-05      cos= 1.0000
           samples (idx | go | npu | |Δ|):
             [   0]   0.000827504   0.000828266   7.62e-07
             [   1]   0.000470302   0.000470161   1.4e-07
             [   2]   0.00124108   0.00124359   2.52e-06
             [   3]   0.0138482   0.0138474   8.88e-07
             [   4]   0.00118504   0.00118637   1.33e-06
             [   5]   0.00202478   0.0020256   8.2e-07
           worst      [2396]  go= 0.0152911  npu= 0.0152512  |Δ|=4e-05

rmsnorm    B=4 S=32 D=128                differ= 99.99%  meanRel=  0.029%  max|Δ|=0.00158    cos= 1.0000
           samples (idx | go | npu | |Δ|):
             [   0]   1.06994   1.07031   0.000376
             [   1]   0.0690857   0.0690918   6.09e-06
             [   2]   1.68807   1.68848   0.000403
             [   3]   0.340685   0.340576   0.000108
             [   4]  -0.0726365  -0.0726318   4.64e-06
             [   5]   0.705206   0.705566   0.00036
           worst      [6192]  go=-1.74354  npu=-1.74512  |Δ|=0.00158

layernorm  B=4 S=32 D=128                differ= 99.99%  meanRel=  0.033%  max|Δ|=0.00134    cos= 1.0000
           samples (idx | go | npu | |Δ|):
             [   0]  -0.739011  -0.739258   0.000247
             [   1]   1.02725   1.02734   9.72e-05
             [   2]   0.205923   0.205811   0.000112
             [   3]  -1.03157  -1.03125   0.000316
             [   4]   0.84344   0.84375   0.00031
             [   5]   1.21961   1.21973   0.000121
           worst      [11802]  go=-2.11657  npu=-2.11523  |Δ|=0.00134

cnn2       N=2 16x16 cin=8 cout=8 k=1x1  differ=100.00%  meanRel=100.253%  max|Δ|=0.0253     cos=-0.0099
           note: Go ref = NHWC 1×1 GEMV (same as w2a)
           samples (idx | go | npu | |Δ|):
             [   0]  -0.00375825   0.00283241   0.00659
             [   1]   0.0011962  -0.00100708   0.0022
             [   2]   0.00937116   0.00435638   0.00501
             [   3]  -0.00334132  -0.00899506   0.00565
             [   4]   0.00101026   0.00152206   0.000512
             [   5]  -0.000229072   0.00465012   0.00488
           worst      [  98]  go= 0.0145503  npu=-0.0107117  |Δ|=0.0253

embedding  vocab=512 dim=64 n=64         differ=100.00%  meanRel=  0.018%  max|Δ|=1.53e-05   cos= 1.0000
           samples (idx | go | npu | |Δ|):
             [   0]  -0.0445252  -0.0445252   1.86e-08
             [   1]   0.00744917   0.0074501   9.3e-07
             [   2]  -0.0465575  -0.0465698   1.23e-05
             [   3]  -0.0276137  -0.0276184   4.66e-06
             [   4]  -0.0106374  -0.0106354   2.03e-06
             [   5]   0.0243686   0.0243683   2.72e-07
           worst      [2333]  go=-0.0337067  npu=-0.0337219  |Δ|=1.53e-05

residual   n=8192                        differ=100.00%  meanRel=  0.111%  max|Δ|=0.000467   cos= 1.0000
           samples (idx | go | npu | |Δ|):
             [   0]  -0.21892  -0.218994   7.42e-05
             [   1]   0.38021   0.380371   0.000161
             [   2]  -0.295891  -0.295898   7.03e-06
             [   3]   0.441031   0.440918   0.000113
             [   4]  -0.0861777  -0.0862427   6.49e-05
             [   5]   0.195778   0.195801   2.27e-05
           worst      [4316]  go=-0.803244  npu=-0.803711  |Δ|=0.000467

swiglu     B=2 in=64 hidden=128          differ=100.00%  meanRel= 99.775%  max|Δ|=0.00082    cos=-0.0628
           note: vs Go SiLU (NPU uses sigmoid gate)
           samples (idx | go | npu | |Δ|):
             [   0]  -1.06657e-06  -0.000138283   0.000137
             [   1]   1.21066e-07  -2.59876e-05   2.61e-05
             [   2]   3.66235e-06   0.00057888   0.000575
             [   3]   3.09861e-06   0.00028944   0.000286
             [   4]  -1.21631e-06  -0.000453472   0.000452
             [   5]  -2.427e-06  -0.000118732   0.000116
           worst      [ 107]  go= 3.05405e-07  npu=-0.00082016  |Δ|=0.00082

swiglu~    B=2 in=64 hidden=128          differ=100.00%  meanRel=  0.218%  max|Δ|=6.46e-07   cos= 1.0000
           note: vs sigmoid-gate ref (NPU semantics)
           samples (idx | go | npu | |Δ|):
             [   0]  -0.00013786  -0.000138283   4.22e-07
             [   1]  -2.61557e-05  -2.59876e-05   1.68e-07
             [   2]   0.00057912   0.00057888   2.39e-07
             [   3]   0.000289276   0.00028944   1.64e-07
             [   4]  -0.000453723  -0.000453472   2.51e-07
             [   5]  -0.000118805  -0.000118732   7.28e-08
           worst      [  42]  go= 0.000308921  npu= 0.000308275  |Δ|=6.46e-07

mha        B=1 S=8 D=32 H=4              differ=100.00%  meanRel= 55.381%  max|Δ|=0.0453     cos= 0.6945
           note: vs Go causal MHA (NPU compose may differ on mask/scale)
           samples (idx | go | npu | |Δ|):
             [   0]   0.0396136   0.00621796   0.0334
             [   1]  -0.00832848  -0.0040741   0.00425
             [   2]  -0.0265939   0.00621796   0.0328
             [   3]   0.037809   0.0055275   0.0323
             [   4]   0.0325191   0.0200043   0.0125
             [   5]  -0.0132462  -0.00494766   0.0083
           worst      [  15]  go=-0.033517  npu= 0.0118179  |Δ|=0.0453
```

Full filtered log also in [`clean_run.txt`](clean_run.txt).

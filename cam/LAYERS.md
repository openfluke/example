# Layers as cams — when / where / why

Any Op Parallel hosts can be a **cam**. Prefer **identical twins** + `CombineAvg`
for sync experiments. Use **concat** for heterogeneous widths.

| Layer | Good as cam when… | Tip |
|-------|-------------------|-----|
| **Dense** | Default / mixcam | Fastest for BranchModes sweeps |
| **CNN1/2/3** | Spatial senses | Sync via `Proj` Dense; put CNN *inside* Parallel |
| **ConvT1/2/3** | Generative / upsample twins | Same |
| **MHA** | Sequence attention hemispheres | Same `DModel` |
| **Mamba / GDN** | SSM / gated ΔNet twins | Seq input `[B,T,D]` |
| **RNN / LSTM** | Temporal cams | Same Hidden×Seq feat |
| **Embedding** | Dual vocab views | Token IDs in; concat or avg emb |
| **SwiGLU** | FFN twins | Width = InputDim |
| **RMSNorm / LayerNorm** | Rare as solo cams | Usually *inside* a Sequential cam |
| **Softmax** | Probability cams | Avg of distributions |
| **Sequential / Residual** | Deep hemisphere | Whole stack = one mind |
| **KMeans** | Prototype cams | Same K |
| **Metacognition** | Observer twin | Dim-matched |

See [`05_layers`](05_layers/) for a runnable twin of each.

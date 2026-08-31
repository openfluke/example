# 25. ConvTranspose family — convt1 / convt2 / convt3

**When:** upsampling / decoder paths (1D, 2D, 3D).  
**Where:** `layers/convt1`, `layers/convt2`, `layers/convt3`  
**Why:** generative twins and U-Net-style expands; same Proj/Dense backend story as CNN.

```bash
cd 25-convt && source ../env.sh && go run .
```

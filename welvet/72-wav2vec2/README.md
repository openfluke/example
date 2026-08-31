# 72. model/wav2vec2 — CTC ASR

**When:** speech to text (greedy CTC) from 16 kHz mono WAV/PCM.  
**Where:** `github.com/openfluke/welvet/model/wav2vec2`  
**Why:** run `facebook/wav2vec2-base-960h` in pure Go.

Always proves Config + Vocab + LoadHFDir error path. Full ASR needs HF snapshot via `WAV2VEC2_DIR` / `WAV2VEC2_WAV`.

```bash
cd 72-wav2vec2 && source ../env.sh && go run .
```

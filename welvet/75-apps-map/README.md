# 75. Apps map — when / where / why

**When:** you need a product shell (chat, TTS, ASR, image) rather than a library Op.  
**Where:** `welvet/apps/<name>` — each is its **own Go module**.  
**Why:** engine packages never import apps; apps replace → welvet.

This chapter asserts every known app directory exists and prints guidance.  
It does **not** download model weights or start interactive CLIs.

```bash
cd 75-apps-map && source ../env.sh && go run .
```

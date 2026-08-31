# 41. model/sampling

**Part:** VI · Model IO  
**Package:** `github.com/openfluke/welvet/model/sampling`  
**Status:** ok — ✅

## When

Checkpoints, HF import, tokenize, sample, or decode (`model/sampling`).

## Where

`import "github.com/openfluke/welvet/model/sampling"`

```bash
cd 41-sampling && source ../env.sh && go run .
```

## Why

Logits → token ID needs ArgMax, TopK+temperature, penalties, and chat hygiene in one place.

## What

ArgMax, SampleTopK, ApplyRepetitionPenalty, BanIDs, SanitizeChatReply.

## Sample output (captured)

```
1
1
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/41-sampling`).

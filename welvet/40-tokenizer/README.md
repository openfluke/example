# 40. model/tokenizer

**Part:** VI · Model IO  
**Package:** `github.com/openfluke/welvet/model/tokenizer`  
**Status:** ok — ✅

## When

Checkpoints, HF import, tokenize, sample, or decode (`model/tokenizer`).

## Where

`import "github.com/openfluke/welvet/model/tokenizer"`

```bash
cd 40-tokenizer && source ../env.sh && go run .
```

## Why

Generate needs encode/decode of HF tokenizer.json without pulling Python.

## What

LoadTokenizer, Encode/Decode, LoadForEntity.

## Sample output (captured)

```
failed to read tokenizer file: open tokenizer.json: no such file or directory
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/40-tokenizer`).

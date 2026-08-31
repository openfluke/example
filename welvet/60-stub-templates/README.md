# 60. stub/templates

**Part:** VIII · Stubs  
**Package:** `github.com/openfluke/welvet/stub/templates`  
**Status:** partial — 🚧

## When

Scaffold / experimental stub `stub/templates` (may be partial).

## Where

`import "github.com/openfluke/welvet/stub/templates"`

```bash
cd 60-stub-templates && source ../env.sh && go run .
```

## Why

Chat prompts must match model families (ChatML, Llama3, BitNet) without app-specific string glue.

## What

Template.BuildPrompt; presets ChatML, Llama3, BitNetInstruction, MicrosoftBitNetChat.

## Sample output (captured)

```
<|im_start|>system
You are helpful.
<|im_start|>user
Say hi
<|im_start|>assistant
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/60-stub-templates`).

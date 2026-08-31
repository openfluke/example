# 43. apps — octo · flux2 · mosstts

**Part:** VII · Apps  
**Package:** `github.com/openfluke/welvet/apps/…`  
**Status:** partial — 🚧

## When

You need the `apps/…` foundation package.

## Where

`import "github.com/openfluke/welvet/apps/…"`

```bash
cd 43-apps && source ../env.sh && go run .
```

## Why

Products must not pollute engine packages. Octo is the model shell; flux2/mosstts are domain apps. Lucy races (AAI test41 / test48 / test50) are benches, not Welvet packages.

## What

octo (own module): download/convert/chat, see §44. flux2: MMDiT image. mosstts: Speak/SpeakToFile pipeline. AAI sandwiches race TrainMode on toys — see §68. Off the v1.0 engine board.

## Sample output (captured)

```
Apps import the engine — the engine never imports apps.
  cd apps/octo && go run .
  cd apps/flux2 && go run .
  cd apps/mosstts && go run .
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/43-apps`).

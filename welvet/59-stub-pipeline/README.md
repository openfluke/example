# 59. stub/pipeline

**Part:** VIII · Stubs  
**Package:** `github.com/openfluke/welvet/stub/pipeline`  
**Status:** partial — 🚧

## When

Scaffold / experimental stub `stub/pipeline` (may be partial).

## Where

`import "github.com/openfluke/welvet/stub/pipeline"`

```bash
cd 59-stub-pipeline && source ../env.sh && go run .
```

## Why

Decoder wavefront stats helpers — not a full Lucy-style pipeline runner yet.

## What

PipelineForwardStats, TokenTimelineSummary.

## Sample output (captured)

```
{PipelineTicks:0 SubLayerOps:0 MaxActiveJobs:0 MaxBlockSpread:0 MaxDistinctBlocks:0 MaxPendingTokens:0 StallFallback:false TokenDoneTick:[]}
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/59-stub-pipeline`).

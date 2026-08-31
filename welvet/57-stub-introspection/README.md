# 57. stub/introspection

**Part:** VIII · Stubs  
**Package:** `github.com/openfluke/welvet/stub/introspection`  
**Status:** partial — 🚧

## When

Scaffold / experimental stub `stub/introspection` (may be partial).

## Where

`import "github.com/openfluke/welvet/stub/introspection"`

```bash
cd 57-stub-introspection && source ../env.sh && go run .
```

## Why

UIs and FFI need to list Grid methods without hardcoding every export.

## What

GetMethods, GetMethodsJSON, GetMethodSignature.

## Sample output (captured)

```
9 <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/57-stub-introspection`).

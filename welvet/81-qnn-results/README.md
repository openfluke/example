# 81. QNN — go_v_qnn RESULTS (large shapes)

**Part:** XI · Hexagon NPU (welvet.qnn)  
**Package:** `w2a/suites/go_v_qnn`  
**Status:** ok — ✅ matrix published

## When

You want the honest Snapdragon X numbers (not toy shapes).

## Where

```bash
cd 81-qnn-results && source ../env.sh && go run .
```

Full tables: [`w2a/suites/go_v_qnn/RESULTS.md`](../../../w2a/suites/go_v_qnn/RESULTS.md).

## Why

Toy shapes lie. Dense prepare-once wins ~93–98% vs Go/native on **DLL exec**. Online layers pay compose cost on the **call** timer.

## What

Prints a RESULTS.md headline snapshot (no DLL required). Re-run matrix:

```bash
cd w2a && go test ./tests/go_v_qnn -v -count=1 -run TestGoVQNNMatrix -timeout 180m
```

See [feature book §81](https://openfluke.github.io/welvet/chapters/81-qnn-results.html).

## Source

Mirrored for the [Welvet feature book](https://openfluke.github.io/welvet/) (`openfluke.github.io` §81).

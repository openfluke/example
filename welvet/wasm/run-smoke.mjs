#!/usr/bin/env node
/**
 * Welvet book WASM smoke — mirrors Go chapters via @openfluke/welvet.
 * Usage (from anywhere):
 *   node /path/to/example/welvet/wasm/run-smoke.mjs
 * Resolves WASM package from WELVET_TS or sibling welvet checkout.
 */
import { createRequire } from "module";
import { pathToFileURL } from "url";
import path from "path";
import fs from "fs";

const require = createRequire(import.meta.url);

function resolveWelvetTS() {
  if (process.env.WELVET_TS) return process.env.WELVET_TS;
  const candidates = [
    path.resolve(path.dirname(new URL(import.meta.url).pathname), "../../../../chaosglue/welvet/apps/w2a/typescript"),
    path.resolve(path.dirname(new URL(import.meta.url).pathname), "../../../../welvet/apps/w2a/typescript"),
    "/home/openfluke/git/welvet/apps/w2a/typescript",
    "/home/openfluke/git/chaosglue/welvet/apps/w2a/typescript",
  ];
  for (const c of candidates) {
    if (fs.existsSync(path.join(c, "dist", "index.js"))) return c;
    if (fs.existsSync(path.join(c, "src", "index.ts"))) return c;
  }
  throw new Error("Set WELVET_TS to apps/w2a/typescript (run npm run build:all there first)");
}

const root = resolveWelvetTS();
const dist = path.join(root, "dist", "index.js");
const mod = await import(pathToFileURL(dist).href);

await mod.init();
mod.assertEngineVersion();

// 01-welvet / 11-dense style
const g = mod.createGrid({ depth: 1, rows: 1, cols: 1, layers_per_cell: 1 });
g.placeDense(JSON.stringify({ in: 8, out: 4, act: "relu", dtype: 1 }));
const x = new Float32Array(8);
for (let i = 0; i < 8; i++) x[i] = i * 0.1;
const fwd = g.forward(x);
if (fwd.error) throw new Error(fwd.error);
console.log("01/11 dense forward", fwd.output.length);

// 31-training
const tgt = new Float32Array(4);
tgt[0] = 1;
const tr = g.trainSGD(x, tgt, 0.05);
console.log("31 trainSGD loss", tr.loss);

// 33-dna
console.log("33 dna bytes", g.extractDNA().length);

// 27-parallel / cameral
const st = mod.createBicameral({ in: 4, hidden: 4, out: 4 });
const modes = mod.listNamedTrainModes();
const r = st.trainStackMSE(new Float32Array(4).fill(0.1), new Float32Array([1, 0, 0, 0]), modes[0], 0.05);
console.log("27/cam train", modes[0], r.loss);

// 36-tanhi (disabled — no panic)
g.configureTanhi(JSON.stringify({ Enabled: false }));
console.log("36 tanhi configure ok");

console.log("=== example welvet WASM smoke OK ===");

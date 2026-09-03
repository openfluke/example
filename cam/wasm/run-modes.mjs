#!/usr/bin/env node
/**
 * Cam cookbook WASM — modes / CamSync / CamKit smoke.
 */
import path from "path";
import fs from "fs";
import { pathToFileURL } from "url";

function resolveWelvetTS() {
  if (process.env.WELVET_TS) return process.env.WELVET_TS;
  const candidates = [
    "/home/openfluke/git/welvet/apps/w2a/typescript",
    "/home/openfluke/git/chaosglue/welvet/apps/w2a/typescript",
  ];
  for (const c of candidates) {
    if (fs.existsSync(path.join(c, "dist", "index.js"))) return c;
  }
  throw new Error("Set WELVET_TS");
}

const mod = await import(pathToFileURL(path.join(resolveWelvetTS(), "dist", "index.js")).href);
await mod.init();
mod.assertEngineVersion();

const modes = mod.listNamedTrainModes();
console.log("01_modes count", modes.length);

const st = mod.createBicameral({ in: 4, hidden: 4, out: 4 });
const inp = new Float32Array([0.1, 0.2, 0.3, 0.4]);
const tgt = new Float32Array([1, 0, 0, 0]);
for (const m of modes.slice(0, 5)) {
  const s = mod.createBicameral({ in: 4, hidden: 4, out: 4 });
  const r = s.trainStackMSE(inp, tgt, m, 0.05);
  if (r.error) throw new Error(m + ": " + r.error);
  console.log("  mode", m, "loss", r.loss);
  s.free();
}

const hem = mod.createHemispheres({ dim: 4, n: 2, combine: "add" });
hem.setCamSync(JSON.stringify({ Enabled: true, Alpha: 1 }));
hem.setCamKit(JSON.stringify({ ShadowCoef: 1, DNAReg: 0 }));
const pr = hem.trainMSE(inp, tgt, modes[0], 0.05);
if (pr.error) throw new Error(pr.error);
console.log("03_camsync/04_kit", pr.loss);

console.log("=== cam WASM smoke OK ===");

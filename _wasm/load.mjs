#!/usr/bin/env node
import { createRequire } from "node:module";
import { existsSync } from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const EXAMPLE_ROOT = path.resolve(__dirname, "..");

export function welvetPackageRoot() {
  if (process.env.WELVET_TS) return path.resolve(process.env.WELVET_TS);
  const sibling = path.resolve(
    EXAMPLE_ROOT,
    "../welvet/apps/w2a/typescript",
  );
  if (existsSync(path.join(sibling, "package.json"))) return sibling;
  const chaos = path.resolve(
    EXAMPLE_ROOT,
    "../chaosglue/welvet/apps/w2a/typescript",
  );
  if (existsSync(path.join(chaos, "package.json"))) return chaos;
  try {
    const req = createRequire(import.meta.url);
    return path.dirname(req.resolve("@openfluke/welvet/package.json"));
  } catch {
    throw new Error(
      "Set WELVET_TS to apps/w2a/typescript, or npm link @openfluke/welvet",
    );
  }
}

export async function loadWelvet() {
  const root = welvetPackageRoot();
  const entry = path.join(root, "dist/index.js");
  if (!existsSync(entry)) {
    throw new Error(`Build welvet first: cd ${root} && npm run build:all`);
  }
  const mod = await import(entry);
  await mod.init({ simd: false });
  return mod;
}

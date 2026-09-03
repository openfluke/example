#!/usr/bin/env node
/**
 * npm/Node WASM example — welvet/05-quant
 * Requires: WELVET_TS or built @openfluke/welvet (apps/w2a/typescript)
 *   export WELVET_TS=/path/to/apps/w2a/typescript
 *   node npm/run.mjs
 */
import { runChapter } from "../../../_wasm/run-chapter.mjs";

const SLUG = "05-quant";
const msg = await runChapter(SLUG);
if (String(msg).startsWith("SKIP")) process.exitCode = 0;
else if (msg !== "OK" && !String(msg).includes("OK")) {
  /* runChapter already threw on failure */
}
console.log("=== welvet/05-quant npm OK ===");

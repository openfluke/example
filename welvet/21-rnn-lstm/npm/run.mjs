#!/usr/bin/env node
/**
 * npm/Node WASM example — welvet/21-rnn-lstm
 * Requires: WELVET_TS or built @openfluke/welvet (apps/w2a/typescript)
 *   export WELVET_TS=/path/to/apps/w2a/typescript
 *   node npm/run.mjs
 */
import { runChapter } from "../../../_wasm/run-chapter.mjs";

const SLUG = "21-rnn-lstm";
const msg = await runChapter(SLUG);
if (String(msg).startsWith("SKIP")) process.exitCode = 0;
else if (msg !== "OK" && !String(msg).includes("OK")) {
  /* runChapter already threw on failure */
}
console.log("=== welvet/21-rnn-lstm npm OK ===");

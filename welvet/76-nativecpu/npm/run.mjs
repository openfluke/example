#!/usr/bin/env node
import { runChapter } from "../../../_wasm/run-chapter.mjs";
const SLUG = "76-nativecpu";
const msg = await runChapter(SLUG);
if (String(msg).startsWith("SKIP")) process.exitCode = 0;
console.log("=== welvet/76-nativecpu npm OK ===");

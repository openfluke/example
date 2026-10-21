#!/usr/bin/env node
import { runChapter } from "../../../_wasm/run-chapter.mjs";
const SLUG = "77-native-layers";
const msg = await runChapter(SLUG);
if (String(msg).startsWith("SKIP")) process.exitCode = 0;
console.log("=== welvet/77-native-layers npm OK ===");

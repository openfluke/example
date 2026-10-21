#!/usr/bin/env node
import { runChapter } from "../../../_wasm/run-chapter.mjs";
const SLUG = "78-native-trainmodes";
const msg = await runChapter(SLUG);
if (String(msg).startsWith("SKIP")) process.exitCode = 0;
console.log("=== welvet/78-native-trainmodes npm OK ===");

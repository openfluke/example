#!/usr/bin/env node
import { loadWelvet } from "./load.mjs";
import { runDemo } from "./demos.mjs";

export async function runChapter(slug, { quiet } = {}) {
  const log = quiet ? () => {} : console.log;
  const mod = await loadWelvet();
  const msg = await runDemo(slug, mod, log);
  return msg;
}

if (import.meta.url === `file://${process.argv[1]}`) {
  const slug = process.argv[2];
  if (!slug) {
    console.error("usage: run-chapter.mjs <slug>");
    process.exit(2);
  }
  runChapter(slug)
    .then((m) => {
      console.log(m);
    })
    .catch((e) => {
      console.error(e);
      process.exit(1);
    });
}

#!/usr/bin/env node
/**
 * Generate npm/ + html/ for every welvet chapter and cam cookbook entry.
 */
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(__dirname, "..");

function npmRunner(slug, kind) {
  const depth = kind === "cam" ? "../../.." : "../../..";
  return `#!/usr/bin/env node
/**
 * npm/Node WASM example — ${kind}/${slug}
 * Requires: WELVET_TS or built @openfluke/welvet (apps/w2a/typescript)
 *   export WELVET_TS=/path/to/apps/w2a/typescript
 *   node npm/run.mjs
 */
import { runChapter } from "${depth}/_wasm/run-chapter.mjs";

const SLUG = "${slug}";
const msg = await runChapter(SLUG);
if (String(msg).startsWith("SKIP")) process.exitCode = 0;
else if (msg !== "OK" && !String(msg).includes("OK")) {
  /* runChapter already threw on failure */
}
console.log("=== ${kind}/${slug} npm OK ===");
`;
}

function htmlPage(slug, kind, title) {
  const relAssets = kind === "cam" ? "../../../_wasm" : "../../../_wasm";
  return `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>${title} — WASM</title>
  <style>
    :root { color-scheme: light dark; font-family: ui-sans-serif, system-ui, sans-serif; }
    body { margin: 1.5rem; max-width: 52rem; line-height: 1.45; }
    h1 { font-size: 1.25rem; font-weight: 650; }
    #out pre { margin: 0.35rem 0; padding: 0.5rem 0.65rem; background: color-mix(in oklab, Canvas 92%, CanvasText 8%); border-radius: 6px; overflow-x: auto; font-size: 0.85rem; }
    .meta { opacity: 0.7; font-size: 0.9rem; }
  </style>
</head>
<body>
  <h1 id="title">${kind}/${slug}</h1>
  <p class="meta">${title}. Serve from the <code>example</code> repo root so <code>/_wasm/assets/main.wasm</code> resolves (see <code>_wasm/README.md</code>).</p>
  <div id="out"></div>
  <script>window.CHAPTER = ${JSON.stringify(slug)};</script>
  <script src="${relAssets}/assets/wasm_exec.js"></script>
  <script type="module" src="${relAssets}/browser-runner.js"></script>
</body>
</html>
`;
}

function packageJson(slug, kind) {
  return JSON.stringify(
    {
      name: `@openfluke/example-${kind}-${slug}`,
      private: true,
      type: "module",
      scripts: {
        start: "node run.mjs",
        test: "node run.mjs",
      },
    },
    null,
    2,
  ) + "\n";
}

const chapters = JSON.parse(
  fs.readFileSync(path.join(ROOT, "welvet/_chapters.json"), "utf8"),
);

let n = 0;
for (const ch of chapters) {
  const dir = path.join(ROOT, "welvet", ch.slug);
  const npmDir = path.join(dir, "npm");
  const htmlDir = path.join(dir, "html");
  fs.mkdirSync(npmDir, { recursive: true });
  fs.mkdirSync(htmlDir, { recursive: true });
  fs.writeFileSync(path.join(npmDir, "run.mjs"), npmRunner(ch.slug, "welvet"));
  fs.writeFileSync(path.join(npmDir, "package.json"), packageJson(ch.slug, "welvet"));
  fs.writeFileSync(
    path.join(htmlDir, "index.html"),
    htmlPage(ch.slug, "welvet", ch.title || ch.slug),
  );
  n++;
}

const camTitles = {
  "01_modes": "Cameral train modes",
  "02_combine": "Hemisphere combine",
  "03_camsync": "CamSync",
  "04_kit": "Cam kit",
  "05_layers": "Cameral layers",
  "06_recipes": "Cameral recipes",
};

for (const slug of Object.keys(camTitles)) {
  const dir = path.join(ROOT, "cam", slug);
  const npmDir = path.join(dir, "npm");
  const htmlDir = path.join(dir, "html");
  fs.mkdirSync(npmDir, { recursive: true });
  fs.mkdirSync(htmlDir, { recursive: true });
  fs.writeFileSync(path.join(npmDir, "run.mjs"), npmRunner(slug, "cam"));
  fs.writeFileSync(path.join(npmDir, "package.json"), packageJson(slug, "cam"));
  fs.writeFileSync(
    path.join(htmlDir, "index.html"),
    htmlPage(slug, "cam", camTitles[slug]),
  );
  n++;
}

console.log(`generated npm+html for ${n} examples`);

/**
 * Browser runner: set window.CHAPTER (or ?chapter=) then include this module
 * after wasm_exec.js. Expects assets under /_wasm/assets/ or ?wasmBase=.
 */
import { runDemo, NATIVE_ONLY } from "./demos.mjs";

function resolveBase() {
  const q = new URLSearchParams(location.search);
  if (q.get("wasmBase")) return q.get("wasmBase").replace(/\/$/, "");
  // Served from example root: /_wasm/assets
  return new URL("/_wasm/assets/", location.origin).href.replace(/\/$/, "");
}

async function boot() {
  const out = document.getElementById("out") || document.body;
  const log = (m) => {
    const line = document.createElement("pre");
    line.textContent = String(m);
    out.appendChild(line);
    console.log(m);
  };
  const slug =
    window.CHAPTER ||
    new URLSearchParams(location.search).get("chapter") ||
    "01-welvet";

  document.title = `Welvet WASM — ${slug}`;
  const h = document.getElementById("title");
  if (h) h.textContent = slug;

  if (NATIVE_ONLY.has(slug)) {
    log(`SKIP ${slug} — native-only`);
    return;
  }

  if (typeof Go !== "function") {
    throw new Error("wasm_exec.js not loaded (Go global missing)");
  }

  const base = resolveBase();
  const wasmURL = `${base}/main.wasm`;
  const go = new Go();
  const res = await fetch(wasmURL);
  if (!res.ok) throw new Error(`fetch ${wasmURL}: ${res.status}`);
  const result = await WebAssembly.instantiateStreaming(res, go.importObject);
  go.run(result.instance);

  // Wait for exports
  for (let i = 0; i < 50; i++) {
    if (typeof createGrid === "function" || typeof createWelvetGrid === "function") break;
    await new Promise((r) => setTimeout(r, 20));
  }

  // Build a thin mod facade from globals
  const mod = {
    assertEngineVersion: () => {
      if (typeof welvetEngineVersion === "function") {
        const v = welvetEngineVersion();
        if (v && v !== "1.1.1") throw new Error(`version ${v}`);
      }
    },
    createGrid: (cfg) => {
      const j = typeof cfg === "string" ? cfg : JSON.stringify(cfg);
      return typeof createGrid === "function" ? createGrid(j) : createWelvetGrid(j);
    },
    createWelvetStore: (...a) => createWelvetStore(...a),
    listWelvetDTypes: () => listWelvetDTypes(),
    listWelvetFormats: () => listWelvetFormats(),
    listWelvetBackends: () => listWelvetBackends(),
    listNamedTrainModes: () =>
      typeof listNamedTrainModes === "function"
        ? listNamedTrainModes()
        : listWelvetNamedTrainModes(),
    createBicameral: (c) =>
      typeof createBicameral === "function"
        ? createBicameral(typeof c === "string" ? c : JSON.stringify(c))
        : createWelvetBicameral(JSON.stringify(c)),
    createHemispheres: (c) =>
      typeof createHemispheres === "function"
        ? createHemispheres(typeof c === "string" ? c : JSON.stringify(c))
        : createWelvetHemispheres(JSON.stringify(c)),
    SeedFrom: (...a) => SeedFrom(...a),
    FountainLTRoundTrip: (...a) => FountainLTRoundTrip(...a),
    MemoryFromGrid: (...a) => MemoryFromGrid(...a),
    ReleaseTransient: (...a) => ReleaseTransient(...a),
    GraftGrids: (...a) => GraftGrids(...a),
    EnsembleMajorityVote: (...a) => EnsembleMajorityVote(...a),
    EvaluatePrediction: (...a) => EvaluatePrediction(...a),
    IntrospectGrid: (...a) => IntrospectGrid(...a),
    TemplateBuildPrompt: (...a) => TemplateBuildPrompt(...a),
    LucyAvailability: (...a) => LucyAvailability(...a),
    LucyScore: (...a) => LucyScore(...a),
    ArgMax: (...a) => ArgMax(...a),
    SampleTopK: (...a) => SampleTopK(...a),
    BuildTransformerPrompt: (...a) => BuildTransformerPrompt(...a),
    ExtractDNA: (...a) => ExtractDNA(...a),
    SpliceDNA: (...a) => SpliceDNA(...a),
    DeserializeEntity: (...a) =>
      typeof DeserializeEntity === "function"
        ? DeserializeEntity(...a)
        : DeserializeGrid(...a),
    DeserializeGrid: (...a) => DeserializeGrid(...a),
    createWelvetStepState: (...a) => createWelvetStepState(...a),
  };

  try {
    const msg = await runDemo(slug, mod, log);
    log(`RESULT: ${msg}`);
  } catch (e) {
    log(`FAIL: ${e?.message || e}`);
    throw e;
  }
}

boot().catch((e) => {
  console.error(e);
  const out = document.getElementById("out") || document.body;
  const line = document.createElement("pre");
  line.style.color = "crimson";
  line.textContent = String(e?.stack || e);
  out.appendChild(line);
});

#!/usr/bin/env node
/**
 * Shared demo bodies for every Welvet book chapter + cam cookbook.
 * Node: pass the @openfluke/welvet module after init().
 * Browser: pass a facade; bare WASM globals are also used as fallbacks.
 */
export const NATIVE_ONLY = new Set([
  "06-simd",
  "07-webgpu",
  "10-fusedgpu",
  "48-stub-donate",
  "50-stub-hardware",
  "51-stub-accel",
  "39-hf",
  "61-stub-universal",
]);

const g = () => globalThis;

function fill(n, v = 0.1) {
  const a = new Float32Array(n);
  for (let i = 0; i < n; i++) a[i] = v * ((i % 7) + 1);
  return a;
}

function ok(r, label) {
  if (r && typeof r === "object" && r.error) throw new Error(`${label}: ${r.error}`);
  return r;
}

function call(mod, name, ...args) {
  if (mod && typeof mod[name] === "function") return mod[name](...args);
  if (typeof g()[name] === "function") return g()[name](...args);
  throw new Error(`missing ${name}`);
}

function grid(mod) {
  return call(mod, "createGrid", { depth: 1, rows: 1, cols: 1, layers_per_cell: 1 });
}

function place(gr, method, spec) {
  const fn = gr[method];
  if (typeof fn !== "function") throw new Error(`missing ${method}`);
  const payload = typeof spec === "string" ? spec : JSON.stringify(spec);
  return ok(fn.call(gr, payload), method);
}

/** @returns {Promise<string>} */
export async function runDemo(slug, mod, log = console.log) {
  if (NATIVE_ONLY.has(slug)) {
    const msg = `SKIP ${slug} — native-only (host SIMD/WebGPU/FS/TCP)`;
    log(msg);
    return msg;
  }

  if (mod?.assertEngineVersion) mod.assertEngineVersion();
  else if (typeof g().welvetEngineVersion === "function") {
    const v = g().welvetEngineVersion();
    if (v && v !== "1.1.1") throw new Error(`version ${v}`);
  }

  const layerPlace = {
    "11-dense": { method: "placeDense", spec: { in: 8, out: 4, act: "relu", dtype: 1 }, inLen: 8 },
    "12-mha": { method: "placeMHA", spec: { DModel: 8, NumHeads: 2, SeqLen: 4, dtype: 1 }, inLen: 32, shape: [1, 4, 8] },
    "13-swiglu": { method: "placeSwiGLU", spec: { InputDim: 8, IntermediateDim: 16, dtype: 1 }, inLen: 8 },
    "14-rmsnorm": { method: "placeRMSNorm", spec: { dim: 8, dtype: 1 }, inLen: 8 },
    "15-layernorm": { method: "placeLayerNorm", spec: { dim: 8, dtype: 1 }, inLen: 8 },
    "16-embedding": { method: "placeEmbedding", spec: { VocabSize: 32, EmbeddingDim: 8, SeqLen: 4, dtype: 1 }, inLen: 4, shape: [1, 4] },
    "17-softmax": { method: "placeSoftmax", spec: { dim: 8, dtype: 1 }, inLen: 8 },
    "18-sequential": { method: "placeSequential", spec: { dim: 8, Depth: 2, dtype: 1 }, inLen: 8 },
    "19-residual": { method: "placeResidual", spec: { dim: 8, Depth: 1, dtype: 1 }, inLen: 8 },
    "20-cnn": { method: "placeCNN1", spec: { InChannels: 1, Filters: 4, SeqLen: 16, Kernel: 3, dtype: 1 }, inLen: 16, shape: [1, 1, 16] },
    "21-rnn-lstm": { method: "placeLSTM", spec: { InputSize: 8, HiddenSize: 8, SeqLen: 4, dtype: 1 }, inLen: 32, shape: [1, 4, 8] },
    "23-gdn": { method: "placeGDN", spec: { HiddenSize: 8, NumKeyHeads: 2, NumValueHeads: 2, KeyHeadDim: 4, ValueHeadDim: 4, ConvKernel: 3 }, inLen: 8, shape: [1, 1, 8] },
    "24-mamba": { method: "placeMamba", spec: { DModel: 8, DState: 8, SeqLen: 4, dtype: 1 }, inLen: 32, shape: [1, 4, 8] },
    "25-convt": { method: "placeConvT1", spec: { InChannels: 4, Filters: 2, SeqLen: 8, Kernel: 3, dtype: 1 }, inLen: 32, shape: [1, 4, 8] },
    "26-kmeans": { method: "placeKMeans", spec: { FeatureDim: 8, NumClusters: 4, dtype: 1 }, inLen: 8 },
    "27-parallel": { method: "placeParallel", spec: { dim: 8, OutFeat: 8, Branches: 2, dtype: 1 }, inLen: 8 },
    "28-metacognition": { method: "placeMetacognition", spec: { Dim: 8, dtype: 1 }, inLen: 8 },
  };

  if (layerPlace[slug]) {
    const def = layerPlace[slug];
    const gr = grid(mod);
    place(gr, def.method, def.spec);
    const x = fill(def.inLen);
    const fwd = def.shape ? gr.forward(x, JSON.stringify(def.shape)) : gr.forward(x);
    ok(fwd, "forward");
    log(`${slug} forward out=${fwd.output?.length}`);
    gr.free?.();
    return "OK";
  }

  switch (slug) {
    case "01-welvet":
    case "02-tree":
    case "09-architecture":
    case "29-forward":
    case "62-w2a":
    case "63-validation":
    case "64-scorecard":
    case "71-dispatch":
    case "73-all-layers":
    case "74-kokoro":
    case "75-apps-map":
    case "43-apps":
    case "44-octo":
    case "08-tiling":
    case "22-seqmix":
    case "72-wav2vec2": {
      const gr = grid(mod);
      place(gr, "placeDense", { in: 8, out: 4, act: "relu", dtype: 1 });
      const fwd = ok(gr.forward(fill(8)), "forward");
      log(`${slug} dense forward ${fwd.output.length}`);
      gr.free?.();
      return "OK";
    }
    case "03-core": {
      const dts = mod?.listDTypes?.() ?? JSON.parse(call(mod, "listWelvetDTypes"));
      const fms = mod?.listFormats?.() ?? JSON.parse(call(mod, "listWelvetFormats"));
      let bes = [];
      try {
        bes = JSON.parse(call(mod, "listWelvetBackends"));
      } catch {
        bes = ["cpu_tiled"];
      }
      log(`dtypes=${dts.length} formats=${fms.length} backends=${bes.length}`);
      return "OK";
    }
    case "04-weights":
    case "05-quant":
    case "65-cross-numeric": {
      const s = mod?.createStore
        ? mod.createStore(8, 8, 1, 0, fill(64))
        : call(mod, "createWelvetStore", 8, 8, 1, 0, fill(64));
      ok(s, "store");
      ok(s.applySGD(Float64Array.from({ length: 64 }, () => 0.01), 0.1), "sgd");
      const flat = s.flattenF32();
      if (flat?.error) throw new Error(flat.error);
      log(`${slug} ApplySGD ok len=${flat.length ?? flat?.byteLength}`);
      s.free?.();
      return "OK";
    }
    case "30-backward":
    case "31-training": {
      const gr = grid(mod);
      place(gr, "placeDense", { in: 4, out: 4, act: "relu", dtype: 1 });
      const x = fill(4);
      ok(gr.forward(x), "fwd");
      const tr = ok(gr.trainSGD(x, fill(4, 0), 0.05), "train");
      log(`${slug} trainSGD loss=${tr.loss}`);
      gr.free?.();
      return "OK";
    }
    case "32-step": {
      const gr = grid(mod);
      place(gr, "placeDense", { in: 4, out: 4, dtype: 1 });
      const st = call(mod, "createWelvetStepState", gr._id);
      ok(st, "step");
      st.setInput(fill(4));
      ok(st.step(false), "step.fwd");
      log("32-step ok");
      st.free?.();
      gr.free?.();
      return "OK";
    }
    case "33-dna":
    case "37-telemetry": {
      const gr = grid(mod);
      place(gr, "placeDense", { in: 4, out: 4, dtype: 1 });
      const dna = gr.extractDNA();
      if (!dna || String(dna).includes('"error"')) throw new Error("dna");
      log(`${slug} dna len=${String(dna).length}`);
      gr.free?.();
      return "OK";
    }
    case "34-evolution": {
      const gr = grid(mod);
      place(gr, "placeDense", { in: 4, out: 4, dtype: 1 });
      const gr2 = grid(mod);
      place(gr2, "placeDense", { in: 4, out: 4, dtype: 1 });
      const out = call(mod, "SpliceDNA", gr._id, gr2._id);
      ok(out, "splice");
      log("34-evolution splice ok");
      out.free?.();
      gr.free?.();
      gr2.free?.();
      return "OK";
    }
    case "35-tween": {
      const gr = grid(mod);
      place(gr, "placeDense", { in: 4, out: 4, dtype: 1 });
      const tr = ok(gr.trainTween(fill(4), fill(4, 0), 0.05), "tween");
      log(`35-tween loss=${tr.loss}`);
      gr.free?.();
      return "OK";
    }
    case "36-tanhi": {
      const gr = grid(mod);
      place(gr, "placeDense", { in: 4, out: 4, dtype: 1 });
      ok(gr.configureTanhi(JSON.stringify({ Enabled: false })), "tanhi");
      log("36-tanhi configure ok");
      gr.free?.();
      return "OK";
    }
    case "38-entity":
    case "46-stub-serialization": {
      const gr = grid(mod);
      place(gr, "placeDense", { in: 4, out: 4, dtype: 1 });
      const ent = gr.serializeEntity();
      if (!(ent instanceof Uint8Array) || ent.length < 4) throw new Error("entity");
      const back = call(mod, "DeserializeGrid", ent);
      ok(back, "deser");
      log(`${slug} entity bytes=${ent.length}`);
      back.free?.();
      gr.free?.();
      return "OK";
    }
    case "40-tokenizer": {
      const i = call(mod, "ArgMax", new Float32Array([0.1, 0.9, 0.2]));
      if (i !== 1) throw new Error("argmax");
      log("40-tokenizer sampling helpers ok");
      return "OK";
    }
    case "41-sampling": {
      const i = call(mod, "SampleTopK", new Float32Array([0.1, 0.9, 0.2]), 2, 1);
      log(`41-sampling topk=${i}`);
      return "OK";
    }
    case "42-transformer": {
      const p = call(mod, "BuildTransformerPrompt", "hi", "sys");
      if (!String(p).includes("hi")) throw new Error("prompt");
      log("42-transformer prompt ok");
      return "OK";
    }
    case "45-stub-seed": {
      const a = mod?.seedFrom
        ? mod.seedFrom("welvet", 42, true)
        : call(mod, "SeedFrom", JSON.stringify(["welvet", 42, true]));
      const b = mod?.seedFrom
        ? mod.seedFrom("welvet", 42, true)
        : call(mod, "SeedFrom", JSON.stringify(["welvet", 42, true]));
      if (String(a) !== String(b)) throw new Error("seed mismatch");
      log(`45-stub-seed ${a}`);
      return "OK";
    }
    case "47-stub-memory": {
      const gr = grid(mod);
      place(gr, "placeDense", { in: 8, out: 8, dtype: 1 });
      const fp = call(mod, "MemoryFromGrid", gr._id);
      log(`47-stub-memory ${fp}`);
      try {
        call(mod, "ReleaseTransient");
      } catch {
        /* optional */
      }
      gr.free?.();
      return "OK";
    }
    case "49-stub-fountain": {
      const rt = call(
        mod,
        "FountainLTRoundTrip",
        [
          new Uint8Array([1, 2, 3, 4, 5, 6, 7, 8]),
          new Uint8Array([9, 10, 11, 12, 13, 14, 15, 16]),
          new Uint8Array([17, 18, 19, 20, 21, 22, 23, 24]),
        ],
        "7",
      );
      ok(rt, "fountain");
      if (!rt.ok) throw new Error("lt not ok");
      log("49-stub-fountain LT ok");
      return "OK";
    }
    case "52-stub-clustering":
    case "53-stub-ensemble":
    case "54-stub-evaluation":
    case "55-stub-grafting":
    case "56-stub-grouping":
    case "57-stub-introspection":
    case "58-stub-observer":
    case "59-stub-pipeline":
    case "60-stub-templates": {
      if (slug === "55-stub-grafting") {
        const g1 = grid(mod);
        place(g1, "placeDense", { in: 4, out: 4, dtype: 1 });
        const g2 = grid(mod);
        place(g2, "placeDense", { in: 4, out: 4, dtype: 1 });
        const p = call(mod, "GraftGrids", [g1._id, g2._id], "concat");
        ok(p, "graft");
        log("55 grafting ok");
        p.free?.();
        g1.free?.();
        g2.free?.();
        return "OK";
      }
      if (slug === "53-stub-ensemble") {
        const v = call(mod, "EnsembleMajorityVote", JSON.stringify([[0, 1], [0, 2], [0, 1]]));
        log(`53 ensemble ${v}`);
        return "OK";
      }
      if (slug === "54-stub-evaluation") {
        const j = call(mod, "EvaluatePrediction", 0, 1, 0.9);
        log(`54 eval ${j}`);
        return "OK";
      }
      if (slug === "57-stub-introspection") {
        const gr = grid(mod);
        place(gr, "placeDense", { in: 4, out: 4, dtype: 1 });
        const m = call(mod, "IntrospectGrid", gr._id);
        log(`57 introspect ${String(m).slice(0, 80)}`);
        gr.free?.();
        return "OK";
      }
      if (slug === "60-stub-templates") {
        const p = call(mod, "TemplateBuildPrompt", "chatml", "hi", "sys");
        if (!String(p).includes("hi")) throw new Error("template");
        log("60 templates ok");
        return "OK";
      }
      const gr = grid(mod);
      place(gr, "placeDense", { in: 4, out: 4, dtype: 1 });
      ok(gr.forward(fill(4)), "fwd");
      log(`${slug} portable smoke ok`);
      gr.free?.();
      return "OK";
    }
    case "66-lucy":
    case "69-lucy-density": {
      const avail = call(mod, "LucyAvailability", 1, 2);
      const score = call(mod, "LucyScore", 10, 50, 0.9);
      log(`${slug} avail=${avail} score=${score}`);
      return "OK";
    }
    case "67-train-modes":
    case "68-cameral":
    case "70-cam-sync": {
      const modes = mod?.listNamedTrainModes?.() ?? JSON.parse(call(mod, "listWelvetNamedTrainModes"));
      const st = call(mod, "createBicameral", { in: 4, hidden: 4, out: 4 });
      ok(st, "bicameral");
      const tr = ok(st.trainStackMSE(fill(4), fill(4, 0), modes[0], 0.05), "train");
      log(`${slug} ${modes[0]} loss=${tr.loss} modes=${modes.length}`);
      if (slug === "70-cam-sync") {
        const hem = call(mod, "createHemispheres", { dim: 4, n: 2, combine: "add" });
        ok(hem.setCamSync(JSON.stringify({ Enabled: true, Alpha: 1 })), "camsync");
        log("70 camsync set ok");
        hem.free?.();
      }
      st.free?.();
      return "OK";
    }
    case "01_modes":
    case "02_combine":
    case "03_camsync":
    case "04_kit":
    case "05_layers":
    case "06_recipes": {
      const modes = mod?.listNamedTrainModes?.() ?? JSON.parse(call(mod, "listWelvetNamedTrainModes"));
      const combine = slug === "02_combine" ? "avg" : "add";
      const p = call(mod, "createHemispheres", { dim: 4, n: 2, combine });
      ok(p, "parallel");
      p.setBranchModes?.(JSON.stringify([modes[0], modes[0]]));
      if (slug === "03_camsync" || slug === "04_kit" || slug === "06_recipes") {
        p.setCamSync?.(JSON.stringify({ Enabled: true, Alpha: 1 }));
        p.setCamKit?.(JSON.stringify({ ShadowCoef: 1, DNAReg: 0, SurpriseThresh: 0 }));
      }
      const tr = ok(p.trainMSE(fill(4), fill(4, 0), modes[0], 0.05), "train");
      log(`cam/${slug} ${modes[0]} loss=${tr.loss}`);
      p.free?.();
      return "OK";
    }
    default: {
      const gr = grid(mod);
      place(gr, "placeDense", { in: 4, out: 4, dtype: 1 });
      ok(gr.forward(fill(4)), "fwd");
      log(`${slug} default dense smoke ok`);
      gr.free?.();
      return "OK";
    }
  }
}

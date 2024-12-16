async function runWasm() {
  const go = new Go();
  const wasmModule = await WebAssembly.instantiateStreaming(
    fetch("./bin/app.wasm"),
    go.importObject
  );

  go.run(wasmModule.instance);
}

runWasm();

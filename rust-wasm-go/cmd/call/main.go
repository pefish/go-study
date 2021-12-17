package main

import (
	"fmt"
	"log"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	wasmBytes, err := Asset("/Users/pefish/Work/rust/add/target/wasm32-unknown-unknown/release/add.wasm")
	if err != nil {
		log.Fatal(err)
	}

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		log.Fatal(err)
	}

	// Gets the `sum` exported function from the WebAssembly instance.
	sum, err := instance.Exports.GetFunction("sum")
	if err != nil {
		log.Fatal(err)
	}

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, err := sum(65464, 37)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result) // 42!
}

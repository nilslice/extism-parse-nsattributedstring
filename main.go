package main

import (
	"context"
	"fmt"
	"os"
	"io"

	extism "github.com/extism/go-sdk"
)

func main() {
	// construct the plugin from this Wasm code:
	// https://modsurfer.dylibso.com/module?hash=08647b42ec879ce597814ba341b5b0c88feda5ae4940a28183a440157ef7dbc2
	// Note: this could also be loaded from disk, etc.
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmUrl{Url: "https://cdn.modsurfer.dylibso.com/api/v1/module/08647b42ec879ce597814ba341b5b0c88feda5ae4940a28183a440157ef7dbc2.wasm"},
		},
	}
	plugin, err := extism.NewPlugin(context.Background(), manifest, extism.PluginConfig{}, nil)
	if err != nil {
		panic(err)
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	// provide the input bytes, expected to be in the NSAttributedString encoding, and return the parsed string as a slice of bytes.
	_, output, err := plugin.Call("parse_nsattributedstring", input)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}

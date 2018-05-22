package main

import (
	"bytes"
	"io/ioutil"
	"fmt"
	"github.com/go-interpreter/wagon/wasm"
)

func main() {
	testWasm()
}

func testWasm() {
	raw, err := ioutil.ReadFile("hello.wasm")
	if err != nil {
		fmt.Println(err)
		return
	}

	r := bytes.NewReader(raw)
	m, err := wasm.ReadModule(r, nil)
	if err != nil {
		fmt.Printf("error reading module %v\n", err)
	}
	if m == nil {
		fmt.Println("error reading module: (nil *Module)")
	}
}

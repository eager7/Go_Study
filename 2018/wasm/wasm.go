package main

import (
	"bytes"
	"io/ioutil"
	"fmt"
	"github.com/go-interpreter/wagon/wasm"
	"github.com/go-interpreter/wagon/exec"
)

func main() {
	testWasm()
}

func testWasm() {
	raw, err := ioutil.ReadFile("basic.wasm")
	if err != nil {
		fmt.Println(err)
		return
	}

	r := bytes.NewReader(raw)
	m, err := wasm.ReadModule(r, nil)
	if err != nil {
		fmt.Printf("error reading module %v\n", err)
		return
	}
	if m == nil {
		fmt.Println("error reading module: (nil *Module)")
		return
	}
	if m.Export == nil {
		panic("module has no export section")
	}

	vm, err := exec.NewVM(m)
	if err != nil {
		fmt.Println("could not create VM:", err)
	}
	for name, e := range m.Export.Entries {
		i := int64(e.Index)
		fidx := m.Function.Types[int(i)]
		ftype := m.Types.Entries[int(fidx)]
		switch len(ftype.ReturnTypes) {
		case 1:
			fmt.Printf("%s() %s => ", name, ftype.ReturnTypes[0])
		case 0:
			fmt.Printf("%s() => ", name)
		default:
			fmt.Println("running exported functions with more than one return value is not supported")
			continue
		}
		if len(ftype.ParamTypes) > 0 {
			fmt.Println("running exported functions with input parameters is not supported")
			continue
		}
		o, err := vm.ExecCode(i)
		if err != nil {
			fmt.Printf("\n")
			fmt.Println("err=%v", err)
		}
		if len(ftype.ReturnTypes) == 0 {
			fmt.Printf("\n")
			continue
		}
		fmt.Printf("%[1]v (%[1]T)\n", o)
	}
}

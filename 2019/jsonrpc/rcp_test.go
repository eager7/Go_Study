package jsonrpc

import (
	"fmt"
	"testing"
)

func TestInitialize(t *testing.T) {
	s := Server{}
	if err := s.Register(new(Int)); err != nil {
		t.Fatal(err)
	}
}

func hello() {
	fmt.Println("hello world!")
}

type Int int
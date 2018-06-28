package unsafe_test

import (
	"testing"
	"github.com/eager7/go_study/2018/unsafe"
)

func TestPointerAndString(t *testing.T) {
	p, err := unsafe.StringToPointer("Test String")
	if err != nil {
		t.Fatal(err)
	}
	str := unsafe.PointerToString(p)
	if str != "Test String" {
		t.Fatal("Error")
	}
}

//go test -v -test.bench=".*"
func BenchmarkPointerAndString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p, err := unsafe.StringToPointer("Test String")
		if err != nil {
			b.Fatal(err)
		}
		str := unsafe.PointerToString(p)
		if str != "Test String" {
			b.Fatal("Error")
		}
	}
}
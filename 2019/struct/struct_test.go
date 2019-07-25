package _struct

import (
	"fmt"
	"testing"
)

func TestAppendSlice(t *testing.T) {
	var list = make([]int, 5)
	AppendSlice(list, 1)
	fmt.Println(list)
}

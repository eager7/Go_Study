package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := []int{1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6,7,8,9,0,1,1,1,1,2}
	fmt.Println(SplitSlice(slice, 1))
	fmt.Println(SplitSlice(slice, 2))
	fmt.Println(SplitSlice(slice, 3))
	fmt.Println(SplitSlice(slice, 4))
	fmt.Println(SplitSlice(slice, 5))
	fmt.Println(SplitSlice(slice, 6))
	fmt.Println(SplitSlice(slice, 7))
	fmt.Println(SplitSlice(slice, 8))

	fmt.Println(SplitSliceLen(slice, 1))
	fmt.Println(SplitSliceLen(slice, 2))
	fmt.Println(SplitSliceLen(slice, 3))
	fmt.Println(SplitSliceLen(slice, 4))
	fmt.Println(SplitSliceLen(slice, 5))
	fmt.Println(SplitSliceLen(slice, 6))
}

func TestCap(t *testing.T) {
	a := make([]string, 0, 0)
	a = append(a, "a", "b", "c")

	b := a
	a = append(a, "d", "e", "f", "g") // 下面代码中 b 能覆盖多少，取决于 a append 了多少
	fmt.Println("cap-a", cap(a), cap(b))
}
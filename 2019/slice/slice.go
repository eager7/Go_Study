package slice

import (
	"fmt"
)

func SplitSlice(slice []int, n int) [][]int {
	var sTemp [][]int
	s := len(slice) / n
	for i := 0; i < s; i++ {
		ll := slice[i*n : (i+1)*n]
		sTemp = append(sTemp, ll)
	}
	sTemp = append(sTemp, slice[s*n:])
	return sTemp
}

func SplitSliceLen(slice []int, length int) [][]int {
	var sTemp [][]int
	n := len(slice) / length
	for i := 0; i < n; i++ {
		ll := slice[i*length : (i+1)*length]
		sTemp = append(sTemp, ll)
	}
	sTemp = append(sTemp, slice[length*n:])
	return sTemp
}


func AppendSlice() {
	type m struct {
		s []interface{}
	}
	a := m{
		s: make([]interface{}, 0, 0),
	}
	a.s = append(a.s, "a")
	a.s = append(a.s, "b")
	a.s = append(a.s, "c")

	b := a
	a.s = append(a.s, "d")
	a.s = append(a.s, "e")

	b.s = append(b.s, "1")
	b.s = append(b.s, "2")
	b.s = append(b.s, "3")

	fmt.Println(1111, a, b)
}


func PSlice() {
	a := make([]string, 0, 8)
	a = append(a, "a", "b", "c")

	b := a
	a = append(a, "d")
	a = append(a, "e")

	fmt.Println(a, b)

	b = append(b, "1")
	b = append(b, "2")
	b = append(b, "3")
	fmt.Println(a, b)
}
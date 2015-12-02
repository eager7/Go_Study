package main

import "fmt"

func main() {
	slice1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slice1)

	slice2 := slice1[:3]
	slice3 := slice1[1:len(slice1)-1]
	slice4 := slice1[1:]
	slice1 = append(slice1[:2],slice1[4:]...)
	fmt.Println(slice1,slice2,slice3,slice4)
}
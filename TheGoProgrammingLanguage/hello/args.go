package main

import (
	"fmt"
	"os"
	"strings"
)

var s, sep string

func main() {
	fmt.Println("args test")
	fmt.Println(len(os.Args), strings.Join(os.Args, "\n"))

	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%s", os.Args[i])
	}
	fmt.Println("")

	for i, v := range os.Args {
		fmt.Println(i, v)
	}

}

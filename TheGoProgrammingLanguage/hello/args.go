package main

import (
	"fmt"
	"os"
)
	
	var s, sep string
func main(){
	fmt.Println("args test")


	for i:=0;i<len(os.Args);i++{
		fmt.Printf("%s",os.Args[i])
	}
	fmt.Println("")
}
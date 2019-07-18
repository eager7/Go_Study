package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main")
	var c chan int
	go func() {
		c <- 1
		fmt.Println("done")
	}()
	time.Sleep(time.Second * 5)
}

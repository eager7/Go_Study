package main

import (
	"fmt"
	_ "github.com/eager7/go_study/2020/jwt"
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

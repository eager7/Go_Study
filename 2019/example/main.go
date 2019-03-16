package main

import (
	"fmt"
	"github.com/eager7/go_study/2019/crontab"
)

func main() {
	fmt.Println("start example...")
	if err := crontab.Initialize(); err != nil {
		panic(err)
	}
	make(chan interface{}) <- struct {}{}
}

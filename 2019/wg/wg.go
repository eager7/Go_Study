package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		fmt.Println("task 1...")
		wg.Done()
	}()
	wg.Wait()

	go func() {
		fmt.Println("task 2...")
		wg.Done()
	}()

	go func() {
		fmt.Println("task 3...")
		wg.Done()
	}()

	fmt.Println("done")
}

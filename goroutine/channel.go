package main
import "fmt"
//import "time"

func main() {
	println("channel test")
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++{
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _,ch := range(chs){
		fmt.Println("Read")
		<-ch
	}
	//<-time.After(time.Millisecond*1000)
}

func Count(ch chan int){
	ch <- 1
	fmt.Println("Counting")
}
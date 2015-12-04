package main
import "fmt"
import "time"

func main() {
	for i:=0;i < 10; i++{
		go Add(i,i)
	}
	time.Sleep(100 * time.Millisecond)
}

func Add(x,y int){
	fmt.Println(x + y)
}
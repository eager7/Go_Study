package main
import "fmt"

func main() {
	chs := make(chan string, 100)
	go func(){
		for ch := range chs{
			fmt.Println("done",ch)
		}
	}()
	
	for ch := range chs{
		fmt.Println("done",ch)
	}

}
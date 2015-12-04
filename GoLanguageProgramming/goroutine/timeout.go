package main
import(
	//"fmt"
	"time"
	)

func main(){
	println("timeout testing")

	timeout := make(chan bool, 1)
	go func(){
		time.Sleep(1e9)
		timeout <- true
	}()

	ch := make(chan int, 1)
	select{
		case <- ch:
			println("Read data form ch")
		case <- timeout:
			println("Timeout")
	}
	close(timeout)
	close(ch)
	_,ok := <-ch
	if ok == false{
		println("ch is close")
	}
}

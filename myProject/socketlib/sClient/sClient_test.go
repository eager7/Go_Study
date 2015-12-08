package sClient
import (
	"fmt"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	fmt.Println("test connect")
	client := NewClient("localhost", 6667)
	client.Init()

	/*client.SocketSendMsg("{\"command\":20,\"sequence\":2}")
	s, err := client.SocketReadMsg()
	if err != nil {
		fmt.Println("read msg err")
		return
	}
	fmt.Println(s)*/


	client.SocketSendMsg("{\"command\":20,\"sequence\":2}")
	s, err := client.SocketReadMsgTime(time.Millisecond*1200)
	if err != nil {
		fmt.Println("read msg err", err)
		return
	}
	fmt.Println(s)

	defer client.Finished()
}


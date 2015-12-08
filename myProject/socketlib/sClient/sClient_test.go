package sClient
import (
	"fmt"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	fmt.Println("test connect")
	client := NewClient()
	client.Init("localhost", 6667)
	time.Sleep(time.Second*1)
	client.Finished()
}


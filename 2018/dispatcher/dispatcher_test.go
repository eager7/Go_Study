package dispatcher_test

import (
	"testing"
	"github.com/eager7/go_study/2018/dispatcher"
	"github.com/eager7/go_study/2018/dispatcher/pb"
	"fmt"
	"sync"
)

type Msg string
func (m *Msg) Type() pb.MsgType {
	return pb.MsgType_APP_MSG_BLK
}
func (m Msg) Data() []byte {
	return []byte(string(m))
}

func TestDispatcher(t *testing.T) {
	dispatcher.InitMsgDispatcher()

	c, err := dispatcher.Subscribe(pb.MsgType_APP_MSG_BLK)
	if err != nil {
		t.Fatal(err)
	}
	m := Msg("test")
	wg := &sync.WaitGroup{}
	wg.Add(2)
	c1, err := dispatcher.Subscribe(pb.MsgType_APP_MSG_BLK)
	dispatcher.Publish(&m)
	go receive(wg, c)
	go receive(wg, c1)
	wg.Wait()
}

func receive(wg *sync.WaitGroup, c <-chan interface{}) {
	defer wg.Done()
	ret := <-c
	fmt.Println(string(ret.(*Msg).Data()))
}
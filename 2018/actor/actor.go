package main

import (
	"reflect"

	"github.com/AsynkronIT/protoactor-go/actor"
	"fmt"
	"time"
)

var server *Server

func main() {
	test()
}

func test() {
	i := 1
	server = &Server{}
	p := &PoolActor{server:server}
	pid, err := NewTxPoolActor(p)
	if err != nil {
		fmt.Println(err)
	}

	timer := time.NewTimer(time.Second * 1)
	for {
		select {
		case <-timer.C:
			pid.Tell(&Test{i})
			i += 1
			timer.Reset(time.Second * 1)
		}
	}
}

type Test struct {
	val int
}

type Server struct {
	value int
}

type Pool struct {
	actor *PoolActor
	server *Server
}

type PoolActor struct {
	server *Server
}

func NewTxPoolActor(p *PoolActor) (*actor.PID, error) {
	props := actor.FromProducer(func() actor.Actor {
		return p
	})
	pid, err := actor.SpawnNamed(props, "TxPoolActor")
	if err != nil {
		return nil, err
	}
	return pid, nil
}


func (l *PoolActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
	case *actor.Restarting:
	case *Test:
		fmt.Println("test value:", msg.val)
		l.server.SetValue(msg)
	default:
		fmt.Println("unknown type message:", msg, "type", reflect.TypeOf(msg))
	}
}

func (s *Server)SetValue(t *Test) {
	s.value = t.val
	fmt.Println("Server:", s.value)
}

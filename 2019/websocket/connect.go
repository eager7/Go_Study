package ws

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"reflect"
	"sync"
)

type Connection struct {
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
	wsConn *websocket.Conn
	send   chan interface{}
	recv   chan interface{}
}

func (c *Connection) Initialize(wsConn *websocket.Conn) *Connection {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), "manager", wg))

	return &Connection{
		wg:     wg,
		ctx:    ctx,
		cancel: cancel,
		wsConn: wsConn,
		send:   make(chan interface{}, 10),
		recv:   make(chan interface{}, 10),
	}
}

func (c *Connection) Start() {
	go c.ReadRoutine(c.ctx, c.wg)
	go c.WriteRoutine(c.ctx, c.wg)
	go c.HandlerRoutine(c.ctx, c.wg)
}

func (c *Connection) Finished() {
	c.cancel()
	c.wg.Wait()
}

func (c *Connection) SendMessage(msg Message) {
	c.send <- msg
}

func (c *Connection) ReadRoutine(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		select {
		default:
			typ, body, err := c.wsConn.ReadMessage()
			if err != nil {
				logger.Println("read message err:", err)
			}
			logger.Println("message:", typ, string(body))
			c.recv <- &BaseMessage{identify: typ, body: body}
		case <-ctx.Done():
			fmt.Println("return read routine...")
			wg.Done()
			return
		}
	}
}

func (c *Connection) WriteRoutine(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		select {
		case in := <-c.send:
			message, ok := in.(Message)
			if ok {
				if err := c.wsConn.WriteMessage(message.Identify(), message.Body()); err != nil {
					log.Println("write message err:", err)
				}
			} else {
				log.Println("can't parse type:", reflect.TypeOf(in))
			}
		case <-ctx.Done():
			fmt.Println("return write routine...")
			wg.Done()
			return
		}
	}
}

func (c *Connection) HandlerRoutine(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		select {
		case in := <-c.recv:
			message, ok := in.(Message)
			if ok {
				log.Println("handle message:", message.Identify(), string(message.Body()))
			} else {
				log.Println("can't parse type:", reflect.TypeOf(in))
			}
		case <-ctx.Done():
			fmt.Println("return handle routine...")
			wg.Done()
			return
		}
	}
}

package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var logger = log.New(os.Stdout, "ws", log.LstdFlags|log.Lshortfile)

type WebSocket struct {
	lock     sync.RWMutex
	errChan  chan *Connection
	upGrader websocket.Upgrader
	listConn map[string]*Connection
}

func Initialize() *WebSocket {
	upGrader := websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
		CheckOrigin: func(r *http.Request) bool { //允许跨域
			return true
		},
		EnableCompression: true, //压缩
	}
	list := make(map[string]*Connection)
	errChan := make(chan *Connection, 10)
	w := &WebSocket{lock: sync.RWMutex{}, errChan: errChan, upGrader: upGrader, listConn: list}
	go w.HandleConnect(context.Background())
	return w
}

func (ws *WebSocket) InitHttpConn(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*Connection, error) {
	wsConn, err := ws.upGrader.Upgrade(w, r, responseHeader)
	if err != nil {
		log.Println("upgrade http connect to ws err:", err)
		return nil, err
	}
	conn := new(Connection).Initialize(wsConn, ws.errChan)
	conn.Start()
	ws.lock.Lock()
	defer ws.lock.Unlock()
	ws.listConn[wsConn.RemoteAddr().String()] = conn
	return conn, err
}

func (ws *WebSocket) HandleConnect(ctx context.Context) {
	for {
		select {
		case conn := <-ws.errChan:
			log.Println("close connect:", conn.wsConn.RemoteAddr())
			conn.Finished()
			ws.lock.Lock()
			delete(ws.listConn, conn.wsConn.RemoteAddr().String())
			ws.lock.Unlock()
		case <-ctx.Done():
			return
		}
	}
}

func (ws *WebSocket) Finished() {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	for _, v := range ws.listConn {
		v.Finished()
	}
}

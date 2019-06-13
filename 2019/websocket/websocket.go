package ws

import (
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
	return &WebSocket{lock: sync.RWMutex{}, upGrader: upGrader, listConn: list}
}

func (ws *WebSocket) InitHttpConn(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*Connection, error) {
	wsConn, err := ws.upGrader.Upgrade(w, r, responseHeader)
	if err != nil {
		log.Println("upgrade http connect to ws err:", err)
		return nil, err
	}
	conn := new(Connection).Initialize(wsConn)
	conn.Start()
	ws.lock.Lock()
	defer ws.lock.Unlock()
	ws.listConn[wsConn.RemoteAddr().String()] = conn
	return conn, err
}

func (ws *WebSocket) Finished() {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	for _, v := range ws.listConn {
		v.Finished()
	}
}

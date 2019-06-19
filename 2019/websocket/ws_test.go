package ws_test

import (
	"fmt"
	"github.com/eager7/go_study/2019/websocket"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestWs(t *testing.T) {
	web := ws.Initialize()

	serverWs := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("receive connect")
		_, err := web.InitHttpConn(w, r, nil)
		if err != nil {
			t.Fatal(err)
		}
	}
	go client(t)
	http.HandleFunc("/ws", serverWs)
	log.Fatal(http.ListenAndServe("127.0.0.1:13142", nil))
}

func client(t *testing.T) {
	time.Sleep(time.Second * 1)
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:13142", Path: "/ws"}
	fmt.Println("connecting to ", u.String())
	cli, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("connect to server success:", cli.RemoteAddr().String(), resp.StatusCode)
	defer func() {
		fmt.Println("close connect:", cli.Close())
	}()
	if resp == nil || resp.StatusCode != http.StatusSwitchingProtocols {
		fmt.Println("status err:", resp)
		t.Fatal(resp)
	}
	fmt.Println("send msg 1")
	if err := cli.WriteMessage(1, []byte("test message 1")); err != nil {
		t.Fatal(err)
	}
	fmt.Println("send msg 2")
	if err := cli.WriteMessage(1, []byte("test message 2")); err != nil {
		t.Fatal(err)
	}
}

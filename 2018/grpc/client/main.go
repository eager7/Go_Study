package main

import (
	"sync"
	"google.golang.org/grpc"
	"github.com/eager7/go_study/2018/grpc/proto"
	"context"
	"github.com/eager7/go/mlog"
)

var (
	wg sync.WaitGroup
)
var log = mlog.NewLogger("client", mlog.NoticeLog)
const (
	networkType = "tcp"
	server      = "127.0.0.1"
	port        = "50001"
	parallel    = 50        //连接并行度
	times       = 100000    //每连接请求次数
)

func main() {
	log.Info("grpc client process")
	call()
}

func call(){
	conn, err := grpc.Dial(server + ":" + port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Debug("connect success")
	client := inf.NewDataClient(conn)

	resp, err := client.GetUser(context.Background(), &inf.UserReq{Id:1})
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(resp.Name)
}
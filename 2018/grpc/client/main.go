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
var log,_ = mlog.NewLogger("client", mlog.NoticeLog)
const (
	networkType = "tcp"
	server      = "127.0.0.1"
	port        = "50001"
	parallel    = 50        //连接并行度
	times       = 100000    //每连接请求次数
)

func main() {
	call()
}

func call(){
	conn, err := grpc.Dial(server + ":" + port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := inf.NewDataClient(conn)

	var req inf.UserReq
	req.Id = 1

	resp, _ := client.GetUser(context.Background(), &req)
	log.Debug(resp.Name)
}
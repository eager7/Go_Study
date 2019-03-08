package main

import (
	"github.com/eager7/go/mlog"
	"github.com/eager7/go_study/2018/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

const port = 50001

type Data struct {
}

var log = mlog.NewLogger("grpc", mlog.NoticeLog)

func main() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	inf.RegisterDataServer(s, &Data{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func (d *Data) GetUser(ctx context.Context, req *inf.UserReq) (*inf.UserResp, error) {
	log.Debug("receive client request:", req)
	resp := &inf.UserResp{Name: strconv.Itoa(int(req.Id)) + ":test"}

	return resp, nil
}

package main

import (
	"github.com/eager7/go/mlog"
	"net"
	"google.golang.org/grpc"
	"github.com/eager7/go_study/2018/tx/proto"
	"context"
	"github.com/eager7/go_study/2018/grpc/proto"
)

const port = "50001"
const server = "127.0.0.1"

type RpcMsg struct {}
var log, _ = mlog.NewLogger("server", mlog.NoticeLog)

func main() {
	lis, err := net.Listen("tcp", ":"+ port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	txmsg.RegisterGetVersionServer(s, &RpcMsg{})
	txmsg.RegisterTransferTxServer(s, &RpcMsg{})
	s.Serve(lis)
}

func (msg *RpcMsg)GetVersion(ctx context.Context, req *txmsg.VersionReq) (*txmsg.VersionResp, error){
	return &txmsg.VersionResp{Version:"v1.0"}, nil
}
func (msg *RpcMsg)TransferTx(ctx context.Context, req *txmsg.TxTransfer) (*txmsg.TxTransferResp, error){
	//TODO: handle tx
	resp := &txmsg.TxTransferResp{Ret:true}
	return resp, nil
}

func Call(req interface{})(resp interface{}){
	conn, err := grpc.Dial(server + ":" + port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Debug("connect success")

	switch t := req.(type) {
	case *txmsg.VersionReq:
		client := txmsg.NewGetVersionClient(conn)
		r, err := client.GetVersion(context.Background(), t)
		if err != nil {
			log.Fatal(err)
		}
		return r
	case *txmsg.TxTransfer:
		client := txmsg.NewTransferTxClient(conn)
		r, err := client.TransferTx(context.Background(), t)
		if err != nil {
			log.Fatal(err)
		}
		return r
	}

	return nil
}

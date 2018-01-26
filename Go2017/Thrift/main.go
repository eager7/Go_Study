package main

import (
	"context"
	"fmt"
	"time"
	"./thrift/gen-go/batu/demo"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type batuThrift struct {
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket("127.0.0.1:9090")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	handler := &batuThrift{}
	processor := demo.NewBatuThriftProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Server Start:127.0.0.1:9090")
	server.Serve()
}

func (b *batuThrift) CallBack(ctx context.Context, callTime int64, name string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("From Client Call:", time.Unix(callTime, 0).Format("1900-01-01 15:00:00"), name, paramMap)
	r = append(r, "key:"+paramMap["a"]+"value:"+paramMap["b"])
	return
}

func (b *batuThrift) Put(ctx context.Context, s *demo.Article) (err error) {
	fmt.Printf("Article--id:%d Title:%s Context:%t Author:%d\n", s.ID, s.Title, s.Content, s.Author)
	return
}

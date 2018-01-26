package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"../thrift/gen-go/batu/demo"
)

func main() {
	startTime := time.Now()

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9090"))
	if err != nil {
		fmt.Println(err)
		return
	}

	useTransport, _ := transportFactory.GetTransport(transport)
	client := demo.NewBatuThriftClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Println(err)
		return
	}
	defer transport.Close()

	for i := 0; i < 10; i++{
		paramMap := make(map[string]string)
		paramMap["a"] = "batu.demo"
		paramMap["b"] = "test" + strconv.Itoa(i)
		r, _ := client.CallBack(nil, time.Now().Unix(), "go client", paramMap)
		 fmt.Println("go client call :", r)
	}

	model := demo.Article{1, "first", "i am here", "pct"}
	client.Put(nil, &model)
	fmt.Println("time:", time.Now().Sub(startTime))
}

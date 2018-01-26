package remote

import (
	"../thrift/gen-go/method"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
)

func ServerThrift(){
	tFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	pFactory := thrift.NewTBinaryProtocolFactoryDefault()

	sTransport, err := thrift.NewTServerSocket("127.0.0.1:9090")
	if err != nil {
		fmt.Println(err)
		return
	}
	handler := NewMethod()
	processor := method.NewMethodProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, sTransport, tFactory, pFactory)
	server.Serve()
}

func ClientThrift(){
	tFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	pFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket("127.0.0.1:9090")
	if err != nil {
		fmt.Println(err)
		return
	}
	utransport, _ := tFactory.GetTransport(transport)
	client := method.NewMethodClientFactory(utransport, pFactory)
	if err := transport.Open(); err != nil {
		fmt.Println(err)
		return
	}
	defer transport.Close()

	if _, err := client.Add(nil, "pct", "panchangtao"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := client.Add(nil, "lyz", "laiyongzheng"); err != nil {
		fmt.Println(err)
		return
	}
	r, err := client.Get(nil, "pct")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
package remote

import (
	"testing"
	"../thrift/gen-go/method"
	"time"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
)

func TestStart(t *testing.T) {

	go ServerThrift()
	time.Sleep(time.Second*1)

	tFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	pFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket("127.0.0.1:9090")
	if err != nil {
		t.Fatal(err)
	}
	utransport, _ := tFactory.GetTransport(transport)
	client := method.NewMethodClientFactory(utransport, pFactory)
	if err := transport.Open(); err != nil {
		t.Fatal(err)
	}
	defer transport.Close()

	if _, err := client.Add(nil, "pct", "panchangtao"); err != nil {
		t.Fatal("Add err:", err.Error())
	}
	if _, err := client.Add(nil, "lyz", "laiyongzheng"); err != nil {
		t.Fatal(err)
	}
	r, err := client.Get(nil, "pct")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}

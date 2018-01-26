package main

import (
	"fmt"
	"net"
	"bytes"
	"time"
)

func initServer(hostAndport string)*net.TCPListener{
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndport)
	if err != nil {
		fmt.Println("Resolve Addr Failed:", err.Error())
		return nil
	}
	listener, err := net.ListenTCP("tcp", serverAddr)
	if err != nil{
		fmt.Println("Listen TCP failed:", err.Error())
		return nil
	}
	fmt.Println("Listen Address:", listener.Addr().String())
	return listener
}

func main(){
	fmt.Println("tcp server start...")
	listener, err := net.Listen("tcp", "10.128.0.101:8080")
	if err != nil{
		fmt.Println("Listern Falied:", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("Accept Falied:", err.Error())
			return
		}
		go doServer(conn)
	}
}

func doServer(conn net.Conn){
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(5*time.Minute))
	for{
		buf := make([]byte, 512)
		l, err := conn.Read(buf)
		if err != nil{
			fmt.Println("Read Len:", l, "Failed:", err.Error())
			return
		} else{
			fmt.Println("Read Len: ", l, "Data:", string(bytes.TrimRight(buf, "\x00")))//去除多余的0
		}
		if _, err := conn.Write([]byte("I am Server"));err != nil{
			fmt.Println("Write Failed:", err.Error())
			return
		}
	}
}


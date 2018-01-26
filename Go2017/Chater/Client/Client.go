package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main(){
	fmt.Println("Client Start...")
	conn, err := net.Dial("tcp", "10.128.0.101:7878")
	if err != nil {
		fmt.Println("Dial Failed:", err.Error())
		return
	}
	defer conn.Close()
	bufferReader := bufio.NewReader(os.Stdin)
	fmt.Println("input your name")
	clientName, _ := bufferReader.ReadString('\n')
	//clientName = strings.Trim(clientName, "\n")
	buffer := make([]byte, 128)
	for{
		fmt.Println("Input Data to Server, Type Q to Quit")
		input, _ := bufferReader.ReadString('\n')
		input = strings.Trim(input, "\n")
		if input == "Q"{
			return
		}
		conn.Write([]byte(clientName+":"+input))
		conn.Read(buffer)
		fmt.Println(strings.TrimRight(string(buffer), "\x00"))
	}
}

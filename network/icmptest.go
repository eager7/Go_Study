package main
import (
	"net"
	"os"
	//"bytes"
	"fmt"
)

func main(){
	println("icmptest program...")

	if len(os.Args) != 2{
		fmt.Println("usage:", os.Args[0], "host")
		os.Exit(1)
	}
	server := os.Args[1]

	conn, err := net.Dial("ip4:icmp", server)
	checkError(err)

	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37

	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	fmt.Println("write data")
	_, err = conn.Write(msg[0:len])
	checkError(err)

	fmt.Println("read data")
	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("Got Response")
	if msg[5] == 13{
		fmt.Println("Identifier matchs")
	}
	if msg[7] == 37{
		fmt.Println("Sequene matchs")
	}
	os.Exit(0)
}

func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0

	for n := 1; n <len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}
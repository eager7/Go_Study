package socketClient
import (
	"fmt"
	"net"
	"os"
	"strconv"
	"io"
	"bytes"
	"time"
	//"io/ioutil"
	)

type SockClient struct {
	conn net.Conn
}

func (s *SockClient) Init(addr string, port int) error {
	fmt.Println("Socket Client Init With", addr, port)
	dial := addr + ":" + strconv.Itoa(port)
	fmt.Println(dial)

	var err error
	s.conn, err = net.Dial("tcp", dial)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s\n", err.Error())
		return err
	}

	return nil
}

func (s *SockClient) Finished() error {
	fmt.Println("Finished Socket Client")
	s.conn.Close()

	return nil
}

func (s *SockClient) SendMsg(msg string) error {
	fmt.Println("Socket Send Msg")
	_, err := s.conn.Write([]byte(msg))
	if err != nil {
		fmt.Fprintf(os.Stderr, "SendMsg Fatal error:%s\n", err.Error())
		return err		
	}
	return nil
}

func (s *SockClient) ReadMsg() (string, error) {
	fmt.Println("Socket Read Msg")
	bufW := bytes.NewBuffer(nil)
	var bufTemp [256]byte
	for {
		n, err := s.conn.Read(bufTemp[0:])
		bufW.Write(bufTemp[:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "SendMsg Fatal error:%s\n", err.Error())
			return "", err
		}
		if n < 256 {
			break
		}
	}

	return string(bufW.Bytes()), nil
}

func (s *SockClient) writeAndReadMsg(ch chan string, msg string){
	fmt.Println("writeAndReadMsg")
	if err := s.SendMsg(msg);err != nil {
		fmt.Fprintf(os.Stderr, "s.SendMsg Fatal error:%s\n", err.Error())
		ch <- ""
	}
	result, err := s.ReadMsg()
	if err != nil {
		fmt.Fprintf(os.Stderr, "s.ReadMsg Fatal error:%s\n", err.Error())
		ch <- ""
	}
	ch <- result
}

func (s *SockClient) SendMsgWithResp(msg string) (string, int) {
	fmt.Println("SendMsgWithResp", msg)
	ch := make(chan string, 1)
	go s.writeAndReadMsg(ch, msg)

	select {
		case x := <- ch:{
			fmt.Println(x)
			return x, 0
		}
		case <- time.After(time.Millisecond*200):{
			fmt.Println("time out")
			return "", -1
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
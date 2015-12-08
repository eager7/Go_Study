package sClient
import (
	"fmt"
	"os"
	"net"
	"strconv"
	"bytes"
	"io"
	"errors"
	"time"
)

type SockClient struct {
	conn net.Conn
	addr string
	port int
	chjob chan string
}

func NewClient(addr string, port int) *SockClient {
	chjob := make(chan string, 1)
	return &SockClient{conn:nil, addr:addr, port:port, chjob:chjob}
}

func (s *SockClient) Init() (err error) {
	fmt.Println("init the socket client with", s.addr, s.port)

	dial := s.addr + ":" + strconv.Itoa(s.port)
	s.conn, err = net.Dial("tcp", dial)
	if err != nil {
		fmt.Fprintf(os.Stderr, "socket client init err:%s\n", err.Error())
		return err
	}
	go s.service(s.chjob) //start recv msg
	return nil
}

func (s *SockClient) Finished() {
	//close(s.chjob)
	s.conn.Close()
}

func (s *SockClient) SocketSendMsg(msg string)error{
	if _, err := s.conn.Write([]byte(msg)); err != nil {
		fmt.Fprintf(os.Stderr, "socket send msg err:%s\n", err.Error())
		return err
	}
	return nil
}

func (s *SockClient) SocketReadMsg() (string, error) {
	msg := <- s.chjob
	if msg == "" {
		return "", errors.New("can't read msg from server")
	}
	return msg, nil			
}

func (s *SockClient) SocketReadMsgTime(timeout time.Duration) (string, error) {
	select {
		case msg := <- s.chjob: {
			if msg == "" {
				return "", errors.New("can't read msg from server")
			}
			return msg, nil			
		}
		case <- time.After(timeout): {
			fmt.Println("time out")
			return "", errors.New("time out")
		}
	}
}

func (s *SockClient) readMsg()(string, error) {
	bufw := bytes.NewBuffer(nil)
	var buftemp [2048]byte
	for {
		n, err := s.conn.Read(buftemp[0:])
		bufw.Write(buftemp[0:n])
		if err != nil {
			if err == io.EOF {
				return string(bufw.Bytes()), err
			}
			fmt.Fprintf(os.Stderr, "read msg from socket err:%s\n", err.Error())
			return "", err
		}
		if n < 2048 {
			return string(bufw.Bytes()), nil
		}
	}
}

func (s *SockClient)service(chjob chan string) {
	for {
		fmt.Println("read msg...")
		result, err := s.readMsg()
		if err != nil {
			fmt.Fprintf(os.Stderr, "socket service err:%s\n", err.Error())
			chjob <- ""
			fmt.Println("socket read service exit..")
			return
		}
		fmt.Println("read msg...:", result)
		chjob <- result
	}
}

func (s *SockClient)reconnect() {
	s.Finished()
	for {
		if err := s.Init(); err == nil {
			return
		}
	}
}
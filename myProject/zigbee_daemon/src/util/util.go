package util

type SockStatus int 
const (
	SOCK_OK 			SockStatus = iota
	SOCK_ERROR
	SOCK_CONNECT_FAIL
	SOCK_DISCONNECT
)



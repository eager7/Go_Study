package main
//import "fmt"

type enum int
const (
	OK enum = iota
	FAIL
	UNKNOW
)

func (e enum)String() string {
	switch e{
		case OK: {
			return "ok"
		}
		case FAIL:{
			return "Failed"
		}
		case UNKNOW:{
			return "unkonw"
		}
	}
	return "nil"
}

func main(){
	
}
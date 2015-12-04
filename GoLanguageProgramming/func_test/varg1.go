package main
import "fmt"

func MyPrintf(args ...interface{}){
    for _,arg := range args{
        switch arg.(type){
            case int:
                fmt.Println(arg, " is a int value.")
            case string:
                fmt.Println(arg, " is a string value.")
            case int64:
                fmt.Println(arg, " is a int64 value.")
            default:
                fmt.Println(arg, " is a unknow type.")
        }
    }
}

func main(){
    v1 := 1
    var v2 int64 = 234
    v3 := "hello"
    v4 := 1.234

    MyPrintf(v1, v2, v3, v4)
}

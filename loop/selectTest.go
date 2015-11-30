package main

import "fmt"

func main(){
    fmt.Println("Hello\n")

    var c1 chan int
    var i1 int
    select {
        case i1 = <- c1:
            println("recvived ", i1, "from c1")
        default:
            println("no communication\n")
    }
}

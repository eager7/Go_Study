package main
import "fmt"

func main(){
    println("this is a closure func test")
    var j int = 5

    a := func()(func()){
        var i int = 10
        return func(){
            fmt.Printf("i is %d, j is %d\n",i,j)
        }
    }()
    a()
    j = j * 2
    a()
}

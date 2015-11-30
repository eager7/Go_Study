package main

import "fmt"

func main(){
    fmt.Println("Hello\n")
    
    var a,b int = 1,15
    
    numbers := [6]int{1,2,3,4}

    for a:= 0; a < 10; a ++{
        fmt.Println("a is ", a)
    }

    for a < b {
        a++
        println("a = ", a)
    }

    for i,x := range numbers{
        fmt.Printf("numbers[%d] = %d\n", i, x)
    }
}

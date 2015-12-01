package main

import "fmt"

func main(){
    fmt.Println("Hello\n")

    var flag_pri int = 0

    for i:= 2; i < 100; i++{
        flag_pri = 0
        for j:=2; j < i; j++{
            if(i%j == 0){
                //println(i, " is not a primer")
                flag_pri = 1
                break;
            }
        }
        if flag_pri == 1{
            //println(i," is not a primer")
        }else{
            println(i, " is a primer")
        }
    }
}

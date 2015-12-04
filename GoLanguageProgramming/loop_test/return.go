package main

func main(){
    println("this is a return test")
    re(2)
}

func re(x int) int{
    if x == 0{
        println("return 0")
        return 1
    }else{
        println("return ",x)
        return x
    }
    //return 0
}

package main

//from fmt import Println
import (
	"fmt"
	"io"
	"log"
	"runtime"
)

const LIM = 41
var fibs [LIM] uint64
var fib uint64
var where = func(){
	_,file,line,_ := runtime.Caller(1)
	log.Printf("%s,%d", file, line)
}
func main(){
	fmt.Println("Function")
	where()
	fmt.Println(Min(2,4,5,7,8,86,6554,3,))
	arr:=[]int{7,9,3,5}
	fmt.Println(Min(arr...))
	log1("Go")
	pri := func (s string){fmt.Println(s)}
	pri("Hello World!")
	fmt.Println(Add2()(3),Adder(2)(3), Adder2()(3),Adder2()(3))
	f:=Adder2()
	fmt.Println(f(3),f(3))
	f2:=Adder2()
	fmt.Println(f2(3),f2(3))
	where()
	for i:=0;i<LIM;i++{
		fmt.Println(i,"-",fibonacci(i))
	}
}

func Min(a ...int)int{
	defer un(trace("Min"))
	if len(a) == 0{
		return 0
	}
	min := a[0]
	for _,v := range a{
		if v<min{
			min = v
		}
	}
	return min
}

func trace(s string)string{
	fmt.Println("Enter:", s)
	return s
}
func un(s string){
	fmt.Println("Leave:", s)
}
func log1(s string)(n int, err error){
	defer func(){
		log.Println("log1:",s,"=",n,",",err)
	}()
	return 7, io.EOF
}

func Add2() (func(b int) int){
	return func(b int) int {
		return b +2
	}
}
func Adder(a int) (func(b int) int){
	return func (b int) int{
		return a+b
	}
}
func Adder2()(func(b int)int){
	var x int
	return func (b int)int{
		x= x+b
		return x
	}
}
func fibonacci(n int)(res uint64){
	if fibs[n] != 0{
		res = fibs[n]
		return res
	}
	if n<=1 {
		res = 1
	}else{
		res = fibonacci(n-1)+fibonacci(n-2)
	}
	fibs[n] = res
	return res
}
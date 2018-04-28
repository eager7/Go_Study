package main

import (
"fmt"
	"runtime"
	"sync"
	"time"
	"github.com/garyburd/redigo/redis"
	"github.com/eager7/go/mlog"
)

func main(){
mlog.Debug.Println("test")
}

func testSql(){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil{
		fmt.Println("Connect to redis error:", err.Error())
		return
	}
	defer c.Close()
	if _, err = c.Do("SET", "mykey", "panchangtao"); err != nil{
		fmt.Println("set failed:", err.Error())
		return
	}
	username,err := redis.String(c.Do("GET", "mykey"))
	if err != nil{
		fmt.Println("get failed:", err.Error())
		return
	}
	fmt.Println("get key :", username)

	time.Sleep(1)
}

func say(s string){
	for i := 0; i < 5; i++{
		runtime.Gosched()
		fmt.Println(s)
	}
}

func produce(p chan <- int){
	for i:=0; i<10;i++{
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer(c <-chan int){
	for i:=0;i<10;i++{
		v:= <-c
		fmt.Println("receive:", v)
	}
}


func chanTest(){
	ch := make(chan int)
	go func() {
		<- ch
		fmt.Println("1")
	}()
	ch <- 1
	fmt.Println(2)
}

func defer_caller(){
	defer func(){fmt.Println("1")}()
	defer func(){fmt.Println("2")}()
	defer func(){fmt.Println("3")}()

	panic("defer caller")
	defer func(){fmt.Println("4")}()
}

type student struct{
	name string
	age int
}

func pase_students(){
	m := make(map[string] *student)
	stus := []student{
		{"zhou", 24},
		{"li", 23},
		{"pan", 29},
	}
	for _, stu := range stus{
		m[stu.name] = &stu //stu是临时变量，指向最后一个地址
	}

	for _, stu := range stus {
		func(x student) {
			m[x.name] = &x
		}(stu)
	}

	fmt.Println(m["li"]) // &{pan 29}
}

func test(){
	runtime.GOMAXPROCS(1)//锁定为单核执行
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {//因为是单核执行，因此go的协程需要等到主线程执行完成后才有机会运行，此时i已经累积到了10，所以打印全部是10
		go func(){
			fmt.Println("i:", i)
			wg.Done()
		}()
	}
	for j := 0; j < 10; j++ {//func中的i是它内部的变量，和外部i无关，因此可以换成j，k都OK，关键是函数创建时内部的参数值已确定
		go func(i int){
			fmt.Println("j:", i)
			wg.Done()
		}(j)
	}
	wg.Wait()
}

type People struct{}
func (p *People)ShowA(){
	fmt.Println("showA")
	p.ShowB()
}
func (p *People)ShowB(){
	fmt.Println("showB")
}
type Teacher struct{
	People
}
func (t *Teacher)ShowB(){
	fmt.Println("teacher showB")
}

func chan_test(){
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select{
	case value:=<-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

func calc(index string, a, b int)int{
	ret := a+b
	fmt.Println(index, a, b, ret)
	return ret
}

func calc_test(){
	a,b := 1,2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

type Man interface {
	Speaker(string)string
}
type Stu struct{}
func (s *Stu)Speaker(think string)string{
	var talk string
	if think == "bitch"{
		talk = "you are a good man"
	}else {
		talk = "hi"
	}
	return talk
}

func testMan(){
	var m Man = &Stu{}
	think := "bitch"
	fmt.Println(m.Speaker(think))
}

func live()Man {
	var s *Stu
	return s
}
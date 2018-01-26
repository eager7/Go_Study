package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("Struct Method")
	test1()

	pers1 := Person{"Chris", "Woodward"}
	fmt.Println(pers1)
	pers1.upPerson()
	fmt.Println(pers1)

	li := &employee{"liming", 100}
	li.giveRaise(53)
	fmt.Println(li)
}

func test1(){
	type st1 struct{
		i1 int
		f1 float32
		str string
	}
	ms := new(st1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "abc"
	fmt.Println(ms)
}
type Person struct{
	firstName string
	lastName string
}
func (p *Person)upPerson(){
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

type employee struct{
	name string
	salary int
}

func (e *employee)giveRaise(n int){
	e.salary += e.salary * n / 100
}
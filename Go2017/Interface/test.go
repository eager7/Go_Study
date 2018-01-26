package main

import "fmt"

func main(){
	fmt.Println("test")

	var b Int = 5
	fmt.Println(b.get())

	c := Circle{1,2 ,4}

	c.set(5)
	fmt.Println(c)
	c.set2(5)
	fmt.Println(c)
}

type Int int
func (a *Int)get()int{
	return int(*a)
}

type Circle struct {
	x,y int
	radius int
}

func (c Circle)set(r int){
	c.radius = r
}

func (c *Circle)set2(r int){
	c.radius = r
}
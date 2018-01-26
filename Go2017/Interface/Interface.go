package main

import "fmt"

func main(){
	fmt.Println("Interface")

	sq := &Square{5}
	var shaper Shaper = sq
	fmt.Println(shaper.Area())
	cy := &Cycle{5}
	shaper = cy
	fmt.Println(shaper.Area())
	if t,ok := shaper.(*Cycle); ok{
		fmt.Printf("type:%T\n", t)
	}else {
		fmt.Println("error")
	}

	simple:=&Simple{4,5}
	var simpler Simpler = simple
	func1(simpler)

	classifier(true, 8, 7.554, "abcedft", nil, shaper)

	l := &List{1,2,3,4,5}
	fmt.Println(l.Len())
	fmt.Println(l.Len2())

	inter(true)

	fmt.Println("over")
}

type Shaper interface{
	Area() float32
	//Length() float32
}

type Square struct{
	side float32
}
func (s *Square) Area()float32{
	return  s.side * s.side
}

type Cycle struct{
	length float32
}
func (c *Cycle)Area()float32{
	return 3.14159*c.length*c.length
}

type Simpler interface{
	Get() int
	Set() int
}
type Simple struct{
	m int
	n int
}
func (s *Simple)Get()int{
	fmt.Println("Get")
	return s.m
}
func (s *Simple)Set()int{
	fmt.Println("Set")
	return s.n
}
func func1(s Simpler){
	s.Get()
	s.Set()
}

func classifier(items ...interface{}){
	for _, x := range items{
		switch i := x.(type) {
		case bool:
			fmt.Println("bool:", i)
		case float32, float64:
			fmt.Println("float:", i)
		case int, int64:
			fmt.Println("int:", i)
		case nil :
			fmt.Println("NULL:", i)
		case string:
			fmt.Println("string:", i)
		default:
			fmt.Println("unkonwn")
		}
	}
}

type List []int
func (l List)Len()int{
	return len(l)
}
func (l *List)Len2()int{
	return len(*l)
}

func inter(b interface{}){
	if _,ok := b.(bool); ok{
		fmt.Println(b.(bool))
	}else {
		fmt.Println("format error")
	}
	fmt.Println(b.(bool))
}
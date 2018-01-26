package main

import (
	"fmt"
	"runtime"
	"errors"
	"strings"
	"bytes"
	"crypto/rand"
)

var prompt = "Enter a digit, e.g. 3" + " or %s to quit"
func init(){
	fmt.Println("Init Function")
	if runtime.GOOS == "windows"{
		prompt = fmt .Sprintf(prompt, "Ctrl+Z, Enter")
	}else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Println(prompt)
}

func tswitch(num int){
	switch num {
	case 98,99:
		fmt.Printf("num:%d\n", num)
	case 100:
		fmt.Println("num:100")
	default:
		fmt.Println("It is not equal to 98 or 100")
	}
}

func tswitch2(){
	k := 6
	switch k {
	case 4:fmt.Println("was <= 4");fallthrough
	case 5:fmt.Println("was <= 5");fallthrough
	case 6:fmt.Println("was <= 6");fallthrough
	case 7:fmt.Println("was <= 7");fallthrough
	case 8:fmt.Println("was <= 8");fallthrough
	default:
		fmt.Println("defaut case")
	}
}
func season(month int)(season string, err error){
	switch month {
	case 3,4,5:
		return "spring", nil
	case 6,7,8:
		return "summer",nil
	case 9,10,11:
		return "autumn", nil
	case 12,1,2:
		return "winter", nil
	default:
		return "None", errors.New("wrong number")
	}
}
func fizzbuzz(){
	for i:=0;i<=100;i++{
		remainder1:=i%3
		remainder2:=i%5
		switch  {
		case remainder1==0 && remainder2==0:
			fmt.Println("FizzBuzz")
		case remainder1==0:
			fmt.Println("Fizz")
		case remainder2==0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
func rettangle_star(){
	for i:=0;i<5;i++{
		fmt.Print("*")
		for j:=0;j<18;j++{
			if i == 0 || i == 4{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println("*")
	}
}
func mult_returnval(x int, y int)(sum int, product int, difference int){
	sum = x + y
	product = x*y
	difference = x-y
	return sum,product,difference
}

func basename(s string) string{
	slash := strings.LastIndex(s, "/")
	fmt.Println("slash:", slash)
	s = s[slash+1:]
	var dot int
	if dot = strings.LastIndex(s, ".");dot > 0{
		s = s[:dot]
	}
	fmt.Println("dot:", dot)

	return s
}
func intToString(values []int)string{
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _,i := range values{
		fmt.Fprintf(&buf, "%d,", i)
	}
	buf.WriteByte(']')
	return buf.String()
}
func randString()string{
	var r [16]byte
	if _,err := rand.Read(r[:]); err==nil{
		fmt.Println(r)
		return string(r[:])
	}
	return ""
}
func main(){
	fmt.Printf("Hello World! RunTime GOOS:%s\n", runtime.GOOS)
	t := true
	if t{
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	tswitch(98)
	tswitch(99)
	tswitch(100)
	tswitch(11)
	tswitch2()
	if month,err := season(2);err==nil{
		fmt.Println(month)
	}
	if month,err := season(18);err==nil{
		fmt.Println(month)
	}else{
		fmt.Println(err.Error())
	}
	fizzbuzz()
	rettangle_star()
	fmt.Println(mult_returnval(5,3))

	fmt.Println(basename("abc.k.l.j"))
	fmt.Println(intToString([]int{2,4,6,3,4,7}))
	fmt.Println(randString())
	fmt.Println(randString())
}

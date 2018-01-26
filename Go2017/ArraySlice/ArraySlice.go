package main

import (
	"fmt"
)

func main(){
	fmt.Println("Array and Slice Test")
	a:=[10]int{}//array
	fibonacci(&a)
	fmt.Println(a)
	b:=a[:]//slice
	f := fibonacci2()
	f(b)
	fmt.Println(b)
	c:=make([]int,11)
	fibonacci2()(c)
	fmt.Println(c)

	var p = new([]int)
	fmt.Println(p)

	s:=[]byte{1,2,3,4,5}
	fmt.Println(len(s), cap(s))
	s = s[2:4]
	fmt.Println(len(s), cap(s))
	for _,l := range s{
		fmt.Print(l,",")
	}
	fmt.Println("")

	s1:=[]byte{1,2,3,4,7,8}
	s2:=[]byte{4,5,6}
	fmt.Println(Append(s1,s2))
	foreach()
	numbers := []float32{1.22,2.22,3.22,4.22,5,22}
	fmt.Println(Sum(numbers))

	s3:=[]int{1,2,3,4,5,6,7,8,9}
	check:=func(n int)bool{
		if n % 2 == 0{
		return true
	}
		return false
	}
	fmt.Println(Filter(s3, check))
	s4:=[]int{10,11,12}
	fmt.Println(InsertStringSlice(s3, s4, 5))

	s3=append(s3[:2], s3[7:]...)
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))

	str := "abcdefghijklmno"
	fmt.Println(strstr(str, 5))
	fmt.Println(str)
	fmt.Println(str[len(str)/2:] + str[:len(str)/2])

	fmt.Println(revert(str))
}

func fibonacci(a*[10]int){//array
	(*a)[0],(*a)[1] = 1,1
	for i:=2;i<len(a);i++{
		(*a)[i] = (*a)[i-1]+(*a)[i-2]
	}
}

func fibonacci2()(func (l []int)){//slice
	return func(l []int){
		l[0],l[1] = 1,1
		for i := 2; i<len(l);i++  {
			l[i] = l[i-1]+l[i-2]
		}
	}
}

func Append(sl, data[]byte)([]byte){
	d:=make([]byte,len(sl)+len(data))
	for i:=0;i<len(sl);i++{
		d[i]=sl[i]
	}
	for i := len(sl); i < cap(d); i++ {
		d[i]=data[i-len(sl)]
	}
	return d
}
func foreach(){
	items := []int{10,20,30,40,50}
	for _,item := range items{
		item *= 2
	}
	fmt.Println(items)
}
func Sum(numbers []float32)(ret float32){
	for _,n := range numbers{
		ret += n
	}
	return ret
}
func Filter(ss[]int, f func(n int)bool)(slice []int){
	for _,s := range ss{
		if f(s) {
			slice = append(slice, s)
		}
	}
	return slice
}

func InsertStringSlice(s1[]int, s2[]int, index int)([]int){
	ss1:=make([]int, 1)
	ss1 = append(ss1, s1[:index]...)
	ss1 = append(ss1, s2...)
	ss1 = append(ss1, s1[index:len(s1)-1]...)
	return ss1
}

func strstr(str string, index int)(str1 string, str2 string){
	bytes:=[]byte(str)
	bytes1 := bytes[:index]
	bytes2 := bytes[index:]
	return string(bytes1), string(bytes2)
}

func revert(str string)string{
	bytes := []byte(str)
	for index := 0; index < len(bytes)/2; index++{
		bytes[index], bytes[len(bytes) - index - 1] = bytes[len(bytes) - index - 1], bytes[index]
	}
	return string(bytes)
}

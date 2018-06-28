package unsafe

import (
	"unsafe"
	"fmt"
	"reflect"
	"strconv"
)

func test1() {
	str := "hello world!\n"
	fmt.Println("string:", str)
	fmt.Println("---------string to pointer---------")
	fmt.Println("str len:", len(str))
	p := unsafe.Pointer(&str)
	fmt.Println("str addr:", p, "type:", reflect.TypeOf(p))

	fmt.Println("---------pointer to string---------")
	str2 := (*string)(unsafe.Pointer(p))
	fmt.Println(*str2)

	fmt.Println("---------pointer to int32---------")
	strPointerInt := fmt.Sprintf("%d", p)
	i, _ := strconv.ParseInt(strPointerInt, 10, 0)
	fmt.Println("pointer to int32:", i)

	fmt.Println("---------int32 to pointer---------")
	var s *string
	s = *(**string)(unsafe.Pointer(&i))
	fmt.Println("s:", s)

	fmt.Println("---------pointer to string---------")
	str3 := (*string)(unsafe.Pointer(s))
	fmt.Println(*str3)
}



func StringToPointer(str string) (uint64, error) {
	strPointerInt := fmt.Sprintf("%d", unsafe.Pointer(&str))
	return strconv.ParseUint(strPointerInt, 10, 0)
}

func PointerToString(pointer uint64) string {
	var s *string
	s = *(**string)(unsafe.Pointer(&pointer))
	str := *(*string)(unsafe.Pointer(s))
	return str
}
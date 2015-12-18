package main

import "fmt"
import "os"
import "bufio"


func checkError(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
func revert_string(str string) string {
	b := []byte(str)
	n := len(b)
	for i := 0; i < n/2; i++ {
		b[i], b[n-1-i] = b[n-1-i], b[i]
	}
	return string(b);
}

func revert_string_p(str *string) {
	b := []byte(*str)
	n := len(b)
	for i := 0; i < n/2; i++ {
		b[i], b[n-1-i] = b[n-1-i], b[i]
	}
	*str = string(b);
}

func main() {
	fmt.Println("revert go project")

	inP := bufio.NewReader(os.Stdin)
	fmt.Println("please input your content:")

	str, err := inP.ReadString('\n')
	checkError(err)
	fmt.Println("your input is:", str)
	str_v := revert_string(str)
	fmt.Println("revert str is:", str_v)
	revert_string_p(&str)
	fmt.Println("your input:", str, "revert:", str)
}
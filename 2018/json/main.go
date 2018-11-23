package main

import (
	"fmt"
	"encoding/json"
)

type Test struct {
	Test1 uint64
	Test2 uint64
	Test3 string
	Test4 []byte
}

func main() {
	data, _ := json.Marshal(Test{
		Test1: 0,
		Test2: 12,
		Test3: "2",
		Test4: nil,
	})
	fmt.Println(string(data))
}

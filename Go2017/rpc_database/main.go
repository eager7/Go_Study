package main

import (
	"fmt"
	"encoding/json"
)

//import "./remote"

func main() {
	//remote.ServerThrift()


	test()
}

func duplicate(ll interface{}) (r interface{}) {
	m := make(map[interface{}]interface{})
	switch t := ll.(type) {
	case []int:
		for _, l := range t {
			m[l] = l
		}
		var list []int
		for i := 0; i < len(m); i++ {
			v, _ := m[i].(int)
			list = append(list, v)
		}
		return list
	default:
		fmt.Println("unsupport type")
	}
	return nil
}

func test(){
	type j struct {
		Name string
		Age int
	}
	p := j{"pct", 30}
	fmt.Println(p)
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", b)
}
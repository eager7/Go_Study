package main

import (
	"fmt"
	"math/big"
	"encoding/json"
)

func main() {
	testJson()
}

func testCPU() {

}

type Envelope struct {
	Type string
	Msg interface{}
}

type Sound struct {
	Des string
	Aut string
}

type Cowbell struct {
	More bool
}

func testJson() {
	s := Envelope {
		Type:"sound",
		Msg: Sound{
			Des:"des",
			Aut:"aut",
		},
	}
	buf, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var env Envelope
	if err := json.Unmarshal(buf, &env); err != nil {
		panic(err)
	}
	var desc string = env.Msg.(map[string]interface{})["Des"].(string)
	fmt.Println(desc)
}

func test2() {
	n := new(big.Int)
	n.SetInt64(-200)
	fmt.Println(n)

	b := n.Bytes()
	bb ,_ := n.GobEncode()
	fmt.Println(b)
	fmt.Println(bb)

	m := new(big.Int).SetBytes(b)
	fmt.Println(m)
	mm := new(big.Int)
	mm.GobDecode(bb)
	fmt.Println(mm)
}


func test() {
	t := newT2()
	t.t.val = 2
	fmt.Println(t)
}


type t1 struct {
	val int
}

type t2 struct {
	t *t1
}

func newT2() *t2{
	t := t2{}
	t.t = &t1{}
	return &t
}
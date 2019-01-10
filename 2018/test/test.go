package main

import (
	"encoding/json"
	"fmt"
	"github.com/ecoball/go-ecoball/common"
	"math/big"
	"sort"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

type AbaBftData struct {
	NumberRound        uint32
	PerBlockSignatures []common.Signature
}

type Mutex struct {
	m sync.Mutex
}

func Defer() {
	fmt.Println("defer")
}

func main() {
	swc(12.6783)
	//mapGet()
}

func mapGet() {
	m := make(map[string]string, 0)
	m["pct"] = "pct"
	var value string
	if v, ok := m["pct"]; ok {
		value = v
	}
	fmt.Println(value, m)
	value = "test"
	fmt.Println(value, m)
}

func swc(value float64) {
	timeNow := time.Now().Unix()
	var num int64 = 0
	for {
		if value > 2.0 {
			value = value - value*0.01
			num++
		} else {
			break
		}
	}
	t := timeNow + num*600
	fmt.Println(time.Unix(t, 0))
}

func slice() {
	l := []byte("test")
	for k, v := range l {
		fmt.Println(k, v)
	}
}

func SliceTest() {
	list := make(map[uint64]uint64, 1)
	list[1] = 3
	list[2] = 2
	list[3] = 1

	fmt.Println(list)
	var keys []uint64
	for k := range list {
		keys = append(keys, k)
	}
	fmt.Println("keys", keys)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	fmt.Println("keys", keys)
	for _, v := range keys {
		fmt.Println(list[v])
	}
}

func StringToUint64(str string) uint64 {
	strPointerInt := fmt.Sprintf("%d", unsafe.Pointer(&str))
	value, err := strconv.ParseUint(strPointerInt, 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func Uint64ToString(value uint64) string {
	var s *string
	s = *(**string)(unsafe.Pointer(&value))
	str := *(*string)(unsafe.Pointer(s))
	return string([]byte(str))
}

func testSlice() {
	list := make([][]byte, 2)
	list[0] = []byte("1")
	list[1] = []byte("2")
	fmt.Println(list)
}

type Envelope struct {
	Type string
	Msg  interface{}
}

type Sound struct {
	Des string `json:"Des"`
	Aut string `json:"Aut"`
}

type Cowbell struct {
	More bool
}
type Resource struct {
	Ram struct {
		Quota float32 `json:"quota"`
		Used  float32 `json:"used"`
	}
	Net struct {
		Staked    float32 `json:"staked"`
		Used      float32 `json:"used"`
		Available float32 `json:"available"`
		Limit     float32 `json:"limit"`
	}
	Cpu struct {
		Staked    float32 `json:"staked"`
		Used      float32 `json:"used"`
		Available float32 `json:"available"`
		Limit     float32 `json:"limit"`
	}
}

func testJson() {
	s := Envelope{
		Type: "sound",
		Msg: Sound{
			Des: "des",
			Aut: "aut",
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

func testJsonString() {
	map1 := make(map[Sound]interface{})
	sound1 := Sound{"sound1", "value"}
	sound2 := Sound{"sound2", "value"}
	map1[sound1] = "hello"
	map1[sound2] = "world"
	//return []byte
	str, err := json.Marshal(map1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(str))
	}
}

func test2() {
	n := new(big.Int)
	n.SetInt64(-200)
	fmt.Println(n)

	b := n.Bytes()
	bb, _ := n.GobEncode()
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

func newT2() *t2 {
	t := t2{}
	t.t = &t1{}
	return &t
}

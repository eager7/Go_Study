package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"reflect"
	"time"
)



func sliceCopy() {
	s := make([]int, 0)
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println("1", s)
	ss := make([]int, len(s))
	copy(ss, s)
	go func() {
		test(ss)
	}()
	s = make([]int, 0)
	time.Sleep(time.Second * 1)
	fmt.Println("3", s)
}

func test(s []int) {
	fmt.Println("2", s)
}

func sliceChange(data [5]int) {
	data[0] = 1
}

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func getContext(contextFilePath string) (string, int64) {
	var lastRunEndDate string
	var startIndex int64
	actionIndexFile, e := os.Open(contextFilePath)
	if e != nil {
		log.Println("can't read action table index:", e.Error())
		return lastRunEndDate, startIndex
	}
	defer actionIndexFile.Close()

	scanner := bufio.NewScanner(actionIndexFile)
	for scanner.Scan() {
		indexLine := scanner.Text()
		fmt.Sscanf(indexLine, "%s %d", &lastRunEndDate, &startIndex)

		log.Printf("date: %s, index: %d", lastRunEndDate, startIndex)
		break
	}

	if lastRunEndDate == "" || startIndex == 0 {
		log.Println("fail to read executed context")
	}

	return lastRunEndDate, startIndex
}

func main() {
	fmt.Println("start example...")
	var m string
	GetMacAddress(&m)
	fmt.Println("m:", m)
}
func GetMacAddress(macAddr *string) {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		//fmt.Println(inter.Name)
		if "en0" == inter.Name {
			mac := inter.HardwareAddr //获取本机MAC地址
			fmt.Println("mac = ", reflect.TypeOf(mac))
			*macAddr = string([]byte(mac))
			fmt.Println("macAddr = ", *macAddr)
		}
	}
}
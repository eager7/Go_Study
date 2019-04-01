package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("start example...")
	HexAndBigInt()
}

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

func HexFormat(s string) string {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return s
}

func HexToUint64(hex string) (uint64, error) {
	n, err := strconv.ParseUint(HexFormat(hex), 16, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func HexAndBigInt() {
	h := "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	hn, err := HexToUint64(h)
	if err != nil {
		panic(err)
	}
	fmt.Printf("hex:%d\n", hn)
	c := 18446744073709551615 - 115792089237316266660066408626602828282606886466848266086008062602462446642046
	fmt.Println("c:", c)
}

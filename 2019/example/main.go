package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("start example...")
	fmt.Println(getContext("/tmp/data.dat"))
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
package main

import (
	"bufio"
	"fmt"
	"github.com/eager7/elog"
	"hash/crc32"
	"math/big"
	"net"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"
)

var log = elog.NewLogger("example", elog.DebugLevel)

func main() {
	log.Debug("start example...")
	select {}
}

func Rou() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("finished")
}

func CRC32(s string) uint32 {
	ieee := crc32.NewIEEE()
	_, _ = ieee.Write([]byte(s))
	return ieee.Sum32()
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
		fmt.Println("can't read action table index:", e.Error())
		return lastRunEndDate, startIndex
	}
	defer actionIndexFile.Close()

	scanner := bufio.NewScanner(actionIndexFile)
	for scanner.Scan() {
		indexLine := scanner.Text()
		fmt.Sscanf(indexLine, "%s %d", &lastRunEndDate, &startIndex)

		fmt.Printf("date: %s, index: %d", lastRunEndDate, startIndex)
		break
	}

	if lastRunEndDate == "" || startIndex == 0 {
		fmt.Println("fail to read executed context")
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
	h := "20a4768c99456c4358c"
	b, f := new(big.Int).SetString(h, 16)
	fmt.Println("hex:", b, f)
	fmt.Printf("%d\n", b)
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

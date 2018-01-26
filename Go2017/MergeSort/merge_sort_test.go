package MergeSort

import (
	"fmt"
	"math/rand"
	"testing"
	"github.com/eager7/go/mlog"
)
var list []int
var listBenchmark []int
func init() {
	list = make([]int, 0)
	for i := 0; i < 100; i ++{
		list = append(list, rand.Intn(100))
	}
	fmt.Println("before:", list)

	const length = 1000 * 10000
	listBenchmark = make([]int, length)
	for i := 0; i < length; i ++{
		listBenchmark[i] = rand.Intn(length)
	}
}

func TestMergeSort(t *testing.T) {
	list1 := MergeSort(list)
	mlog.Notice.Println("after:", list1)
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(listBenchmark)
	}
}

func TestMergeSortGo(t *testing.T) {
	chRet := make(chan []int)
	go MergeSortGo(list, chRet)
	mlog.Debug.Println("after:", <-chRet)
}


func BenchmarkMergeSortGo(b *testing.B) {
	chRet := make(chan []int)
	for i := 0; i < b.N; i++ {
		go MergeSortGo(listBenchmark, chRet)
		<-chRet
	}
}


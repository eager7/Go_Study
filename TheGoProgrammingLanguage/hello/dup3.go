package main
import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	println("dup3 test")

	counts := make(map[string]int)
	for _, filename := range os.Args[1:]{
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n"){
			counts[line]++
		}
	}

	for line, n := range counts{
		fmt.Println(n, line)
	}
}
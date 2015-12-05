package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	println("dup2 test")

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			defer f.Close()
		}
	}

	for key, value := range counts {
		fmt.Println(key, value)
	}
}


func countLines(f *os.File, counts map[string]int){
	fmt.Println("countLines")
	input := bufio.NewScanner(f)
	for input.Scan(){
		if input.Text() == "q"{
			return
		} else {
			counts[input.Text()]++ 
		}
	}
}
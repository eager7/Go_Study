package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fmt.Println("dup testing 1")

	counts := make(map[string] int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++
		if input.Text() == "q"{
			break
		}
	}

	for line, n := range counts {
			fmt.Printf("%d\t%s\n", n, line)
	}
}
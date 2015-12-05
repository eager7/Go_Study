package main
import (
	"os"
	"io/ioutil"
	"net/http"
	"fmt"
)

func main(){
	fmt.Println("fetch test")

	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Println("fetch: reading", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)
	}
}
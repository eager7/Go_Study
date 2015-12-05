package main
import (
	"io"
	"net/http"
	"fmt"
	"strings"
	"os"
)

func main(){
	fmt.Println("fetch test")

	for _, url := range os.Args[1:] {
		
		isprefix := strings.HasPrefix(url, "http://")
		if isprefix != true {
			url = "http://" + url
		}

		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Get Body from", url)
		_, err = io.Copy(os.Stdout, response.Body)
		if err != nil {
			fmt.Println("fetch: reading", url, err)
			os.Exit(1)
		}
		fmt.Println(response.Status)
		response.Body.Close()
	}
}


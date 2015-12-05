package main
import (
	"fmt"
	"os"
	"time"
	"net/http"
	"io"
	"io/ioutil"
	"strings"
)

func main(){
	start := time.Now()

	ch := make(chan string)
	for _, url := range os.Args[1:] {
		isPre := strings.HasPrefix(url, "http://")
		if !isPre {
			url = "http://" + url
		}
		go fetch(url, ch)
	}

	for _, _ = range os.Args[1:]{
		fmt.Println(<-ch)
		fmt.Println("get chan")
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string){
	fmt.Println("fetch", url)
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

package main
import (
	"fmt"
	"net/http"
	"log"
	"sync"
)

var mu sync.Mutex
var count int

func main(){
	fmt.Println("server2 test")

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println("handler")
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request){
	fmt.Println("counter")
	fmt.Fprintf(w, "%q\n", r.URL.Path)
	mu.Lock()
	fmt.Fprintf(w, "the count of request is %d\n", count)
	mu.Unlock()
}
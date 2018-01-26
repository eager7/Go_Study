package main

import (
	"fmt"
	"net/http"
	. "github.com/eager7/go/log"
	"html/template"
)

func main(){
	Debug.Println("Web Test")
	http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request){
		fmt.Fprint(w, "hello world")
		req.ParseForm()
		Debug.Println("Server")
		Info.Println(req.Form)
		Info.Println(req.URL.Path)
		Info.Println(req.URL.Scheme)
		Info.Println(req.Form["url_long"])
		for k,v := range req.Form{
			Info.Println("key:", k)
			Info.Println("value:", v)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request){
		Success.Println("Method:", r.Method)
		r.ParseForm()
		if r.Method == "GET"{
			t,_ := template.ParseFiles("login.gtpl")
			t.Execute(w, nil)
		}else{
			Success.Println("name:", r.Form["username"])
			Success.Println("passwd:", r.Form["password"])
		}
	})
	Error.Fatal(http.ListenAndServe(":9090", nil))
}

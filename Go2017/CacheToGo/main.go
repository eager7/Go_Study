package main

import (
	"github.com/muesli/cache2go"
	"time"
	"fmt"
	"github.com/eager7/go/mlog"
)

type myStruct struct{
	text string
	moreData []byte
}

func main() {
	cache := cache2go.Cache("myCache")

	cache.SetLogger(mlog.Info)
	val := myStruct{"This is a test", []byte{}}
	cache.Add("someKey", time.Second*5, &val)

	res, err := cache.Value("someKey")
	if err == nil {
		fmt.Println(res.Data().(*myStruct).text)
	}else {
		mlog.Info.Println()
		mlog.Error.Println("Error retrieving value form cache:", err)
	}

	time.Sleep(time.Second*6)
	res, err = cache.Value("someKey")
	if err == nil {
		fmt.Println(res.Data().(*myStruct).text)
	}else{
		mlog.Error.Println(err)
	}

	cache.Add("someKey", time.Second*5, &val)
	cache.Delete("someKey")
	cache.Flush()

	if err == nil {
		fmt.Println(res.Data().(*myStruct).text)
	}else{
		mlog.Error.Println(err)
	}
}

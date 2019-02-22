package gorequest

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

var url = "https://forex.1forge.com/1.0.3/convert?from=USD&to=CNH&quantity=1&api_key=irSQh8AuW0x0Cu6JSbd8VkTGLBOoCdkn"

func Initialize() {
	resp, body, errs := gorequest.New().Timeout(time.Second*5).Get(url).End()
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(resp, body)
}

func GetCurrency(){
	resp := struct {
		Value     float64
		Text      string
		Timestamp int64
	}{}
	ret, _, errs := gorequest.New().Timeout(time.Second * 5).Get(url).EndStruct(&resp)
	if errs != nil || ret.StatusCode != http.StatusOK {
		panic(errs)
	} else {
		fmt.Println(resp)
	}
}


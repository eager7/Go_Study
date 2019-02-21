package gorequest

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

func Initialize() {
	resp, body, errs := gorequest.New().Timeout(time.Second*5).Get("https://forex.1forge.com/1.0.3/convert?from=USD&to=CNH&quantity=1&api_key=irSQh8AuW0x0Cu6JSbd8VkTGLBOoCdkn").End()
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(resp, body)
}





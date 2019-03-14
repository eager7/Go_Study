package gorequest

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

var url = "https://forex.1forge.com/1.0.3/convert?from=USD&to=CNH&quantity=1&api_key=irSQh8AuW0x0Cu6JSbd8VkTGLBOoCdkn"

func Initialize() {
	resp, body, errs := gorequest.New().Timeout(time.Second * 5).Get(url).End()
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(resp, body)
}

func GetCurrency() {
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

type TokenPrice struct {
	Symbol      string
	ExecAccount string
	EOSPrice    float64
	USDPrice    float64
	Source      string
	Timestamp   time.Time
}

func FetchTokenPriceListFromNewDex() ([]TokenPrice, error) {
	type tTokenTicker struct {
		Symbol   string  `json:"symbol"`
		Contract string  `json:"contract"`
		Currency string  `json:"currency"`
		Last     float64 `json:"last"`
		Change   float64 `json:"change"`
		High     float64 `json:"high"`
		Low      float64 `json:"low"`
		Amount   float64 `json:"amount"`
		Volume   float64 `json:"volume"`
	}
	type tResponse struct {
		Code int            `json:"code"`
		Data []tTokenTicker `json:"data"`
	}

	fmt.Println("begin get token list...", time.Now().UTC())
	var response tResponse
	resp, _, errs := gorequest.New().Get("https://api.newdex.io/v1/ticker/all").EndStruct(&response)
	if errs != nil {
		return nil, errs[0]
	}
	fmt.Println("end get token list...", time.Now().UTC())
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("fail to fetch token price from newdex. status: " + resp.Status)
	}

	now := time.Now()
	priceList := make([]TokenPrice, len(response.Data))
	for i, ticker := range response.Data {
		priceList[i] = TokenPrice{
			Symbol:      ticker.Currency,
			ExecAccount: ticker.Contract,
			EOSPrice:    ticker.Last,
			Source:      "newdex",
			Timestamp:   now,
		}
	}
	return priceList, nil
}

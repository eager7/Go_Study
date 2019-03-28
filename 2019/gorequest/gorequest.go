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
		//Change   float64 `json:"change"`
		//High     float64 `json:"high"`
		//Low      float64 `json:"low"`
		//Amount   float64 `json:"amount"`
		//Volume   float64 `json:"volume"`
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

func GetEthChainId() {
	resp := struct {
		JsonRpc string `json:"jsonrpc"`
		Id      int32  `json:"id"`
		Result  string `json:"result"`
	}{}
	ret, body, errs := gorequest.New().Timeout(time.Second*5).
		Post("http://47.52.157.31:8545").
		Set("Content-Type", "application/json").
		Send(`{"jsonrpc":"2.0","method":"net_version","params":[],"id":67}`).EndStruct(&resp)
	if errs != nil || ret.StatusCode != http.StatusOK {
		fmt.Println("request chain id error:", errs)
	}
	fmt.Println(resp)
	fmt.Println("body:", body)
}

func GetEtherScanTrxList() {
	type Trx struct {
		BlockNumber       string `json:"blockNumber"`
		BlockHash         string `json:"blockHash"`
		TimeStamp         string `json:"timeStamp"`
		Hash              string `json:"hash"`
		Nonce             string `json:"nonce"`
		TransactionIndex  string `json:"transactionIndex"`
		From              string `json:"from"`
		To                string `json:"to"`
		Value             string `json:"value"`
		Gas               string `json:"gas"`
		GasPrice          string `json:"gasPrice"`
		Input             string `json:"input"`
		ContractAddress   string `json:"contractAddress"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		TxReceipt         string `json:"txreceipt_status"`
		GasUsed           string `json:"gasUsed"`
		ConfirmAtions     string `json:"confirmations"`
		IsError           string `json:"isError"`
	}
	type Resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  []Trx  `json:"result"`
	}
	r := Resp{}
	res := gorequest.New()
	ret, body, errs := res.Timeout(time.Second*10).Get("http://api.etherscan.io/api").
		Set("module", "account").Set("action", "txlist").Set("address", "0x1BcF4912fEecad5c1F7c74705273A562fA6002e5").
		Set("startblock", "0").Set("endblock", "99999999").Set("sort", "asc").Set("apikey", "4EVHX9TZ7EPWY8MMFANBJMHTH6FGSH81K9").EndStruct(&r)
	if errs != nil || ret.StatusCode != http.StatusOK {
		req, err := res.MakeRequest()
		if err == nil {
			fmt.Printf("request body:%+v\n", req)
		}
		fmt.Println(errs, string(body))
		return
	}
	fmt.Printf("%+v", r)
}

func Example() {
	res := gorequest.New()
	ret, body, errs := res.Timeout(time.Second*20).Get("http://api.etherscan.io/api?module=account&action=txlist&address=0x1BcF4912fEecad5c1F7c74705273A562fA6002e5&startblock=0&endblock=99999999&sort=asc&apikey=4EVHX9TZ7EPWY8MMFANBJMHTH6FGSH81K9").End()
	if errs != nil || ret.StatusCode != http.StatusOK {
		req, err := res.MakeRequest()
		if err == nil {
			fmt.Printf("request body:%+v\n", req)
		}
		fmt.Println(errs, string(body))
		return
	}
	fmt.Println("resp:", body)
}
package gorequest

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"testing"
	"time"
)

func TestInitialize(t *testing.T) {
	GetCurrency()
}

func TestFetchTokenPriceListFromNewDex(t *testing.T) {
	if list, err := FetchTokenPriceListFromNewDex(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(list)
	}
}

func TestGetEthChainId(t *testing.T) {
	fmt.Println(fmt.Sprintf("%v", true))
	GetEthChainId()
}

func TestGetEtherScanTrxList(t *testing.T) {
	GetEtherScanTrxList()
}

func TestExample(t *testing.T) {
	Example()
}

func TestCK(t *testing.T) {
	fmt.Println("test ck")
	url := `https://api.cryptokitties.co/kitties/%d`
	for i := 1; i < 200; i++ {
		fmt.Println("url:", fmt.Sprintf(url, i))
		resp, body, errs := gorequest.New().Timeout(time.Second * 10).Get(fmt.Sprintf(url, i)).End()
		if errs != nil || resp.StatusCode != http.StatusOK {
			t.Log(resp.StatusCode)
			t.Fatal(errs)
		}
		fmt.Println("body:", string(body[:10]))
	}
}

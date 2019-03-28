package gorequest

import (
	"fmt"
	"testing"
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
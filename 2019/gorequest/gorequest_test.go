package gorequest

import "testing"

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
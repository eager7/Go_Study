package dapp_sheild

import (
	"testing"
)

func TestRequestContractList(t *testing.T) {
	resp, err := RequestContractList()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	//HttpRequest()

}

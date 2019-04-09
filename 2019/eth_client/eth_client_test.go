package eth_client

import (
	"fmt"
	"github.com/BlockABC/wallet_eth_client/common/context"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestInitialize(t *testing.T) {
	eth, err := new(Eth).Initialize(new(context.Context).Initialize(), "wss://mainnet.infura.io/ws") //"wss://ropsten.infura.io/ws" "https://mainnet.infura.io"
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(eth.BalanceAt("0x06012c8cf97BEaD5deAe237070F9587f8E7A266d", nil))
	go eth.SubscribeNewHeader()
	make(chan interface{}) <- struct{}{}
}

func TestRun(t *testing.T) {
	fmt.Println(new(common.Address).Hex())
}

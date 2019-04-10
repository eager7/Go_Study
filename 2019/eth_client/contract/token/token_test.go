package token_test

import (
	"fmt"
	"github.com/BlockABC/wallet_eth_client/common/context"
	"github.com/eager7/go_study/2019/eth_client"
	"github.com/eager7/go_study/2019/eth_client/contract/token"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"testing"
)

func TestTokenInfo(t *testing.T) {
	eth, err := new(eth_client.Eth).Initialize(new(context.Context).Initialize(), "wss://mainnet.infura.io/ws") //"wss://ropsten.infura.io/ws" "https://mainnet.infura.io"
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token.ReadTokenInfo("0xB8c77482e45F1F44dE1745F52C74426C631bDD52", eth.Client()))
}

func TestListenTransferEvent(t *testing.T) {
	eth, err := new(eth_client.Eth).Initialize(new(context.Context).Initialize(), "wss://mainnet.infura.io/ws") //"wss://ropsten.infura.io/ws" "https://mainnet.infura.io"
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token.ListenTransferEvent(eth.CTX().Context(), eth.Client()), "0xB8c77482e45F1F44dE1745F52C74426C631bDD52")
}

func TestAbi(t *testing.T) {
	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(contractAbi.Events)
}
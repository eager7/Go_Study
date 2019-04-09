package eth_client

import (
	"fmt"
	"github.com/BlockABC/wallet-webserver/common/utils"
	"github.com/BlockABC/wallet_eth_client/common/context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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

func TestEth_GetChainDataByNumber(t *testing.T) {
	eth, err := new(Eth).Initialize(new(context.Context).Initialize(), "wss://mainnet.infura.io/ws") //"wss://ropsten.infura.io/ws" "https://mainnet.infura.io"
	if err != nil {
		t.Fatal(err)
	}
	chain, err := eth.GetChainDataByNumber(new(big.Int).SetUint64(7533908))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(chain.Transactions), len(chain.Tokens), len(chain.Events), len(chain.Accounts), len(chain.Contracts), len(chain.Asserts))
	fmt.Println(utils.JsonString(chain.Events))
}

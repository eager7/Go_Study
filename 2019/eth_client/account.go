package eth_client

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (e *Eth) BalanceAt(address string, number *big.Int) (*big.Int, error) {
	return e.client.BalanceAt(e.ctx.Context(), common.HexToAddress(address), number)
}

func (e *Eth) PendingBalanceAt(address string) (*big.Int, error) {
	return e.client.PendingBalanceAt(e.ctx.Context(), common.HexToAddress(address))
}

func (e *Eth) PendingNonceAt(address string) (uint64, error) {
	return e.client.PendingNonceAt(e.ctx.Context(), common.HexToAddress(address))
}

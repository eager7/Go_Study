package eth_client

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"regexp"
)

func (e *Eth) IsValidAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func (e *Eth) IsContractAddress(address string, number *big.Int) (bool, error) {
	code, err := e.client.CodeAt(e.ctx.Context(), common.HexToAddress(address), number)
	if err != nil {
		return false, errors.New(fmt.Sprintf("get address' code err:%v", err))
	}
	return len(code) > 0, nil
}


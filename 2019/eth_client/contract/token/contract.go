package token

import (
	"context"
	"errors"
	"fmt"
	"github.com/BlockABC/wallet_eth_client/common/evm"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

func ReadTokenInfo(address string, client *ethclient.Client) (string, string, uint8, *big.Int, error) {
	instance, err := NewToken(common.HexToAddress(address), client)
	if err != nil {
		return "", "", 0, nil, errors.New(fmt.Sprintf("new erc20 err:%v", err))
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		fmt.Println("can't get name:", err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, nil, errors.New(fmt.Sprintf("symbol err:%v", err))
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		fmt.Println("decimals err:", err)
	}
	supply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, nil, errors.New(fmt.Sprintf("supply err:%v", err))
	}
	return name, symbol, decimals, supply, nil
}

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func ListenTransferEvent(ctx context.Context, client *ethclient.Client, address ...common.Address) error {
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: address,
	}
	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(TokenABI)))
	if err != nil {
		return err
	}
	logTransferSigHash := evm.EIP165Event("Transfer(address,address,uint256)")
	for _, l := range logs {
		switch l.Topics[0].Hex() {
		case logTransferSigHash:
			fmt.Printf("Log Name: Transfer\n")
			var transferEvent LogTransfer
			err := contractAbi.Unpack(&transferEvent, "Transfer", l.Data)
			if err != nil {
				return err
			}
			transferEvent.From = common.HexToAddress(l.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(l.Topics[2].Hex())
			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())
		default:
			fmt.Println("unknown topic:", l.Topics[0].Hex())
		}
	}
	return nil
}

func AnalysisLogs() {

}

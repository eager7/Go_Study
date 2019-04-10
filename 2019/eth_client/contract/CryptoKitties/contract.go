package CryptoKitties

import (
	"context"
	"fmt"
	"github.com/BlockABC/wallet_eth_client/common/evm"
	"github.com/BlockABC/wallet_eth_client/common/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

type LogTransfer struct {
	From   common.Address `json:"from"`
	To     common.Address `json:"to"`
	Tokens *big.Int       `json:"tokens"`
}

type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func ListenTransferEvent(ctx context.Context, client *ethclient.Client, address ...common.Address) error {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(7537735),
		ToBlock:   nil,
		Addresses: address,
	}
	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(CryptoKittiesABI)))
	if err != nil {
		return err
	}
	logTransferSigHash := evm.EIP165Event("Transfer(address,address,uint256)")
	for _, l := range logs {
		switch utils.HexFormat(l.Topics[0].Hex()) {
		case logTransferSigHash:
			fmt.Println("Log Name: Transfer\n", utils.JsonString(l))
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

func AnalysisLogs(logs ...types.Log) error {
	contractAbi, err := abi.JSON(strings.NewReader(string(CryptoKittiesABI)))
	if err != nil {
		return err
	}
	logTransferSigHash := evm.EIP165Event("Transfer(address,address,uint256)")
	for _, l := range logs {
		switch utils.HexFormat(l.Topics[0].Hex()) {
		case logTransferSigHash:
			fmt.Println("Log Name: Transfer\n", utils.JsonString(l))
			if len(l.Data) == 0 && len(l.Topics) == 4 {
				fmt.Println("from:", l.Topics[1].Hex())
				fmt.Println("to:", l.Topics[2].Hex())
				fmt.Println("value:", l.Topics[3].Big())
			} else {
				var transferEvent LogTransfer
				err := contractAbi.Unpack(&transferEvent, "Transfer", l.Data)
				if err != nil {
					return err
				}
				transferEvent.From = common.HexToAddress(l.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(l.Topics[2].Hex())
				fmt.Printf("From: %s\n", transferEvent.From.Hex())
				fmt.Printf("To: %s\n", transferEvent.To.Hex())
				fmt.Printf("Tokens: %s\n", transferEvent.Tokens)
			}

		default:
			fmt.Println("unknown topic:", l.Topics[0].Hex())
		}
	}
	return nil
}

package eth_client

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (e *Eth) HeaderByNumber(number *big.Int) (*types.Header, error) {
	return e.client.HeaderByNumber(e.ctx.Context(), number)
}

type Transaction struct {
	*types.Transaction
	Time    uint64
	From    string
	GasUsed uint64
	Index   uint16
}

func (e *Eth) BlockByNumber(number *big.Int) (*types.Block, error) {
	block, err := e.client.BlockByNumber(e.ctx.Context(), number)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get block by number err:%v", err))
	}
	return block, nil
}

func (e *Eth) BlockByHash(hash common.Hash) (*types.Block, error) {
	block, err := e.client.BlockByHash(e.ctx.Context(), hash)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get block by number err:%v", err))
	}
	return block, nil
}

func (e *Eth) AnalysisTransactions(block *types.Block) (map[common.Hash]*Transaction, error) {
	chainID, err := e.client.NetworkID(e.ctx.Context())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get chain id err:%v", err))
	}
	txs := make(map[common.Hash]*Transaction, 1)
	for index, tx := range block.Transactions() {
		t := Transaction{
			Transaction: block.Transaction(tx.Hash()),
			Time:        block.Time(),
			Index:       uint16(index),
		}
		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
			t.From = msg.From().Hex()
		}
		txs[tx.Hash()] = &t
	}
	return txs, err
}

func (e *Eth) TransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	return e.client.TransactionReceipt(e.ctx.Context(), hash)
}

func (e *Eth) AnalysisReceipts(block *types.Block) (map[common.Hash]*types.Receipt, error) {
	receipts := make(map[common.Hash]*types.Receipt, 1)
	for _, tx := range block.Transactions() {
		receipt, err := e.TransactionReceipt(tx.Hash())
		if err != nil {
			return nil, errors.New(fmt.Sprintf("analysis receipts err:%v", err))
		}
		receipts[tx.Hash()] = receipt
	}
	return receipts, nil
}

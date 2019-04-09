package eth_client

import (
	"errors"
	"fmt"
	"github.com/BlockABC/wallet_eth_client/common/evm"
	"github.com/BlockABC/wallet_eth_client/common/utils"
	"github.com/BlockABC/wallet_eth_client/database/tables"
	"github.com/eager7/go_study/2019/eth_client/contract/token"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

type ChainData struct {
	Block        *tables.TableBlockInfo
	Accounts     []*tables.TableAccountInfo
	Asserts      []*tables.TableAssertsInfo
	Contracts    []*tables.TableContractInfo
	Events       []*tables.TableEventInfo
	Tokens       []*tables.TableTokenInfo
	Transactions []*tables.TableTransactionInfo
}

func (e *Eth) GetChainDataByNumber(number *big.Int) (*ChainData, error) {
	chain := ChainData{}
	block, err := e.BlockByNumber(number)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get chain block err:%v", err))
	}
	chain.FormatBlock(block)
	txs, err := e.AnalysisTransactions(block)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get chain txs err:%v", err))
	}
	receipts, err := e.AnalysisReceipts(block)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get chain receipts err:%v", err))
	}
	chain.FormatTransactions(txs, receipts)
	if err := chain.FormatReceipts(e, txs, receipts); err != nil {
		return nil, errors.New(fmt.Sprintf("get chain receipts format err:%v", err))
	}
	return &chain, nil
}

func (c *ChainData) FormatBlock(block *types.Block) {
	c.Block = &tables.TableBlockInfo{
		Number:           block.Number().Uint64(),
		Difficulty:       utils.BigIntToHex(block.Difficulty()),
		ExtraData:        utils.ByteToHex(block.Extra()),
		GasLimit:         block.GasLimit(),
		GasUsed:          block.GasUsed(),
		Hash:             utils.HexFormat(block.Hash().Hex()),
		LogsBloom:        utils.ByteToHex(block.Bloom().Bytes()),
		Miner:            utils.HexFormat(block.Coinbase().Hex()),
		MixHash:          utils.HexFormat(block.MixDigest().Hex()),
		Nonce:            block.Nonce(),
		ParentHash:       utils.HexFormat(block.ParentHash().Hex()),
		ReceiptsRoot:     utils.HexFormat(block.ReceiptHash().Hex()),
		Sha3Uncles:       utils.HexFormat(block.UncleHash().Hex()),
		Size:             block.Size().String(),
		StateRoot:        utils.HexFormat(block.Root().Hex()),
		Timestamp:        block.Time(),
		TotalDifficulty:  utils.BigIntToHex(block.DeprecatedTd()),
		TransactionsNum:  uint16(block.Transactions().Len()),
		TransactionsRoot: utils.HexFormat(block.TxHash().Hex()),
	}
	for _, h := range block.Uncles() {
		c.Block.UnclesHash += utils.HexFormat(h.Hash().Hex())
	}
}

func (c *ChainData) FormatTransactions(txs map[common.Hash]*Transaction, receipts map[common.Hash]*types.Receipt) {
	if c.Block == nil {
		return
	}
	for _, tx := range txs {
		v, s, r := tx.RawSignatureValues()
		txInfo := tables.TableTransactionInfo{
			BlockHash:        c.Block.Hash,
			BlockNumber:      c.Block.Number,
			Timestamp:        c.Block.Timestamp,
			From:             tx.From,
			Gas:              tx.Gas(),
			GasUsed:          receipts[tx.Hash()].GasUsed,
			GasPrice:         utils.BigIntToHex(tx.GasPrice()),
			Hash:             utils.HexFormat(tx.Hash().Hex()),
			InputFlag:        1,
			Input:            utils.ByteToHex(tx.Data()),
			Nonce:            tx.Nonce(),
			TransactionIndex: uint16(receipts[tx.Hash()].TransactionIndex),
			Value:            utils.BigIntToHex(tx.Value()),
			V:                utils.BigIntToHex(v),
			S:                utils.BigIntToHex(s),
			R:                utils.BigIntToHex(r),
		}
		if tx.To() != nil {
			txInfo.To = utils.HexFormat(tx.To().Hex())
		}
		//合约代码存储到合约表单
		if receipts[tx.Hash()].ContractAddress.Hex() != new(common.Address).Hex() {
			txInfo.InputFlag = 2
			txInfo.Input = ""
		}
		c.Transactions = append(c.Transactions, &txInfo)
	}
}

func (c *ChainData) FormatReceipts(e *Eth, txs map[common.Hash]*Transaction, receipts map[common.Hash]*types.Receipt) error {
	for _, receipt := range receipts {
		if receipt.ContractAddress.Hex() != new(common.Address).Hex() { //创建合约交易
			contract := tables.TableContractInfo{
				Address:     utils.HexFormat(receipt.ContractAddress.Hex()),
				Creator:     utils.HexFormat(txs[receipt.TxHash].From),
				Timestamp:   c.Block.Timestamp,
				BlockNumber: c.Block.Number,
				Transaction: utils.HexFormat(receipt.TxHash.Hex()),
				Code:        utils.ByteToHex(txs[receipt.TxHash].Data()),
			}
			c.Contracts = append(c.Contracts, &contract)
			//根据合约二进制代码检测合约是否为代币合约，如果是代币合约则放入代币表单中
			if t := evm.CheckTokenInterface(contract.Code); t != evm.None {
				name, symbol, decimals, supply, err := token.ReadTokenInfo(receipt.ContractAddress.Hex(), e.Client())
				if err != nil {
					return errors.New(fmt.Sprintf("get chain data err:%s", err.Error()))
				}
				tok := tables.TableTokenInfo{
					Address:     utils.HexFormat(receipt.ContractAddress.Hex()),
					Type:        uint8(t),
					Name:        name,
					Symbol:      symbol,
					SymbolValue: utils.CRC32(symbol),
					Decimals:    decimals,
					Timestamp:   c.Block.Timestamp,
					Supply:      utils.BigIntToHex(supply),
					Issue:       "",
				}
				c.Tokens = append(c.Tokens, &tok)
			}
		} else { //普通交易或者调用合约交易，事件会放到日志的topics中，普通转账交易则没有日志
			contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
			if err != nil {
				return errors.New(fmt.Sprintf("create abi err:%v", err))
			}
			for _, logger := range receipt.Logs {
				switch utils.HexFormat(logger.Topics[0].Hex()) {
				case evm.EIP165Event("Transfer(address,address,uint256)"):
					event := tables.TableEventInfo{
						Address:          utils.HexFormat(logger.Address.Hex()),
						BlockHash:        utils.HexFormat(logger.BlockHash.Hex()),
						BlockNumber:      c.Block.Number,
						Timestamp:        c.Block.Timestamp,
						Hash:             utils.HexFormat(logger.TxHash.Hex()),
						InputFlag:        1,
						Input:            utils.ByteToHex(logger.Data),
						TransactionIndex: uint16(logger.TxIndex),
					}
					var transferEvent token.LogTransfer
					err := contractAbi.Unpack(&transferEvent, "Transfer", logger.Data)
					if err != nil {
						fmt.Println("can't unpack log data:", len(logger.Topics), utils.JsonString(logger))
						return errors.New(fmt.Sprintf("abi unpack err:%v", err))
					}
					if transferEvent.From.Hex() == new(common.Address).Hex() && len(logger.Topics) == 3 { //符合标准协议的ERC20 Transfer事件-event Transfer(address indexed _from, address indexed _to, uint256 _value)
						transferEvent.From = common.HexToAddress(logger.Topics[1].Hex())
						transferEvent.To = common.HexToAddress(logger.Topics[2].Hex())
					} else if transferEvent.From.Hex() == new(common.Address).Hex() && len(logger.Topics) == 4 { //符合标准协议的ERC721 Transfer事件event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId);
						transferEvent.From = common.HexToAddress(logger.Topics[1].Hex())
						transferEvent.To = common.HexToAddress(logger.Topics[2].Hex())
						transferEvent.Value = logger.Topics[3].Big()
					} else {
						fmt.Println("can't find handle func:", len(logger.Topics), utils.JsonString(logger))
						continue
					}
					event.From = utils.HexFormat(transferEvent.From.Hex())
					event.To = utils.HexFormat(transferEvent.To.Hex())
					event.Value = utils.BigIntToHex(transferEvent.Value)
					c.Events = append(c.Events, &event)
					if err := c.GetAssertInfo(e, logger.Address.Hex(), transferEvent.From.Hex(), transferEvent.To.Hex()); err != nil {
						return errors.New(fmt.Sprintf("get assert data err:%s", err.Error()))
					}
				}
				if err := c.GetAccountInfo(e, txs[receipt.TxHash].From, txs[receipt.TxHash].To().Hex()); err != nil {
					return errors.New(fmt.Sprintf("get account data err:%s", err.Error()))
				}
			}
		}
	}
	return nil
}

func (c *ChainData) GetAccountInfo(e *Eth, address ...string) error {
	for _, addr := range address {
		balance, err := e.BalanceAt(addr, nil)
		if err != nil {
			return err
		}
		nonce, err := e.PendingNonceAt(addr)
		if err != nil {
			return err
		}
		acc := tables.TableAccountInfo{
			Address:   utils.HexFormat(addr),
			Nonce:     nonce,
			Balance:   utils.BigIntToHex(balance),
			Timestamp: c.Block.Timestamp,
			Type:      1,
		}
		c.Accounts = append(c.Accounts, &acc)
	}
	return nil
}

func (c *ChainData) GetAssertInfo(e *Eth, contract string, address ...string) error {
	for _, addr := range address {
		v, err := token.BalanceAt(addr, contract, e.client)
		if err != nil {
			return err
		}
		assert := tables.TableAssertsInfo{
			Address:  utils.HexFormat(addr),
			Balance:  utils.BigIntToHex(v),
			Contract: utils.HexFormat(contract),
		}
		c.Asserts = append(c.Asserts, &assert)
	}
	return nil
}

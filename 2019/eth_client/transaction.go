package eth_client

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/BlockABC/wallet_eth_client/common/evm"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

func (e *Eth) SendTransaction(private, to string, amount *big.Int, gasLimit uint64, data []byte) error {
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		return errors.New(fmt.Sprintf("send transaction private err:%v", err))
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := e.PendingNonceAt(from.Hex())
	if err != nil {
		return errors.New(fmt.Sprintf("send transaction nonce err:%v", err))
	}
	gasPrice, err := e.client.SuggestGasPrice(e.ctx.Context())
	if err != nil {
		return errors.New(fmt.Sprintf("send transaction gas price err:%v", err))
	}
	tx := types.NewTransaction(nonce, common.HexToAddress(to), amount, gasLimit, gasPrice, data)
	chainID, err := e.client.NetworkID(e.ctx.Context())
	if err != nil {
		return errors.New(fmt.Sprintf("send transaction chain id err:%v", err))
	}
	sigTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return errors.New(fmt.Sprintf("send transaction sign err:%v", err))
	}
	return e.client.SendTransaction(e.ctx.Context(), sigTx)
}

func (e *Eth) SendErcTransaction(private, contract, to string, amount *big.Int, gasLimit uint64, gasPrice *big.Int) error {
	var data []byte
	toAddr := common.HexToAddress(to)

	methodID := evm.EIP165Method("transfer(address,uint256)")
	data = append(data, methodID...)

	paddedAddress := common.LeftPadBytes(toAddr.Bytes(), 32)
	data = append(data, paddedAddress...)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	data = append(data, paddedAmount...)

	gasLimit, err := e.client.EstimateGas(e.ctx.Context(), ethereum.CallMsg{To: &toAddr, Data: data})
	if err != nil {
		return errors.New(fmt.Sprintf("send transaction estimate gas err:%v", err))
	}
	return e.SendTransaction(private, contract, big.NewInt(0), gasLimit, data)
}

func (e *Eth) TransactionToRaw(tx *types.Transaction) string {
	return hex.EncodeToString(types.Transactions{tx}.GetRlp(0))
}

func (e *Eth) TransactionFromRaw(rawTx string) (*types.Transaction, error) {
	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("transaction from raw err:%v", err))
	}
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(rawTxBytes, &tx); err != nil {
		return nil, errors.New(fmt.Sprintf("transaction from raw rlp err:%v", err))
	}
	return tx, nil
}

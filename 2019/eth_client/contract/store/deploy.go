package store

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func DeployContract(ctx context.Context, client *ethclient.Client, private string) (*common.Address, *types.Transaction, *Store, error) {
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		return nil, nil, nil, errors.New(fmt.Sprintf("send transaction private err:%v", err))
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, from)
	if err != nil {
		return nil, nil, nil, errors.New(fmt.Sprintf("send transaction nonce err:%v", err))
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, nil, nil, errors.New(fmt.Sprintf("send transaction gas price err:%v", err))
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := DeployStore(auth, client, "1.0")
	if err != nil {
		return nil, nil, nil, errors.New(fmt.Sprintf("deploy contract err:%v", err))
	}
	return &address, tx, instance, nil
}

func New(address common.Address, client *ethclient.Client) (*Store, error) {
	return NewStore(address, client)
}

func Version(address common.Address, client *ethclient.Client) (string, error) {
	instance, err := New(address, client)
	if err != nil {
		return "", err
	}
	return instance.Version(nil)
}

func SetItem(ctx context.Context, address common.Address, client *ethclient.Client, private string) (*types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("send transaction private err:%v", err))
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, from)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("send transaction nonce err:%v", err))
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("send transaction gas price err:%v", err))
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	instance, err := New(address, client)
	if err != nil {
		return nil, err
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))
	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

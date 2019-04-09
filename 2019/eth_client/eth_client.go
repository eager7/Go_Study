package eth_client

import (
	"errors"
	"fmt"
	"github.com/BlockABC/wallet_eth_client/common/context"
	"github.com/BlockABC/wallet_eth_client/common/elog"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var log = elog.Log

type Eth struct {
	url    string
	client *ethclient.Client
	ctx    *context.Context
}

func (e *Eth) Initialize(ctx *context.Context, url string) (*Eth, error) {
	e.ctx = ctx
	e.url = url
	client, err := ethclient.Dial(url) //"https://mainnet.infura.io"
	if err != nil {
		return nil, errors.New(fmt.Sprintf("initialize eth client err:%v", err))
	}
	e.client = client
	return e, nil
}

func (e *Eth) Client() *ethclient.Client {
	return e.client
}

func (e *Eth) CTX() *context.Context {
	return e.ctx
}

func (e *Eth) Close() {
	e.client.Close()
}

func (e *Eth) Reset(url string) error {
	if url == e.url {
		return nil
	}
	n, err := e.Initialize(e.ctx, url)
	if err != nil {
		return errors.New(fmt.Sprintf("reset eth error:%v", err))
	}
	e.Close()
	e.url = url
	e.client = n.client
	return nil
}

func (e *Eth) SubscribeNewHeader() {
	headers := make(chan *types.Header)
	sub, err := e.client.SubscribeNewHead(e.ctx.Context(), headers)
	if err != nil {
		panic(err)
	}
	e.ctx.Add()
	for {
		select {
		case err := <-sub.Err():
			fmt.Println("receive err signal:", err)
			return
		case header := <-headers:
			fmt.Println("receive new header:", header.Number, header.Hash().Hex())
		case <-e.ctx.Done():
			fmt.Println("stop header subscribe")
			e.ctx.TaskDone()
			return
		}
	}
}

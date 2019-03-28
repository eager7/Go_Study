package main

import (
	"encoding/json"
	"fmt"
)

const text = `{"jsonrpc":"2.0","id":1,"result":{"blockHash":"0xc58072a8a7e21325a2efd391f79010849e4d309c8c178099737a77c61f06c3d3","blockNumber":"0x3e9c97","from":"0x9e0b766e690d56c9a32194f43396a23b8ab39b70","gas":"0x186a0","gasPrice":"0x3b9aca00","hash":"0xb9658b9299a1698669ac2851299177dd78b1452783a91ea59c996f8b9a7a7d78","input":"0x","nonce":"0x11","to":"0x5ba1583b5fb624f66858386971ac9d5362ab5352","transactionIndex":"0x12","value":"0x38d7ea4c68000","v":"0x2b","r":"0xadb80657ab64eeb6a77fa13d63e812e5930ed40a0bf8c1215be8744bf9ced0f2","s":"0x3c289c993e682dfa68ee8bb0943fd40ad5fd5cd97c84c8e674df15abc02d166e"}}`
const text2 = `{"status":"1","message":"OK","result":[{"blockNumber":"7287233","blockHash":"0xd56d64b16ece41c9b2be1e7abfca289a1ac47e0c155f6dcfa72e8ed93ec70c52","timeStamp":"1551489930","hash":"0x5c4095f690f98fed770574a75b024a232c827cbceae8232f4545a7da3294b7a0","nonce":"51392","transactionIndex":"5","from":"0xb9a4873d8d2c22e56b8574e8605644d08e047549","to":"0xdb1260b4b02c37e17cdc9b9663a11969cf1a5df0","value":"209084260000000000","gas":"90000","gasPrice":"30000000000","input":"0x","contractAddress":"","cumulativeGasUsed":"237750","txreceipt_status":"1","gasUsed":"21000","confirmations":"168353","isError":"0"},{"blockNumber":"7312974","blockHash":"0x79fca12ca672f6e425ffb02b46fbe89c4fc835fc0ef4da1582c6a862e7c4b1ed","timeStamp":"1551837726","hash":"0xa4f07c2c14496eaf364d8c0f0940b527ae8272ed1ade24f2b171b2668cf3ce29","nonce":"0","transactionIndex":"76","from":"0xdb1260b4b02c37e17cdc9b9663a11969cf1a5df0","to":"0x5dd70b8bfd8b753886fb44e89b8348b271743a9d","value":"208500000000000000","gas":"60000","gasPrice":"9000000000","input":"0x","contractAddress":"","cumulativeGasUsed":"7481687","txreceipt_status":"1","gasUsed":"21000","confirmations":"142612","isError":"0"}]}`
type Transaction struct {
	BlockHash        string `json:"block_hash"`
	BlockNum         string `json:"block_num"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gas_price"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transaction_index"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseResp struct {
	JsonRpc string      `json:"jsonrpc"`
	Id      int32       `json:"id"`
	Error   Error       `json:"error"`
}

type Trx struct {
	BlockNumber       string `json:"blockNumber"`
	BlockHash         string `json:"blockHash"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	TxReceipt         string `json:"txreceipt_status"`
	GasUsed           string `json:"gasUsed"`
	ConfirmAtions     string `json:"confirmations"`
	IsError           string `json:"isError"`
}
type Resp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []Trx  `json:"result"`
}

func main() {
	r := Resp{}
	json.Unmarshal([]byte(text2), &r)
	fmt.Printf("%+v", r)
}

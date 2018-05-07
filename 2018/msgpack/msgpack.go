package main

import (
	"crypto"
	"fmt"

	"github.com/vmihailenco/msgpack"
)

const HashLen = 32

type Hash [HashLen]byte

type Transaction struct {
	Version uint32
	TxType  uint32
	Nonce   uint32
	Payload []byte
	Sig     Signature
	hash    Hash
}

type Signature struct {
	PubKey  crypto.PublicKey
	SigData []byte
}

func main() {
	testMsgPack()
}



func testMsgPack(){
	t := &Transaction{
		Version:1,
		TxType:1,
		Nonce:0,
		Payload:[]byte("test"),
		Sig: struct {
			PubKey  crypto.PublicKey
			SigData []byte
		}{PubKey: "01aa76587cdcbf7e8a8ca66ed27fbba5ffd9a15f", SigData: []byte("12345678")},
		hash:Hash{},
	}

	b, err := msgpack.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	var tx Transaction
   	err = msgpack.Unmarshal(b, &tx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tx)
}

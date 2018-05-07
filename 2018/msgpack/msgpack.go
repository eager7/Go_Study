package main

import (
	"crypto"
	"github.com/vmihailenco/msgpack"
	"fmt"
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
		}{PubKey: nil, SigData: []byte("")},
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

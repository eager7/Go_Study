package eip55

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"strings"
)

func Hex(addr string) string {
	if len(addr) > 1 {
		if addr[0:2] == "0x" || addr[0:2] == "0X" {
			addr = addr[2:]
		}
	}
	if len(addr)%2 == 1 {
		addr = "0" + addr
	}
	unCheckSummed := strings.ToLower(addr)
	fmt.Println("unCheckSummed:", unCheckSummed)
	sha := sha3.NewLegacyKeccak256()
	sha.Write([]byte(unCheckSummed))
	hash := sha.Sum(nil)

	result := []byte(unCheckSummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32 //转换为大写
		}
	}
	return "0x" + string(result)
}

func Address() string {
	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	pub := privateKeyECDSA.PublicKey
	pubData := elliptic.Marshal(crypto.S256(), pub.X, pub.Y)
	sha := sha3.NewLegacyKeccak256()
	sha.Write([]byte(pubData[1:]))
	pubHash := sha.Sum(nil)[12:]

	return "0x" + hex.EncodeToString(pubHash)
}
//
//func AddressFromPriv(priv string) string {
//	ecdsa.PrivateKey{}.
//}
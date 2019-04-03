package main

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func main() {
	fmt.Println("transfer:", EIP165Method("transfer(address,uint256)"))
	fmt.Println("event:", EIP165Event(`Transfer(address,address,uint256)`))

}

func EIP165Method(method string) string {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(method))
	return hex.EncodeToString(hash.Sum(nil)[0:4])
}

func EIP165Event(event string) string {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(event))
	return hex.EncodeToString(hash.Sum(nil))
}
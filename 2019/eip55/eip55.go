package eip55

import (
	"fmt"
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

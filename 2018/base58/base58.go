package main

import (
	"github.com/btcsuite/btcutil/base58"
	"fmt"
	"math/big"
	"regexp"
)

func main() {
	testRegexp()
}

func testBase58() {
	name := "pct"
	base := base58.Encode([]byte(name))
	uuid := new(big.Int).SetBytes([]byte(base)).Uint64()
	fmt.Println(uuid)

	uuid2 := new(big.Int).SetUint64(uuid)
	base2 := base58.Decode(string(uuid2.Bytes()))
	fmt.Println(string(base2))
}

func testRegexp() {
	reg := `^[1-5a-z]{1,12}$`
	rgx := regexp.MustCompile(reg)
	s := []string{"pct", "Pct", "panchangtaoaa"}
	for _, v := range s {
		fmt.Println(rgx.MatchString(v))
	}
}
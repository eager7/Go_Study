package eip55

import (
	"testing"
)

func TestEthAddress(t *testing.T) {
	Check("0xaAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa", t)
	Check("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed", t)
	Check("0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359", t)
	Check("0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB", t)
	Check("0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb", t)
}

func Check(address string, t *testing.T) {
	if addr := Hex(address); addr != address {
		t.Fatal("unequal:", addr)
	}
}

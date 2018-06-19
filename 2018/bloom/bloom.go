package main

import (
	"fmt"
	"github.com/irfansharif/cfilter"
	"github.com/willf/bloom"
	"github.com/willf/bitset"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/util"
	"bytes"
	"math/big"
	"gitlab.quachain.net/aba/aba/common"
)

func main() {
	myBloom()
	//Cfilter()
}

func xBloom() {
	fmt.Println("Test Bloom")
	fmt.Println(bitset.Cap())
	b := bloom.New(10000000, 50)
	b.Add([]byte("pct"))
	fmt.Println(b.Test([]byte("pct")))
	fmt.Println(b.Test([]byte("pc2t")))

	var bf bytes.Buffer
	if _, err := b.WriteTo(&bf); err != nil {
		fmt.Println(err)
		return
	}
	by := bf.Bytes()
	fmt.Println("len of by:", len(by))

	b2 := bloom.New(1, 1)
	b2.ReadFrom(bytes.NewBuffer(by))
	fmt.Println(b2.Test([]byte("pct")))
	fmt.Println(b2.Test([]byte("pc2t")))

}

func Cfilter() {
	cfilter.New()
	fmt.Println("Test Cfilter")
}

func LevelDBBloom(){
	f := filter.NewBloomFilter(10)
	f.NewGenerator().Add([]byte("pct"))
	b := &util.Buffer{}
	f.NewGenerator().Generate(b)
	fil := b.Bytes()
	fmt.Println(f.Contains(fil, []byte("pct")))
}





////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
const (
	// BloomByteLength represents the number of bytes used in a header log bloom.
	BloomByteLength = 256

	// BloomBitLength represents the number of bits used in a header log bloom.
	BloomBitLength = 8 * BloomByteLength
)

// Bloom represents a 2048 bit bloom filter.
type Bloom [BloomByteLength]byte

// SetBytes sets the content of b to the given bytes.
// It panics if d is not of suitable size.
func (b *Bloom) SetBytes(d []byte) {
	if len(b) < len(d) {
		panic(fmt.Sprintf("bloom bytes too big %d %d", len(b), len(d)))
	}
	copy(b[BloomByteLength-len(d):], d)
}

// BytesToBloom converts a byte slice to a bloom filter.
// It panics if b is not of suitable size.
func BytesToBloom(b []byte) Bloom {
	var bloom Bloom
	bloom.SetBytes(b)
	return bloom
}

func (b *Bloom) AddKey(data []byte) {
	b.Add(new(big.Int).SetBytes([]byte(data)))
}

// Add adds d to the filter. Future calls of Test(d) will return true.
func (b *Bloom) Add(d *big.Int) {
	bin := new(big.Int).SetBytes(b[:])
	bin.Or(bin, bloom9(d.Bytes()))
	b.SetBytes(bin.Bytes())
}

// Big converts b to a big integer.
func (b Bloom) Big() *big.Int {
	return new(big.Int).SetBytes(b[:])
}

func (b Bloom) Bytes() []byte {
	return b[:]
}

func bloom9(b []byte) *big.Int {
	b = common.Keccak256Hash(b[:]).Bytes()

	r := new(big.Int)

	for i := 0; i < 6; i += 2 {
		t := big.NewInt(1)
		b := (uint(b[i+1]) + (uint(b[i]) << 8)) & 2047
		r.Or(r, t.Lsh(t, b))
	}

	return r
}
func (b Bloom) Test(test *big.Int) bool {
	return BloomLookup(b, test.Bytes())
}

func (b Bloom) TestBytes(test []byte) bool {
	return b.Test(new(big.Int).SetBytes(test))

}
func BloomLookup(bin Bloom, key []byte) bool {
	bloom := bin.Big()
	cmp := bloom9(key)

	return bloom.And(bloom, cmp).Cmp(cmp) == 0
}

func myBloom() {
	b := BytesToBloom(nil)
	b.AddKey([]byte("pct"))
	fmt.Println(b.TestBytes([]byte("pct")))
	fmt.Println(b.TestBytes([]byte("pct2")))
}
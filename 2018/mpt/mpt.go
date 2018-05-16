package main

import (
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb"
)

func main() {
	testTrie()
}

func testTrie() {
	diskDb, _ := ethdb.NewMemDatabase()
	tree, err := trie.New(common.Hash{}, trie.NewDatabase(diskDb))
	if err != nil {
		fmt.Println(err)
	}
	//tree.Update([]byte("pct"), []byte("panchangtao"))
	tree.Update([]byte("doe"), []byte("reindeer"))
	tree.Update([]byte("dog"), []byte("puppy"))
	tree.Update([]byte("dogglesworth"), []byte("cat"))
	fmt.Println("root:", tree.Hash().String())

	value := tree.Get([]byte("dog"))
	fmt.Println("dog value:", string(value))

	fmt.Println("delete dog")
	tree.Delete([]byte("dog"))
	tree.Commit(nil)

	value = tree.Get([]byte("dog"))
	fmt.Println("dog value:", string(value))




}

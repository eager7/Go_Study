package main

import (
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/core/state"
)

func main() {
	//testTrie()
	testDiskDB()
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

func testDiskDB() {
	diskDb, _ := ethdb.NewLDBDatabase("Test", 0, 0)
	iterate := diskDb.NewIterator()
	for iterate.Next() {
		fmt.Println(string(iterate.Key()), string(iterate.Value()))
	}
	iterate.Release()
	if err := iterate.Error(); err != nil {
		fmt.Println(err)
		return
	}

	tree, err := trie.New(common.Hash{}, trie.NewDatabase(diskDb))
	if err != nil {
		fmt.Println(err)
	}

	value := tree.Get([]byte("dog"))
	fmt.Println("dog value:", string(value))

	tree.Update([]byte("doe"), []byte("reindeer"))
	tree.Update([]byte("dog"), []byte("puppy"))
	tree.Update([]byte("dogglesworth"), []byte("cat"))
	fmt.Println("root:", tree.Hash().String())

	value = tree.Get([]byte("dog"))
	fmt.Println("dog value:", string(value))

	tree.Commit(nil)

}

func test3() {
	diskDb, _ := ethdb.NewLDBDatabase("Test", 0, 0)
	Db := state.NewDatabase(diskDb)
	Db.OpenTrie()

}

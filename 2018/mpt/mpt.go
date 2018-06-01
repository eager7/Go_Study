package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
	aba_common "gitlab.quachain.net/aba/aba/common"
	"gitlab.quachain.net/aba/aba/core/store"
	aba_state "gitlab.quachain.net/aba/aba/core/state"
)

func main() {
	//testTrie()
	//testDiskDB()
	test4()
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
	fmt.Println(tree.Hash().String())
}

func test3() {
	diskDb, _ := ethdb.NewLDBDatabase("Test", 0, 0)
	Db := state.NewDatabase(diskDb)

	root := common.HexToHash("0xd4cd937e4a4368d7931a9cf51686b7e10abb3dce38a39000fd7902a092b64585")
	fmt.Println(root.String())
	t, err := Db.OpenTrie(root)
	if err != nil {
		fmt.Println(err)
		t, _ = Db.OpenTrie(common.Hash{})
	}
	fmt.Println("Root:", t.Hash().String())
	value, err := t.TryGet([]byte("dog"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("dog value:", string(value))

	t.TryUpdate([]byte("doe"), []byte("reindeer"))
	t.TryUpdate([]byte("dog"), []byte("puppy"))
	t.TryUpdate([]byte("dogglesworth"), []byte("cat"))
	fmt.Println("root:", t.Hash().String())
	t.Commit(nil)

	trie := Db.TrieDB()
	trie.Commit(t.Hash(), false)

}

func test4() {
	diskDb, _ := store.NewLevelDBStore("Test", 0, 0)
	Db := aba_state.NewDatabase(diskDb)

	root := aba_common.HexToHash("0xd4cd937e4a4368d7931a9cf51686b7e10abb3dce38a39000fd7902a092b64585")
	fmt.Println(root.HexString())
	t, err := Db.OpenTrie(root)
	if err != nil {
		fmt.Println(err)
		t, _ = Db.OpenTrie(aba_common.Hash{})
	}
	fmt.Println("Root0:", t.Hash().HexString())
	value, err := t.TryGet([]byte("dog"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("dog value:", string(value))

	t.TryUpdate([]byte("doe"), []byte("reindeer"))
	fmt.Println("root1:", t.Hash().HexString())
	t.TryUpdate([]byte("dog"), []byte("puppy"))
	fmt.Println("root2:", t.Hash().HexString())
	t.TryUpdate([]byte("dogglesworth"), []byte("cat"))
	fmt.Println("root3:", t.Hash().HexString())
	t.TryUpdate([]byte("dogglesworth"), []byte("cat"))
	fmt.Println("root4:", t.Hash().HexString())
	t.Commit(nil)
	fmt.Println("root5:", t.Hash().HexString())

	trie := Db.TrieDB()
	trie.Commit(t.Hash(), false)
	fmt.Println("root6:", t.Hash().HexString())
}
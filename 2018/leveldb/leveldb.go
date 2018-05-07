package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"fmt"
	"bytes"
	"encoding/binary"
	"runtime"
	"time"
)

func main() {
	fmt.Println("level db test")
	testLevelDB()
}

func testLevelDB(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	var start,end int64

	o := opt.Options{
		NoSync: false,
		Filter: filter.NewBloomFilter(10),
	}
	db, err := leveldb.OpenFile("test", &o)
	if err != nil {
		fmt.Println(err)
		return
	}

	bath := new(leveldb.Batch)
	bath.Put([]byte("pct"), []byte("panchangtao"))

	if err := db.Write(bath, nil); err != nil {
		fmt.Println("write err:", err)
	}

	var v []byte
	if v, err = db.Get([]byte("pct"), nil); err != nil {
		fmt.Println("read error:", err)
	}
	fmt.Printf("get pct value:%s\n", v)

	start = time.Now().UnixNano()
	var i int32
	for i = 1; i < 10000000; i++ {
		bufByte := bytes.NewBuffer([]byte{})
		binary.Write(bufByte, binary.BigEndian, i)
		//fmt.Println(i, bufByte.Bytes(), len(bufByte.Bytes()))
		db.Put(bufByte.Bytes(), bufByte.Bytes(), nil)
	}
	end = time.Now().UnixNano()
	fmt.Println("Put time:", (end - start)/1000000)

	start = time.Now().UnixNano()
	k := []byte{0x00,0x00,0x03,0xe7}
	v,_ = db.Get(k, nil)
	fmt.Println(v)
	end = time.Now().UnixNano()
	fmt.Println("Batch Put time:", (end - start)/1000000)



	db.Close()
}
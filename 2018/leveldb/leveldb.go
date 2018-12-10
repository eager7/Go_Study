package main

import (
	"github.com/ecoball/go-ecoball/common/elog"
	"github.com/ecoball/go-ecoball/core/store"

	"os"
	"log"
	"runtime/pprof"
	"runtime"
	"fmt"
	"runtime/trace"
	"flag"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	{//TRACE
		ft, err := os.Create("ProfileTrace")
		if err != nil {
			log.Fatal("could not create trace profile: ", err)
		}
		defer ft.Close()

		trace.Start(ft)
		defer trace.Stop()
	}

	os.RemoveAll("/tmp/store_benchmark")
	s, _ := store.NewBlockStore("/tmp/store_benchmark")

	for i := 0; i < 1000; i++ {
		s.Put([]byte(fmt.Sprintf("key:%d", i)), []byte("value"))
		v, _ := s.Get([]byte(fmt.Sprintf("key:%d", i)))
		elog.Log.Debug(string(v))
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}

func F() {
	os.RemoveAll("/tmp/store_benchmark")
	s, _ := store.NewBlockStore("/tmp/store_benchmark")

	for i := 0; i < 1000; i++ {
		s.Put([]byte(fmt.Sprintf("key:%d", i)), []byte("value"))
		v, _ := s.Get([]byte(fmt.Sprintf("key:%d", i)))
		elog.Log.Debug(string(v))
	}
}
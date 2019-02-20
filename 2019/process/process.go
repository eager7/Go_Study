package process

import (
	"fmt"
	"github.com/jbenet/goprocess"
	"github.com/jbenet/goprocess/periodic"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Initialize() {
	periodic := func(worker goprocess.Process) {
		fmt.Println("process running...")
	}
	process := periodicproc.Tick(time.Second*1, periodic)
	process.Go(periodic)

	Pause()
}

func Pause() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(interrupt)
	sig := <-interrupt
	fmt.Println(" program received exit signal:", sig)
}
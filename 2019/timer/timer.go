package timer

import (
	"fmt"
	"time"
)

func TimerTick() {
	ticker1 := time.NewTicker(time.Second)
	ticker2 := time.NewTicker(time.Minute	)
	for {
		select {
		case <-ticker1.C:
			fmt.Println("ticker 1")
		case <-ticker2.C:
			fmt.Println("ticker 2")
		}
	}
}

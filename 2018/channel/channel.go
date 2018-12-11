package channel

import (
	"context"
	"fmt"
	"time"
)

func Select() {
	ctx := context.Background()
	outbox := make(chan (<-chan string), 0)
	go routine(ctx, outbox)
	fmt.Println(<-outbox)
	fmt.Println(<-outbox)
}

func routine(ctx context.Context, c chan (<-chan string)) {
	defer close(c)
	for ; ;  {
		oneTimeUse := make(chan string, 1)
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		case c <- oneTimeUse:
			fmt.Println("outbox")
		}
		close(oneTimeUse)
		time.Sleep(time.Millisecond*100)
	}
}
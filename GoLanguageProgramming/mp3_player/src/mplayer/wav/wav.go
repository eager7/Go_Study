package wav
import (
	"fmt"
	"time"
)

type WAVPlayer struct{
	stat int
	progress int
}

func (p *WAVPlayer)Play(source string){
	fmt.Println("wav player, source", source)

	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	println("\nFinished Playing", source)
}
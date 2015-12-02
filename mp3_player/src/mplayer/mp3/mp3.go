package mp3
import (
	"fmt"
	"time"
)

type MP3Player struct{
	stat int
	progress int
}

func (p *MP3Player)Play(source string){
	fmt.Println("mp3 player, source", source)

	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	println("\nFinished Playing", source)
}
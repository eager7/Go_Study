package mplayer
import "fmt"
import "mplayer/mp3"
import "mplayer/wav"

type Player interface{
	Play(source string)
}

func Play(source, mtype string){
	var p Player

	switch mtype{
		case "MP3":
			fmt.Println("MP3 format")
			p = &mp3.MP3Player{}
		case "WAV":
			fmt.Println("WAV format")
			p = &wav.WAVPlayer{}
		default:
			fmt.Println("Unknow format", mtype,"please input MP3 or WAV")
			return
	}
	p.Play(source)
}


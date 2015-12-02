package mp3
import "testing"

func TestOps(t *testing.T){
	println("mp3 player testing")

	mp := &MP3Player{}
	mp.Play("www.baidu.com")
}
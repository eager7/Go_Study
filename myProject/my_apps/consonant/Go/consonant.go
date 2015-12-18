package main
import "fmt"
import "bufio"
import "os"

func main() {
	fmt.Println("Consonant Go Test")
	Inp := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your str...")
	str, err := Inp.ReadString('\n');
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("your input is:", str)
	str = ConsonantString(str)
	fmt.Println("consonant str is:", str)
}

func ConsonantString(str string) string {
	conChar := []byte {0,'b',0,'d',0,'f','g','h','i','j','k',0,'m','n',0,'p','q','r','s','t',0,'v',0,0,0,'z'}
	headChar := []byte(str)
	if headChar[0] == conChar[headChar[0] - 0x61] {
		fmt.Println("your input is a consonant")
		str = str[1:len(str)-1] + "-" + string(headChar[0]) + "ay"
		return str
	}
	return str
}
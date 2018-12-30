package main
import "fmt"
import (
	"os"
	"bufio"
	"regexp"
)

func main(){
	fmt.Println("dict test program")
	buffer, err := ReadFile("dictionary.js")
	if err != nil {
		fmt.Println("read file error:", err)
	}
	fmt.Println(len(buffer))
	ParseString(
		`$a:"n. A As 或 A''s  安 ampere  a  art.一 n.字母A  [军] Analog.Digital 模拟 数字   =account of  帐上",
$aaal:"American Academy of Arts and Letters 美国艺术和文学学会",
$aachen:" 亚琛[德意志联邦共和国西部城市]",
`)

}

func ParseString(str string) {
	reg := regexp.MustCompile(`(^.*,)`)
	lists := reg.FindAllString(str, -1)
	fmt.Println("word number:", len(lists), lists)
	fmt.Println("first word:", lists[0])
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fmt.Println("the file size is:", stat.Size())
	buffer := make([]byte, stat.Size())
	reader := bufio.NewReader(file)
	if n, err := reader.Read(buffer); err != nil {
		return nil, err
	} else {
		fmt.Println("read file size:", n)
	}
	return buffer, nil
}

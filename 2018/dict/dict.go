package main
import "fmt"
import (
	"os"
	"bufio"
	"regexp"
)

const example = `$a:"n. A As 或 A''s  安 ampere  a  art.一 n.字母A  [军] Analog.Digital 模拟 数字   =account of  帐上",
$aaal:"American Academy of Arts and Letters 美国艺术和文学学会",
$aachen:" 亚琛[德意志联邦共和国西部城市]",`

var plainchant = "chant"

func main(){
	fmt.Println("dict test program")
	buffer, err := ReadFile("dictionary.js")
	if err != nil {
		fmt.Println("read file error:", err)
	}
	fmt.Println(len(buffer))
	ParseString(string(buffer))
}

func ParseString(str string) {
	reg := regexp.MustCompile(`(?m)(^.*,)`)//(?m)多行模式,(^.*)分组匹配多个
	lists := reg.FindAllString(str, -1)
	for k, v := range lists {
		reg := regexp.MustCompile(`.[a-z]+`)//跳过开头的$符号，查找连续的小写字母
		lists := reg.FindAllString(v, 1)
		if len(lists) != 0 && len(lists[0]) == 12 { //只查找到一个长度为12的字符时，这个就是单词
			//fmt.Println(k, v)
		}
		reg = regexp.MustCompile(plainchant)
		lists = reg.FindAllString(v, 1)
		if len(lists) != 0 {
			fmt.Println(k, v)
		}
	}
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

package main
import (
	"encoding/json"
	"os"
	"fmt"
	"log"
)

func main(){
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			fmt.Println(err)
			log.Println(err)
			return
		}
		fmt.Println(v)
		/*for k := range v {
			if k != "Title" {
				v[k] = nil, false
			}
		}*/
		if err := enc.Encode(&v,); err != nil {
			log.Println(err)
		}
	}
}
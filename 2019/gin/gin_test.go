package gin

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestInitializeGin(t *testing.T) {
	InitializeGin()
}

type Resp struct {
	Value     float64
	Text      string
	Timestamp int64
}

var url = "https://forex.1forge.com/1.0.3/convert?from=USD&to=CNH&quantity=1&api_key=irSQh8AuW0x0Cu6JSbd8VkTGLBOoCdkn"

func TestSendRequest(t *testing.T) {
	/*if body, err := SendRequest("GET", nil, url); err != nil {
		t.Fatal(err)
	} else {
		t.Log(string(body))
	}*/

	resp := Resp{}
	if err := SendAndRecv(url, &resp); err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

var text = `{"value":6.7329,"text":"1 USD is worth 6.7329 CNH","timestamp":1550627859}`

func TestJson(t *testing.T) {
	resp := &Resp{}
	fmt.Println(reflect.ValueOf(resp))
	if err := json.Unmarshal([]byte(text), &resp); err != nil {
		t.Fatal("Unmarshal err:", err)
	}
	t.Log(resp)
}
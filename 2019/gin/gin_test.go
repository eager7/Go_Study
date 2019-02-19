package gin

import "testing"

func TestInitializeGin(t *testing.T) {
	InitializeGin()
}

type Resp struct {
	value     float64
	text      string
	timestamp int64
}

func TestSendRequest(t *testing.T) {
	resp := Resp{}
	if err := SendAndRecv("https://forex.1forge.com/1.0.3/convert?from=USD&to=CNH&quantity=1&api_key=irSQh8AuW0x0Cu6JSbd8VkTGLBOoCdkn", &resp); err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

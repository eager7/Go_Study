package kafka

import "testing"

func TestInitialize(t *testing.T) {
	addresses := []string{`192.168.0.1:9092`}
	topic := []string{`test1`, `test2`}
	group := `test`
	Initialize(addresses, group, topic)
}

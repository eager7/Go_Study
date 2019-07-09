package kafka

import "testing"

func TestInitialize(t *testing.T) {
	addresses := []string{`47.52.147.232:9092`, `47.244.47.241:9092`, `47.75.179.251:9092`}
	topic := []string{`eos_blockchain_transaction`, `eos_blockchain_block`}
	group := `eos_blockchain`
	Initialize(addresses, group, topic)
}

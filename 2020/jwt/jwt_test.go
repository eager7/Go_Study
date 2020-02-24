package jwt

import "testing"

func TestInit(t *testing.T) {
	tok := NewToken("pct")
	//tok := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODIzODQ3NDZ9.lH2vy4AkjnzZhJVz1gSkrdhNzxNsq2jIMUCOVCjDN5Y"
	if err := Verify("pct", tok); err != nil {
		t.Fatal(err)
	}
}

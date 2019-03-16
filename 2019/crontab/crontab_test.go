package crontab

import (
	"testing"
)

func TestInitialize(t *testing.T) {
	if err := Initialize(); err != nil {
		t.Fatal(err)
	}
	make(chan struct{})<- struct{}{}

}

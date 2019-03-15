package crontab

import (
	"testing"
	"time"
)

func TestInitialize(t *testing.T) {
	if err := Initialize(); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second*100)
}

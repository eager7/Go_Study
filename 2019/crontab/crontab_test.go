package crontab

import (
	"fmt"
	"testing"
	"time"
)

func TestInitialize(t *testing.T) {
	fmt.Println(time.Now().Local())
	if err := Initialize(); err != nil {
		t.Fatal(err)
	}
	make(chan struct{})<- struct{}{}

}

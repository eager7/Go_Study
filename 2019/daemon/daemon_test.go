package daemon_test

import (
	"fmt"
	"github.com/eager7/go_study/2019/daemon"
	"testing"
)

func TestInitialize(t *testing.T) {
	c, err := daemon.Initialize()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("run daemon")
	if err := c.Release(); err != nil {
		t.Fatal(err)
	}
}

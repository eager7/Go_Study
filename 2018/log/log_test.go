package log_test

import (
	"testing"
	"github.com/eager7/go_study/2018/log"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestLog(t *testing.T) {
	l := log.Logger()
	l.Debug("Debug")
}
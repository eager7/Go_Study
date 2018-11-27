package logger_test

import (
	"testing"
	"github.com/eager7/go_study/log/logger"
)

func TestLogger(t *testing.T) {
	l := logger.NewLogger("", 0)
	for i := 0; i < 10; i ++ {
		l.Debug("debug------------------")
		l.Info("info----------------------")
		l.Warn("warn------------------")
		l.Error("error---------------------")
	}

}

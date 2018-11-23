package logger_test

import (
	"testing"
	"github.com/eager7/go_study/log/logger"
	"github.com/eager7/go_study/log/logbunny"
)

func TestLogger(t *testing.T) {
	l := logger.NewLogger("", 0)
	for i := 0; i < 10; i ++ {
		l.Debug("debug------------------")
		l.Info("info----------------------")
		l.Warn("warn------------------")
		l.Error("error---------------------")
		ll := l.GetLogger()
		ll.Debug("debug", logbunny.Float64("pct", 2.3456678))
		ll.Info("info", logbunny.Float64("pct", 2.3456678))
		ll.Warn("warn")
		ll.Error("error")
	}

}

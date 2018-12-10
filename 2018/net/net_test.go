package net_test

import (
	"testing"
	"github.com/eager7/go/mlog"
	"github.com/eager7/go_study/2018/net"
	"context"
)

func TestNet(t *testing.T) {
	mlog.L.Debug("net test program...")

	ctx, cancel := context.WithCancel(context.Background())
	net.InitNetWork(ctx)
	net.StartNetWork(nil)

	cancel()
}

package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	suger := logger.Sugar()
	suger.Debug("debug")
}

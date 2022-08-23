package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	path := "zap.log"
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{path, "stdout"}
	logger, _ := cfg.Build()
	defer logger.Sync()

	url := "http://google.com"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))
}

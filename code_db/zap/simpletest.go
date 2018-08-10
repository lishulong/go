package main

import (
	"go.uber.org/zap"
	"time"
)

func TestSugar() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	sugar.Infof("failed to fetch URL: %s", "http://example.com")
}

func TestExampleLogger() {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func TestSugarAndDesugar() {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("", zap.Any("type", logger))

	sugar := logger.Sugar()
	sugar.Debugf("logger: %#v", logger)
	sugar.Debugf("logger: %#v", sugar)

	plain := sugar.Desugar()
	plain.Debug("hello")
}

func main() {
	//TestSugar()

	//TestExampleLogger()

	TestSugarAndDesugar()
}

package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// zap
	// zap 是 Go 语言中一个高性能的日志库，支持 结构化日志、高吞吐量 和 低延迟，适用于高并发环境。
	// 安装： go get -u go.uber.org/zap

	logger, _ := zap.NewProduction()
	defer logger.Sync() // 确保日志写入

	logger.Info(
		"hello zap",
		// 指定类型可以触发反射机制，性能高
		zap.String("level", "debug"),
		zap.Int("port", 8080),
	)
	logger.Debug("This is a debug message")
	logger.Warn("This is a warning")
	logger.Error("This is an error message")

	sugar := logger.Sugar()
	sugar.Info("This is a info message")
	sugar.Warn("This is a warning message")
	sugar.Infow("This is a info message")
	sugar.Errorw("This is a error message")
	sugar.Infof("This is a info message  %s %s", "key", "value")
	sugar.Debugf("This is a debug message  %s %s", "key", "value")

	// 自定义
	config := zap.Config{
		Encoding:         "console", // "json" 或 "console"
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"stdout", "app.log"}, // 输出到控制台和文件
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:      "time",
			LevelKey:     "level",
			CallerKey:    "caller",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.CapitalColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger2, _ := config.Build()
	logger2.Info("Custom log format")

}

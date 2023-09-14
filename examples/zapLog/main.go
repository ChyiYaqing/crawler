package main

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	logger, _ := zap.NewProduction()
	// 在进程退出之前落盘所有缓冲区的日志条目
	defer logger.Sync()

	url := "www.google.com"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))

	// 修改Zap打印时间的格式
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	// 日志切割
	w := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    500, // 日志的最大大小，以M为单位
		MaxBackups: 3,   // 保留旧日志文件的最大数量
		MaxAge:     28,  // 保留旧日志文件的最大天数
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger2 := zap.New(core)
	logger2.Info("this is test")
}

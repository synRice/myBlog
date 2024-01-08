package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:         "json",
		OutputPaths:      []string{"./log/app.log"}, // 输出到 log 文件夹下
		ErrorOutputPaths: []string{"./log/error.log"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}

func Info(ctx context.Context, msg string) {
	Logger.Info(msg)
}

func Error(ctx context.Context, err error) {
	Logger.Error(err.Error(), zap.Any("user", ctx.Value("user")))
}

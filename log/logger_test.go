package log

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:         "json",
		OutputPaths:      []string{"stdout", "./app.log"}, // 输出到 log 文件夹下
		ErrorOutputPaths: []string{"stderr", "./error.log"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}

func TestInfo(t *testing.T) {
	Info(context.Background(), "abc")
}

func TestError(t *testing.T) {
	Error(context.Background(), fmt.Errorf("%s", "lkd"))
}

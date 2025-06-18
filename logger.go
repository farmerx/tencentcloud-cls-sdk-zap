package tencentcloud_cls_sdk_zap

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a thin wrapper for zap.Logger that adds Ctx method.
type Logger struct {
	*zap.Logger
}

const contextKey = "context"

func (l *Logger) Sugar() *SugaredLogger {
	return &SugaredLogger{
		SugaredLogger: l.Logger.Sugar(),
	}
}

func (l *Logger) Ctx(ctx context.Context) *Logger {
	return l
}

func (l *Logger) With(fields ...zapcore.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
	}
}

type SugaredLogger struct {
	*zap.SugaredLogger
}

func (l *SugaredLogger) Ctx(ctx context.Context) *SugaredLogger {
	return &SugaredLogger{
		SugaredLogger: l.SugaredLogger,
	}
}

package ulog

import (
	"context"
	"io"
)

type nilWriter struct{}

func (nilWriter) Write(p []byte) (n int, err error) { return len(p), nil }

var (
	_nilWriter    io.Writer = nilWriter{}
	DefaultLogger           = New()
)

func Debug(ctx context.Context, msg string, data ...any) {
	DefaultLogger.Debug(ctx, msg, data...)
}
func Info(ctx context.Context, msg string, data ...any) {
	DefaultLogger.Info(ctx, msg, data...)
}

func Warn(ctx context.Context, msg string, data ...any) {
	DefaultLogger.Warn(ctx, msg, data...)
}

func Error(ctx context.Context, msg string, data ...any) {
	DefaultLogger.Error(ctx, msg, data...)
}

func Panic(ctx context.Context, msg string, data ...any) {
	DefaultLogger.Panic(ctx, msg, data...)
}

func Fatal(ctx context.Context, msg string, data ...any) {
	DefaultLogger.Fatal(ctx, msg, data...)
}

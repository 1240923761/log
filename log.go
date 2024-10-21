package log

import (
	"context"
	"github.com/1240923761/log/formatter"
	"github.com/1240923761/log/hook"
	"github.com/1240923761/log/ulog"
	"github.com/1240923761/log/util"
	"io"
)

func SetLogLevel(level util.LogLevel) {
	ulog.DefaultLogger.SetLogLevel(level)
}

func SetFormatter(formatter formatter.Formatter) {
	ulog.DefaultLogger.SetFormatter(formatter)
}

func SetWriter(writer io.Writer) {
	ulog.DefaultLogger.SetWriter(writer)
}

func AddHooks(hooks ...hook.Hook) {
	ulog.DefaultLogger.AddHooks(hooks...)
}

func Debug(ctx context.Context, msg string, data ...any) {
	ulog.Debug(ctx, msg, data...)
}

func Info(ctx context.Context, msg string, data ...any) {
	ulog.Info(ctx, msg, data...)
}

func Warn(ctx context.Context, msg string, data ...any) {
	ulog.Warn(ctx, msg, data...)
}

func Error(ctx context.Context, msg string, data ...any) {
	ulog.Error(ctx, msg, data...)
}

func Panic(ctx context.Context, msg string, data ...any) {
	ulog.Panic(ctx, msg, data...)
}

func Fatal(ctx context.Context, msg string, data ...any) {
	ulog.Fatal(ctx, msg, data...)
}

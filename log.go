package log

import (
	"context"
	"fmt"
	"github.com/1240923761/log/ulog"
	"github.com/google/uuid"
)

const (
	TraceKey = "X-Trace-Id"
)

func _mix(ctx context.Context, msg string) string {
	if ctx == nil {
		return fmt.Sprintf("%s | %s", uuid.Must(uuid.NewV7()).String(), msg)
	}

	traceId := ctx.Value(TraceKey)
	if traceId == nil {
		return fmt.Sprintf("%s | %s", uuid.Must(uuid.NewV7()).String(), msg)
	}

	return fmt.Sprintf("%s | %s", traceId, msg)
}
func SetWXAddress(addr string) {
	ulog.DefaultLogger.SetWXAddress(addr)
}
func SetLogLevel(level ulog.LogLevel) {
	ulog.DefaultLogger.SetLogLevel(level)
}
func Debug(ctx context.Context, msg string, data ...any) {
	ulog.Debug(_mix(ctx, msg), data...)
}

func Info(ctx context.Context, msg string, data ...any) {
	ulog.Info(_mix(ctx, msg), data...)
}

func Warn(ctx context.Context, msg string, data ...any) {
	ulog.Warn(_mix(ctx, msg), data...)
}

func Error(ctx context.Context, msg string, data ...any) {
	ulog.Error(_mix(ctx, msg), data...)
}

func Panic(ctx context.Context, msg string, data ...any) {
	ulog.Panic(_mix(ctx, msg), data...)
}

func Fatal(ctx context.Context, msg string, data ...any) {
	ulog.Fatal(_mix(ctx, msg), data...)
}
func WX(ctx context.Context, msg string, data ...any) {
	ulog.WX(_mix(ctx, msg), data...)
}

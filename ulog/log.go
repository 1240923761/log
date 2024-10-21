package ulog

import (
	"context"
	"fmt"
	"github.com/1240923761/log/entity"
	"github.com/1240923761/log/formatter"
	"github.com/1240923761/log/hook"
	"github.com/1240923761/log/util"
	"io"
	"sync"
)

type logger struct {
	sync.Mutex
	formatter                       formatter.Formatter
	writer                          io.Writer
	debug, info, warn, error, fatal io.Writer
	level                           util.LogLevel
	hooks                           []hook.Hook
}

func (l *logger) SetLogLevel(level util.LogLevel) {
	l.Lock()
	defer l.Unlock()

	if level > util.LogLevelDebug {
		l.debug = _nilWriter
	} else {
		l.debug = l.writer
	}

	if level > util.LogLevelInfo {
		l.info = _nilWriter
	} else {
		l.info = l.writer
	}

	if level > util.LogLevelWarn {
		l.warn = _nilWriter
	} else {
		l.warn = l.writer
	}

	if level > util.LogLevelError {
		l.error = _nilWriter
	} else {
		l.error = l.writer
	}

	l.level = level
}

func (l *logger) SetFormatter(formatter formatter.Formatter) {
	l.Lock()
	defer l.Unlock()

	l.formatter = formatter
}

func (l *logger) SetWriter(writer io.Writer) {
	l.Lock()
	defer l.Unlock()

	l.writer = writer
}

func (l *logger) AddHooks(hooks ...hook.Hook) {
	l.Lock()
	defer l.Unlock()

	for idx := range hooks {
		if hooks[idx] == nil {
			l.hooks = append(l.hooks, hooks[idx])
		}
	}
}

func (l *logger) Debug(ctx context.Context, msg string, data ...any) {
	e := entity.NewEntity(ctx)
	e.Debug(msg, data...)

	fmt.Fprintln(l.debug, l.formatter.Format(e))
}

func (l *logger) Info(ctx context.Context, msg string, data ...any) {
	e := entity.NewEntity(ctx)
	e.Info(msg, data...)
	fmt.Fprintln(l.info, l.formatter.Format(e))
}

func (l *logger) Warn(ctx context.Context, msg string, data ...any) {
	e := entity.NewEntity(ctx)
	e.Warn(msg, data...)
	fmt.Fprintln(l.warn, l.formatter.Format(e))
}

func (l *logger) Error(ctx context.Context, msg string, data ...any) {
	e := entity.NewEntity(ctx)
	e.Error(msg, data...)
	fmt.Fprintln(l.error, l.formatter.Format(e))
}

func (l *logger) Panic(ctx context.Context, msg string, data ...any) {
	e := entity.NewEntity(ctx)
	e.Panic(msg, data...)
	panic(l.formatter.Format(e))
}

func (l *logger) Fatal(ctx context.Context, msg string, data ...any) {
	e := entity.NewEntity(ctx)
	e.Fatal(msg, data...)
	fmt.Fprintln(l.writer, l.formatter.Format(e))
}

package formatter

import (
	"context"
	"fmt"
	"github.com/1240923761/log/entity"
	"github.com/1240923761/log/util"
	"github.com/google/uuid"
	"strings"
	"sync"
	"time"
)

type TextFormatter struct {
	Prefix        string
	TimeFn        func() string
	LevelFormatFn func(level util.LogLevel) string
	Delimiter     string
	TraceFn       func(ctx context.Context) string
}

var (
	textPool = &sync.Pool{
		New: func() interface{} {
			return make([]string, 0, 8)
		},
	}
	DefaultTextFormatter = &TextFormatter{
		Prefix: "FILEBEAT",
		TimeFn: func() string { return time.Now().Format("2006-01-02T15:04:05") },
		LevelFormatFn: func(level util.LogLevel) string {
			switch level {
			case util.LogLevelDebug:
				return "DEBUG"
			case util.LogLevelInfo:
				return "INFO "
			case util.LogLevelWarn:
				return "WARN "
			case util.LogLevelError:
				return "ERROR"
			case util.LogLevelFatal:
				return "FATAL"
			case util.LogLevelPanic:
				return "PANIC"
			default:
				return "EMPTY"
			}
		},
		Delimiter: "|",
		TraceFn: func(ctx context.Context) string {
			var trace string

			if ctx == nil {
				goto END
			}

			if trace, _ = ctx.Value(util.TraceKey).(string); trace != "" {
				return trace
			}
		END:
			return uuid.Must(uuid.NewV7()).String()
		},
	}
)

func (f *TextFormatter) Format(entity *entity.Entity) string {
	msgs := textPool.Get().([]string)
	defer func() {
		textPool.Put(msgs[:0])
		entity.Close()
	}()

	if f.Prefix != "" {
		msgs = append(msgs, f.Prefix)
	}

	if level := f.LevelFormatFn(entity.Level); level != "" {
		msgs = append(msgs, level)
	}

	if now := f.TimeFn(); now != "" {
		msgs = append(msgs, now)
	}

	if trace := f.TraceFn(entity.Ctx); trace != "" {
		msgs = append(msgs, trace)
	}

	msgs = append(msgs, fmt.Sprintf(entity.Msg, entity.Args...))
	for _, item := range entity.Fields {
		msgs = append(msgs, fmt.Sprintf("%s=%v", item.Key, item.Value))
	}

	return strings.Join(msgs, " "+f.Delimiter+" ")
}

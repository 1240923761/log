package ulog

import (
	"github.com/fatih/color"
	"io"
	"sync"
	"time"
)

type LogLevel uint32

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelPanic
	LogLevelFatal
	LogLevelWX
)

type logger struct {
	sync.Mutex
	timeFormat string
	writer     io.Writer
	level      LogLevel
	debug      func(prefix, timestamp, msg string, data ...any)
	info       func(prefix, timestamp, msg string, data ...any)
	warn       func(prefix, timestamp, msg string, data ...any)
	error      func(prefix, timestamp, msg string, data ...any)
	panic      func(prefix, timestamp, msg string, data ...any)
	fatal      func(prefix, timestamp, msg string, data ...any)
	wx         func(wxAddress, prefix, timestamp, msg string, data ...any)
	wxAddress  string
}

var (
	red    = color.New(color.FgRed)
	hired  = color.New(color.FgHiRed)
	green  = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
	white  = color.New(color.FgWhite)
	blue   = color.New(color.FgBlue)
)

func (l *logger) SetTimeFormat(format string) {
	l.Lock()
	defer l.Unlock()
	l.timeFormat = format
}
func (l *logger) SetWXAddress(addr string) {
	l.Lock()
	defer l.Unlock()
	l.wxAddress = addr
}
func (l *logger) SetLogLevel(level LogLevel) {
	l.Lock()
	defer l.Unlock()

	if level > LogLevelDebug {
		l.debug = nilLogger
	} else {
		l.debug = normalLogger
	}

	if level > LogLevelInfo {
		l.info = nilLogger
	} else {
		l.info = normalLogger
	}

	if level > LogLevelWarn {
		l.warn = nilLogger
	} else {
		l.warn = normalLogger
	}

	if level > LogLevelError {
		l.error = nilLogger
	} else {
		l.error = normalLogger
	}

	if level > LogLevelPanic {
		l.panic = nilLogger
	} else {
		l.panic = panicLogger
	}

	if level > LogLevelFatal {
		l.fatal = nilLogger
	} else {
		l.fatal = fatalLogger
	}
	if level > LogLevelWX {
		l.wx = wxNilLogger
	} else {
		l.wx = wxLogger
	}
}

func (l *logger) Debug(msg string, data ...any) {
	l.debug(white.Sprint("FILEBEATS [DEBUG] "), time.Now().Format(l.timeFormat), msg, data...)
}

func (l *logger) Info(msg string, data ...any) {
	l.info(green.Sprint("FILEBEATS [INFO] "), time.Now().Format(l.timeFormat), msg, data...)
}

func (l *logger) Warn(msg string, data ...any) {
	l.warn(yellow.Sprint("FILEBEATS [WARN] "), time.Now().Format(l.timeFormat), msg, data...)
}

func (l *logger) Error(msg string, data ...any) {
	l.error(red.Sprint("FILEBEATS [ERROR] "), time.Now().Format(l.timeFormat), msg, data...)
}

func (l *logger) Panic(msg string, data ...any) {
	l.panic(hired.Sprint("FILEBEATS [PANIC] "), time.Now().Format(l.timeFormat), msg, data...)
}

func (l *logger) Fatal(msg string, data ...any) {
	l.fatal(hired.Sprint("FILEBEATS [FATAL] "), time.Now().Format(l.timeFormat), msg, data...)
}
func (l *logger) WX(msg string, data ...any) {
	l.wx(l.wxAddress, blue.Sprint("FILEBEATS [WX] "), time.Now().Format(l.timeFormat), msg, data...)
}

type WroteLogger interface {
	Info(msg string, data ...any)
}

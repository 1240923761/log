package util

type LogLevel uint32

const (
	LogLevelUnknown = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelPanic
	LogLevelFatal
)

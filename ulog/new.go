package ulog

import (
	"os"
	"sync"
)

func New(wxAddress string) *logger {
	return &logger{
		Mutex:      sync.Mutex{},
		timeFormat: "2006-01-02T15:04:05",
		writer:     os.Stdout,
		level:      LogLevelInfo,
		debug:      nilLogger,
		info:       normalLogger,
		warn:       normalLogger,
		error:      normalLogger,
		panic:      panicLogger,
		fatal:      fatalLogger,
		wx:         wxLogger,
		wxAddress:  wxAddress,
	}
}

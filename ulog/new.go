package ulog

import (
	"github.com/1240923761/log/formatter"
	"github.com/1240923761/log/util"
	"os"
	"sync"
)

func New() *logger {
	l := &logger{
		Mutex:     sync.Mutex{},
		writer:    os.Stdout,
		level:     util.LogLevelInfo,
		formatter: formatter.DefaultTextFormatter,
	}

	l.debug, l.info, l.warn, l.error = l.writer, l.writer, l.writer, l.writer

	return l
}

package ulog

import (
	"fmt"
	"github.com/1240923761/log/entity"
	"io"
	"os"
)

type printType uint8

const (
	printTypeNormal printType = iota
	printTypePanic
	printTypeFatal
)

func (l *logger) fprint(w io.Writer, e *entity.Entity, pt printType) {
	for _, hook := range l.hooks {
		hook.Process(e)
	}

	switch pt {
	case printTypePanic:
		panic(l.formatter.Format(e))
	case printTypeFatal:
		fmt.Fprintln(w, l.formatter.Format(e))
		os.Exit(1)
	}

	fmt.Fprintln(w, l.formatter.Format(e))
}

package entity

import (
	"context"
	"github.com/1240923761/log/util"
	"sync"
)

type Entity struct {
	Ctx    context.Context
	Level  util.LogLevel
	Msg    string
	Args   []any
	Fields []*Field
}

type Field struct {
	Key   string
	Value any
}

var (
	nf = func() *Entity {
		return &Entity{
			Args:   make([]any, 0, 2),
			Fields: make([]*Field, 0),
		}
	}

	pool = &sync.Pool{
		New: func() interface{} {
			return nf()
		},
	}
)

func NewEntity(ctx context.Context) *Entity {
	var (
		ok bool
		e  *Entity
	)

	e, ok = pool.Get().(*Entity)
	if !ok {
		e = nf()
	}

	e.Ctx = ctx

	return e
}

func (e *Entity) Close() {
	e.Ctx = nil
	e.Fields = e.Fields[:0]
	e.Level = util.LogLevelUnknown
	e.Args = e.Args[:0]
	// todo
	e.Msg = string([]byte(e.Msg)[:0])
	pool.Put(e)
}

func (e *Entity) SetLevel(level util.LogLevel) {
	e.Level = level
}

func (e *Entity) SetMsg(msg string, args ...any) {
	e.Msg, e.Args = msg, append(e.Args, args...)
}

func (e *Entity) Debug(msg string, args ...interface{}) {
	e.SetLevel(util.LogLevelDebug)
	e.SetMsg(msg, args...)
}

func (e *Entity) Info(msg string, args ...interface{}) {
	e.SetLevel(util.LogLevelInfo)
	e.SetMsg(msg, args...)
}

func (e *Entity) Warn(msg string, args ...interface{}) {
	e.SetLevel(util.LogLevelWarn)
	e.SetMsg(msg, args...)
}

func (e *Entity) Error(msg string, args ...interface{}) {
	e.SetLevel(util.LogLevelError)
	e.SetMsg(msg, args...)
}

func (e *Entity) Fatal(msg string, args ...interface{}) {
	e.SetLevel(util.LogLevelFatal)
	e.SetMsg(msg, args...)
}

func (e *Entity) Panic(msg string, args ...interface{}) {
	e.SetLevel(util.LogLevelPanic)
	e.SetMsg(msg, args...)
}

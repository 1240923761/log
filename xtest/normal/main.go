package main

import (
	"github.com/1240923761/log"
	"github.com/1240923761/log/entity"
	"github.com/1240923761/log/util"
	"time"
)

type tth struct{}

func (t *tth) Process(v *entity.Entity) error {
	v.Fields = append(v.Fields, &entity.Field{Key: "time", Value: time.Now().Format("2006-01-02T15:04:05")})
	return nil
}

func main() {
	th := &tth{}
	log.AddHooks(th)
	log.Info(nil, "hello %s", "world")
	log.SetLogLevel(util.LogLevelFatal)

	log.Info(nil, "hidden")

	log.Panic(nil, "no no no save me!!!")

	log.Fatal(nil, "quit")
}

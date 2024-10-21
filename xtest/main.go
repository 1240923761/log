package main

import (
	"github.com/1240923761/log"
	"github.com/1240923761/log/util"
)

func main() {
	log.Info(nil, "hello %s", "world")
	log.SetLogLevel(util.LogLevelFatal)
	log.Info(nil, "hidden")
}

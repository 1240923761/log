package main

import "github.com/1240923761/log"

func main() {
	log.SetWXAddress("")
	log.WX(nil, "trace_id")
}

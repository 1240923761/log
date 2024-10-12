package main

import "gitcode.com/peachesone/log"

func main() {
	log.SetWXAddress("")
	log.WX(nil, "带traceid")
}

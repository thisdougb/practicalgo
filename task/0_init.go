package task

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// package function as a var so we can easily override it in unit tests
var pause = func(seconds int) {
	time.Sleep(time.Duration(time.Duration(seconds) * time.Second))
}

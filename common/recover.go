package common

import (
	"log"
	"runtime"
)

func PrintRecoverFromPanic(param ...interface{}) {
	if err := recover(); err != nil {
		panicStack(err)
	}
}

func panicStack(err interface{}) {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	log.Printf("panic %v \n%s\n", err, string(buf[:n]))
}

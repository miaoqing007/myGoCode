package common

import (
	"fmt"
	"log"
	"runtime"
)

func PrintRecoverFromPanic(param ...interface{}) {
	var PanicStack []interface{}
	if x := recover(); x != nil {
		PanicStack = append(PanicStack, fmt.Sprintf("PanicStack Recover:%s↓\n", x))
		i := 0
		funcName, file, line, ok := runtime.Caller(i)
		for ok {
			PanicStack = append(PanicStack, fmt.Sprintf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line))
			i++
			funcName, file, line, ok = runtime.Caller(i)
		}
	}
	log.Println(PanicStack)
}

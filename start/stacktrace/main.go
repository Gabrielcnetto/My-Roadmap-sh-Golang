package main

import (
	"log"
	"runtime"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			PrintStackTrace()
		}
	}()

	causePanic()
}

func causePanic() {
	panic("test panic")
}

func PrintStackTrace() {
	buf := make([]byte, 1024*1024*1024) // Allocate a big enough buffer.
	stackSize := runtime.Stack(buf, false)
	log.Printf("%s\n", buf[:stackSize])
}

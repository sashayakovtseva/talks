package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"syscall"
)

func main() {
	var ws sync.WaitGroup

	ws.Add(1)
	go func() {
		defer ws.Done()
		runtime.LockOSThread()
		syscall.Setresuid(1, 1, 1)
		runtime.UnlockOSThread()
	}()

	ws.Wait()

	for i := 0; i < 1000; i++ {
		ws.Add(1)
		go func() {
			fmt.Println(os.Geteuid(), os.Getuid())
			ws.Done()
		}()
	}

	ws.Wait()
}

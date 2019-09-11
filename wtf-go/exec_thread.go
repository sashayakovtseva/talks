package main

import (
	"log"
	"os"
	"runtime"
	"sync"
	"syscall"
)

func init() {
	log.Printf("Init TID=%d", syscall.Gettid())
}

func main() {
	log.Printf("Main TID=%d", syscall.Gettid())

	var ws sync.WaitGroup

	ws.Add(1)
	go func() {
		defer ws.Done()

		log.Println("Locking thread")
		runtime.LockOSThread()
		log.Printf("Locked TID=%d", syscall.Gettid())
		err := syscall.Setresuid(1000, 1000, 1000)
		if err != nil {
			log.Printf("Could not setresuid: %v", err)
		}
		log.Println("Unlocking thread")
		runtime.UnlockOSThread()
	}()

	ws.Wait()

	for i := 0; i < 1000; i++ {
		ws.Add(1)
		go func() {
			log.Printf("TID=%d, euid=%d,  uid=%d", syscall.Gettid(), os.Geteuid(), os.Getuid())
			ws.Done()
		}()
	}

	ws.Wait()
}

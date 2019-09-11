package main

import (
	"log"
	"runtime"
	"sync"
	"syscall"
)

func init() {
	log.Printf("Init TID=%d", syscall.Gettid())
}

var lockedTID int

func main() {
	log.Printf("Main TID=%d", syscall.Gettid())

	var ws sync.WaitGroup

	ws.Add(1)
	go func() {
		defer ws.Done()

		log.Println("Locking thread")
		runtime.LockOSThread()
		lockedTID = syscall.Gettid()
		log.Printf("Locked TID=%d", lockedTID)
		err := syscall.Setresuid(1000, 1000, 1000)
		if err != nil {
			log.Printf("Could not setresuid: %v", err)
		}
		// log.Println("Unlocking thread")
		// runtime.UnlockOSThread()
	}()

	ws.Wait()

	for i := 0; i < 100000000; i++ {
		ws.Add(1)
		go func() {
			myTID := syscall.Gettid()
			if myTID == lockedTID {
				panic("thread is reused!")
			}
			// log.Printf("TID=%d, euid=%d,  uid=%d", myTID, os.Geteuid(), os.Getuid())
			ws.Done()
		}()
	}

	ws.Wait()
}

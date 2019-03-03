package main

import (
	"log"
	"os"
	"runtime"
	"syscall"
)

func main() {
	f, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	runtime.SetFinalizer(f, nil)
	fd := int(f.Fd())
	runtime.GC()
	if err := syscall.Close(fd); err != nil {
		log.Println(err)
	}
	log.Println(fd)
}

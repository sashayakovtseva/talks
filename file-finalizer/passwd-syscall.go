package main

import (
	"log"
	"os"
	"runtime"
	"syscall"
)

func main() {
	fd, err := syscall.Open("/etc/passwd", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	runtime.GC()
	if err := syscall.Close(fd); err != nil {
		log.Println(err)
	}
	log.Println(fd)
}

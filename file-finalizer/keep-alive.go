package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
)

type File struct {
	d int
}

func main() {
	testFile := filepath.Join(os.TempDir(), "keep_alive")
	err := ioutil.WriteFile(testFile, []byte("hello, finalizer!"), 0755)
	if err != nil {
		log.Fatalf("Could not write test file: %v", err)
	}
	d, err := syscall.Open(testFile, syscall.O_RDONLY, 0)
	if err != nil {
		log.Fatalf("Could not open test file: %v", err)
	}
	p := &File{d}
	runtime.SetFinalizer(p, func(p *File) {
		log.Printf("Closed file with error: %v", syscall.Close(p.d))
	})
	var buf [10]byte
	// p can be freed here // HLerror
	n, err := syscall.Read(p.d, buf[:])
	runtime.KeepAlive(p)
	log.Printf("Read %d bytes with error %v", n, err)
}

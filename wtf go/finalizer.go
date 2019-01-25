package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
)

func main() {
	f, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	fd := int(f.Fd())
	log.Printf("File %s, fd = %d", f.Name(), fd)
	runtime.SetFinalizer(f, nil)
	ll("/proc/self/fd")
	log.Println("running gc")
	runtime.GC()
	ll("/proc/self/fd")
	log.Println("closing fd")
	if err := syscall.Close(fd); err != nil {
		log.Printf("Could not close fd: %v", err)
	}
}

func ll(dir string) {
	fii, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("Could not read %d: %v", dir, err)
		return
	}
	log.Printf("%s:", dir)
	for _, fi := range fii {
		link, _ := filepath.EvalSymlinks(filepath.Join(dir, fi.Name()))
		log.Printf("\t%s -> %s", fi.Name(), link)
	}
}

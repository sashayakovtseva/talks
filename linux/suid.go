package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func main() {
	printIDs()

	if len(os.Args) == 1 {
		return
	}

	uid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Could not parse uid to set: %v\n", err)
		os.Exit(1)
	}

	err = syscall.Setuid(uid)
	if err != nil {
		fmt.Printf("Setuid returned error: %v\n", err)
	}

	printIDs()
}

func printIDs() {
	fmt.Printf("uid=%v, euid=%v\n", os.Getuid(), os.Geteuid())
}

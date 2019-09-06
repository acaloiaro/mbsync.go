package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	help := "Please run with a valid golang duration: e.g. go run mbsync.go 5m"

	if len(os.Args) < 2 {
		fmt.Println(help)
		return
	}

	interval, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Println(help)
		return
	}

	ticker := time.Tick(interval)
	wait := time.Tick(1 * time.Minute)

	// Perform a sync before entering the loop
	sync()

	for {
		select {
		case <-ticker:
			sync()
		case <-wait:
			log.Println("Waiting...")
		}
	}
}

func sync() {
	log.Println("Syncing...")
	cmd := exec.Command("mbsync", "-a")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("ERR:", err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}

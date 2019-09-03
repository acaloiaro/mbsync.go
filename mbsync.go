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

	for {
		select {
		case <-ticker:
			log.Println("Syncing...")
			cmd := exec.Command("mbsync", "-a")
			stdoutStderr, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", stdoutStderr)
		case <-wait:
			log.Println("Waiting...")
		}
	}
}

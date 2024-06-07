package main

import (
	"fmt"
	"os"
)

func main() {
	// Option 1: Using the os package
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error retrieving hostname: %v\n", err)
	} else {
		fmt.Printf("Hostname using os package: %s\n", hostname)
	}
}

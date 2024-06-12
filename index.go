package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Option 1: Using the os package
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error retrieving hostname: %v\n", err)
	} else {
		// fmt.Printf("Hostname using os package: %s\n", hostname)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World! from %s", hostname)
		})
		fmt.Println("Starting api server...")
		err := http.ListenAndServe(":8000", nil)
		fmt.Println("Started api server on http://localhost:8000")
		if err != nil {
			fmt.Printf("Error starting server: %v\n", err)
			os.Exit(1)
		}
	}
}

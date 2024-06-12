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
		http.ListenAndServe(":9090", nil)
	}
}

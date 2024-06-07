package main

import (
	"os"
	"testing"
)

func TestHostname(t *testing.T) {
	hostname, err := os.Hostname()
	if err != nil {
		t.Fatalf("Failed to get hostname: %v", err)
	}
	if hostname == "" {
		t.Fatal("Hostname should not be empty")
	}
}

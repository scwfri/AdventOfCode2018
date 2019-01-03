package main

import "testing"

func testChecksum(t *testing.T) {
	if checksum(testInput) != 12 {
		t.Error("Expected checksum of 12 for testInput")
	}
	if 1 != 2 {
		t.Error("Test error")
	}
}

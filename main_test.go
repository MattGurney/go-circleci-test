package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if 1+1 != 2 {
		t.Errorf("Expected 2, got %d", 1+1)
	}
}

func TestFailure(t *testing.T) {
	if 1+1 != 3 {
		t.Errorf("Expected 3, got %d (This test is supposed to fail)", 1+1)
	}
}

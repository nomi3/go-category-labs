package main

import (
	"testing"
)

// TestSystemJoin tests the joining of two systems and checks connectivity.
func TestSystemJoin(t *testing.T) {
	// Create first system and connect A to B
	s1 := NewSystem()
	s1.Connect("A", "B")

	// Create second system and connect B to C
	s2 := NewSystem()
	s2.Connect("B", "C")

	// Join the two systems
	joinedSystem := Join(s1, s2)

	// Test if A and C are connected in the joined system
	if !joinedSystem.IsConnected("A", "C") {
		t.Errorf("Expected A and C to be connected in the joined system, but they are not")
	}
}

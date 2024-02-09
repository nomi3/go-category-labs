package main

import (
	"fmt"
)

// Point represents a point in the system.
type Point string

// System represents a collection of points and connections between them.
type System struct {
	connections map[Point][]Point
}

// NewSystem initializes a new System.
func NewSystem() *System {
	return &System{
		connections: make(map[Point][]Point),
	}
}

// Connect connects two points in the system.
func (s *System) Connect(p1, p2 Point) {
	s.connections[p1] = append(s.connections[p1], p2)
	s.connections[p2] = append(s.connections[p2], p1) // Assuming undirected connections
}

// IsConnected checks if two points are connected using BFS.
func (s *System) IsConnected(p1, p2 Point) bool {
	visited := make(map[Point]bool)
	queue := []Point{p1}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // Dequeue

		if current == p2 {
			return true
		}

		if visited[current] {
			continue
		}
		visited[current] = true

		for _, neighbor := range s.connections[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor) // Enqueue
			}
		}
	}

	return false
}

// Join combines two systems into one by merging their connections.
func Join(s1, s2 *System) *System {
	joined := NewSystem()
	// Combine connections from s1 and s2 into joined
	// This can be more complex, ensuring transitive closure is maintained
	return joined
}

func main() {
	// Example usage
	s1 := NewSystem()
	s1.Connect("A", "B")

	s2 := NewSystem()
	s2.Connect("B", "C")

	joinedSystem := Join(s1, s2)

	fmt.Println(joinedSystem.IsConnected("A", "C")) // Should print true if implementation is complete
}

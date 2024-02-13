package main

import (
	"fmt"
)

type Point string

type System struct {
	connections map[Point][]Point
}

func NewSystem() *System {
	return &System{
		connections: make(map[Point][]Point),
	}
}

func (s *System) Connect(p1, p2 Point) {
	s.connections[p1] = append(s.connections[p1], p2)
	s.connections[p2] = append(s.connections[p2], p1)
}

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

	// Add connections from s1 to the joined system
	for p1, connections := range s1.connections {
		for _, p2 := range connections {
			joined.Connect(p1, p2)
		}
	}

	// Add connections from s2 to the joined system
	for p1, connections := range s2.connections {
		for _, p2 := range connections {
			// Connect checks if the connection already exists; if not, it adds it
			joined.Connect(p1, p2)
		}
	}

	return joined
}


func main() {
	s1 := NewSystem()
	s1.Connect("A", "B")

	s2 := NewSystem()
	s2.Connect("B", "C")

	joinedSystem := Join(s1, s2)

	fmt.Println(joinedSystem.IsConnected("A", "C"))
}

package main

// System represents a collection of points and connections between them.
type System struct {
	connections map[string][]string
}

// NewSystem initializes a new System with no connections.
func NewSystem() *System {
	return &System{
		connections: make(map[string][]string),
	}
}

func (s *System) AddConnection(point string, connections []string) {
	s.connections[point] = connections
}

// Connect creates a bidirectional connection between two points.
func (s *System) Connect(p1, p2 string) {
	// Avoid duplicating connections
	if !s.IsConnected(p1, p2) {
		s.connections[p1] = append(s.connections[p1], p2)
	}
	if !s.IsConnected(p2, p1) {
		s.connections[p2] = append(s.connections[p2], p1)
	}
}

func (s *System) IsConnected(p1, p2 string) bool {
	visited := make(map[string]bool)
	return s.checkConnectivity(p1, p2, visited)
}

// checkConnectivity is a helper function that uses depth-first search to check connectivity.
func (s *System) checkConnectivity(current, target string, visited map[string]bool) bool {
	if current == target {
		return true
	}

	if visited[current] {
		return false
	}

	visited[current] = true

	for _, neighbor := range s.connections[current] {
		if s.checkConnectivity(neighbor, target, visited) {
			return true
		}
	}

	return false
}

// Join combines two systems into one by merging their connections.
func Join(s1, s2 *System) *System {
	joined := NewSystem()
	for p, connections := range s1.connections {
		for _, p2 := range connections {
			joined.Connect(p, p2)
		}
	}
	for p, connections := range s2.connections {
		for _, p2 := range connections {
			joined.Connect(p, p2)
		}
	}
	return joined
}

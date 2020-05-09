package icarus

import "sync"

// Thread safe graph to work with goroutines
type ThreadSafeGraph struct {
	Graph map[string][]string
	*sync.RWMutex
}

// Create a new thread safe graph
func NewThreadSafeGraph() *ThreadSafeGraph {
	return &ThreadSafeGraph{Graph: make(map[string][]string)}
}

// Add a new edge to this graph
func (graph *ThreadSafeGraph) addEdge(from string, to string) {
	graph.Graph[from] = append(graph.Graph[from], to)
}

// Check if a node exists
func (graph *ThreadSafeGraph) nodeExists(node string) bool {
	_, ok := graph.Graph[node]
	return ok
}

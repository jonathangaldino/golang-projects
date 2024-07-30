package Graph

import (
	Queue "algorithms/queue"
	"fmt"
)

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(src, dest int) {
	g.adjacencyList[src] = append(g.adjacencyList[src], dest)
	g.adjacencyList[dest] = append(g.adjacencyList[dest], src) // For undirected graph
}

// BFS performs Breadth-First Search on the graph starting from a given node
func (g *Graph) BFS(startNode int) {
	visited := make(map[int]bool)
	queue := Queue.NewQueue[int]()

	visited[startNode] = true
	queue.Enqueue(startNode)

	for !queue.IsEmpty() {
		currentNode, _ := queue.Dequeue()
		fmt.Printf("Visited %d\n", currentNode)

		for _, neighbor := range g.adjacencyList[currentNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.Enqueue(neighbor)
			}
		}

	}
}

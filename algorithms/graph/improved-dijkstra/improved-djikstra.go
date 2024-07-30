package ImprovedDjikstra

import (
	PriorityQueue "algorithms/priority-queue"
	"container/heap"
	"math"
)

type Edge struct {
	node   int
	weight int
}

type Graph struct {
	adjacencyList map[int][]Edge
}

func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[int][]Edge)}
}

func (g *Graph) AddEdge(source, destination, weight int) {
	g.adjacencyList[source] = append(g.adjacencyList[source], Edge{node: destination, weight: weight})
}

func (g *Graph) Dijkstra(start int) map[int]int {
	distances := make(map[int]int)
	for node := range g.adjacencyList {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0

	pq := make(PriorityQueue.PriorityQueue, 0)

	heap.Init(&pq)
	heap.Push(&pq, &PriorityQueue.Item{Node: start, Priority: 0})

	visited := make(map[int]bool)

	// Loop while the priority queue has items
	for pq.Len() > 0 {
		// Pull from the heap
		current := heap.Pop(&pq).(*PriorityQueue.Item)
		currentNode := current.Node

		// If the node is already visited, we skip it
		if visited[currentNode] {
			continue
		}

		// If not, lets visit it then
		visited[currentNode] = true

		// Loop through all edges of the current node
		for _, edge := range g.adjacencyList[currentNode] {
			// Take only the unvisited ones
			if !visited[edge.node] {
				// Calculate the distance
				newDist := distances[currentNode] + edge.weight

				if newDist < distances[edge.node] {
					// Update the distance if it cost less
					distances[edge.node] = newDist

					// Push the less expensive node to the priority queue
					heap.Push(&pq, &PriorityQueue.Item{Node: edge.node, Priority: newDist})
				}
			}
		}

		// The loop should end when we have calculated all the costs
		// or visited all the necessary nodes (the less expensive ones)
	}

	return distances
}

package Dijkstra

import "math"

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

// Dijkstra encontra o caminho mais curto de um nó de origem para todos os outros nós no grafo.
func (g *Graph) Dijkstra(start int) map[int]int {
	distances := make(map[int]int)
	for node := range g.adjacencyList {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0

	visited := make(map[int]bool)
	for len(visited) < len(g.adjacencyList) {
		// Encontre o nó não visitado com a menor distância
		currentNode := -1
		currentDist := math.MaxInt32

		for node, dist := range distances {
			if !visited[node] && dist < currentDist {
				currentNode = node
				currentDist = dist
			}
		}

		if currentNode == -1 {
			break
		}

		visited[currentNode] = true

		// Update neighbors' distance
		// This can be improved using a priority queue
		for _, edge := range g.adjacencyList[currentNode] {
			if !visited[edge.node] {
				newDist := distances[currentNode] + edge.weight
				if newDist < distances[edge.node] {
					distances[edge.node] = newDist
				}
			}
		}
	}

	return distances
}

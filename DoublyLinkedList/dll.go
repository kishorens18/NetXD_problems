package main

import (
	"fmt"
	"math"
)

type Edge struct {
	to     int
	weight int
}

func Dijkstra(graph map[int][]Edge, start int) map[int]int {
	distances := make(map[int]int)
	visited := make(map[int]bool)

	// Initialize distances and visited map
	for node := range graph {
		distances[node] = math.MaxInt32
		visited[node] = false
	}
	distances[start] = 0

	for {
		// Find the unvisited node with the smallest distance
		minDistance := math.MaxInt32
		var minNode int
		for node, distance := range distances {
			if !visited[node] && distance < minDistance {
				minDistance = distance
				minNode = node
			}
		}

		// If all nodes have been visited, exit
		if minDistance == math.MaxInt32 {
			break
		}

		visited[minNode] = true

		// Update distances to neighboring nodes
		for _, edge := range graph[minNode] {
			if newDistance := distances[minNode] + edge.weight; newDistance < distances[edge.to] {
				distances[edge.to] = newDistance
			}
		}
	}

	return distances
}

func main() {
	graph := map[int][]Edge{
		0: {{1, 4}, {2, 2}},
		1: {{3, 5}},
		2: {{1, 1}, {3, 8}},
		3: {},
	}

	startNode := 0
	distances := Dijkstra(graph, startNode)

	fmt.Println("Shortest distances from node", startNode, ":")
	for node, distance := range distances {
		fmt.Printf("Node %d: %d\n", node, distance)
	}
}

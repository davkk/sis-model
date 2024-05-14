package networks

import (
	"math/rand"
)

type Node[T any] struct {
	K     int
	Value T
}

type Graph[T any] struct {
	Nodes   []Node[T]
	AdjList [][]int
}

func (graph *Graph[T]) MaxDegreeNode() int {
	curr := 0
	for idx := 0; idx < len(graph.Nodes); idx++ {
		if graph.Nodes[curr].K < graph.Nodes[idx].K {
			curr = idx
		}
	}
	return curr
}

func (graph *Graph[T]) MinDegreeNode() int {
	curr := 0
	for idx := 0; idx < len(graph.Nodes); idx++ {
		if graph.Nodes[curr].K > graph.Nodes[idx].K {
			curr = idx
		}
	}
	return curr
}

func BarabasiAlbert[T any](rng *rand.Rand, N int, m int, defaultValue T) Graph[T] {
	nodes := make([]Node[T], N)
	adjList := make([][]int, N)
	nodePool := []int{}

	// initial nodes
	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			adjList[i] = append(adjList[i], j)
			adjList[j] = append(adjList[j], i)
			nodePool = append(nodePool, i, j)
		}
	}

	// populate graph
	for t := m; t < N; t++ {
		edges := []int{}
		currNode := &nodes[t]

		for len(edges) < m {
			randNode := nodePool[rng.Intn(len(nodePool))]
			for node := range edges {
				if node == randNode {
					continue
				}
			}

			nodePool = append(nodePool, randNode, t)
			edges = append(edges, randNode)

			adjList[randNode] = append(adjList[randNode], t)
			nodes[randNode].K++
		}

		adjList[t] = edges

		currNode.K = 2
		currNode.Value = defaultValue
	}

	return Graph[T]{
		Nodes:   nodes,
		AdjList: adjList,
	}
}

// TODO erdos renyi

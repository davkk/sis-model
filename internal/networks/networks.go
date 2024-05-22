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
		currNode := graph.Nodes[curr]
		if currNode.K > graph.Nodes[idx].K {
			curr = idx
		}
	}
	return curr
}

func BarabasiAlbert[T any](rng *rand.Rand, n int, m int, defaultValue T) Graph[T] {
	nodes := make([]Node[T], n)
	adjList := make([][]int, n)
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
	for t := m; t < n; t++ {
		edges := []int{}
		tmpPool := []int{}

	outer:
		for len(edges) < m {
			randNode := nodePool[rng.Intn(len(nodePool))]
			for _, node := range edges {
				if node == randNode {
					continue outer
				}
			}

			tmpPool = append(tmpPool, randNode, t)
			edges = append(edges, randNode)

			adjList[randNode] = append(adjList[randNode], t)
			nodes[randNode].K++
		}

		nodePool = append(nodePool, tmpPool...)
		adjList[t] = edges

		currNode := &nodes[t]
		currNode.K = 2
		currNode.Value = defaultValue
	}

	return Graph[T]{
		Nodes:   nodes,
		AdjList: adjList,
	}
}

func ErdosRenyi[T any](rng *rand.Rand, n int, p float32, defaultValue T) Graph[T] {
	nodes := make([]Node[T], n)
	adjList := make([][]int, n)

	for i := 0; i < n; i++ {
		nodes[i].Value = defaultValue
		nodes[i].K = 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if rng.Float32() < p {
				adjList[i] = append(adjList[i], j)
				nodes[i].K++

				adjList[j] = append(adjList[j], i)
				nodes[j].K++
			}
		}
	}

	return Graph[T]{
		Nodes:   nodes,
		AdjList: adjList,
	}
}

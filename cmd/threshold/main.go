package main

import (
	"fmt"
	"sis-model/internal/networks"
)

func main() {
	n := 1000
	m := 2

	graph := networks.BarabasiAlbert(n, m, 0)

	k := 0.0
	k2 := 0.0

	for _, node := range graph.Nodes {
		k += float64(node.K)
		k2 += float64(node.K * node.K)
	}

	k /= float64(len(graph.Nodes))
	k2 /= float64(len(graph.Nodes))

	fmt.Println(k, k2, k/k2)
}

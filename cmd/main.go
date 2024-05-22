package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"sis-model/internal/networks"
)

type State byte

const (
	Susceptible State = iota
	Infected
)

type Config struct {
	Network     string  `json:"network"`
	Init        string  `json:"init"`
	N           int     `json:"n"`
	M           int     `json:"m"`
	P           float32 `json:"p"`
	Beta        float32 `json:"beta"`
	Gamma       float32 `json:"gamma"`
	rng         *rand.Rand
}

func parseConfig() Config {
	if len(os.Args) != 2 {
		panic("Usage: ./main.out path/to/config")
	}

	path := os.Args[1]
	var config Config

	file, err := os.ReadFile(path)
	if err != nil {
		panic("error while reading config file")
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		panic("error while parsing JSON")
	}

	config.rng = rand.New(rand.NewSource(2001))

	return config
}

func prob(p float32, k int, maxK int) float32 {
	if p < 0 {
		return float32(k) / float32(maxK)
	}
	return p
}

func main() {
	sim := parseConfig()
	steps := int(1e3)

	if sim.N == 0 {
		panic("n must be greater than zero")
	}

	var graph networks.Graph[State]
	switch sim.Network {
	case "ba":
		graph = networks.BarabasiAlbert(sim.rng, sim.N, sim.M, Susceptible)
		if sim.M == 0 || sim.M < 2 {
			panic("m must be greater than 2")
		}
	case "er":
		graph = networks.ErdosRenyi(sim.rng, sim.N, sim.P, Susceptible)
		if sim.P == 0 {
			panic("p must be greater than 0")
		}
	default:
		panic("invalid network type")
	}

	switch sim.Init {
	case "rand":
		graph.Nodes[sim.rng.Intn(sim.N)].Value = Infected
	case "max":
		graph.Nodes[graph.MaxDegreeNode()].Value = Infected
	case "min":
		graph.Nodes[graph.MinDegreeNode()].Value = Infected
	default:
		panic("invalid init type")
	}

	infected := 1
	maxK := graph.Nodes[graph.MaxDegreeNode()].K

	for step := 0; step < steps; step++ {
		fmt.Println(step, infected, sim.N-infected)

		for idx := 0; idx < sim.N; idx++ {
			node := &graph.Nodes[idx]

			if node.K > 0 && node.Value == Infected {
				for _, nnIdx := range graph.AdjList[idx] {
					nnNode := &graph.Nodes[nnIdx]

					if nnNode.Value == Susceptible && sim.rng.Float32() < prob(sim.Beta, node.K, maxK) {
						nnNode.Value = Infected
						infected++
					}
				}

				if sim.rng.Float32() < prob(sim.Gamma, node.K, maxK) {
					node.Value = Susceptible
					infected--
				}
			}
		}
	}
}

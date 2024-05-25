package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"

	"sis-model/internal/networks"
)

type State byte

const (
	Susceptible State = iota
	Infected
)

type Simulation struct {
	NetworkType string
	InitType    string
	KDep        string
	N           int
	M           int
	P           float64
	Beta        float64
	Gamma       float64
}

func (sim *Simulation) parseConfig() {
	flag.StringVar(&sim.NetworkType, "network", "", "network type")
	flag.StringVar(&sim.InitType, "init", "", "initial infection type")
	flag.StringVar(&sim.KDep, "k", "none", "dependece on node degree")
	flag.IntVar(&sim.N, "n", 0, "number of nodes")
	flag.IntVar(&sim.M, "m", 0, "BA: m parameter")
	flag.Float64Var(&sim.P, "p", 0, "ER: p parameter")
	flag.Float64Var(&sim.Beta, "beta", 0, "SIS: beta parameter")
	flag.Float64Var(&sim.Gamma, "gamma", 0, "SIS: gamma parameter")
	flag.Parse()

	sim.NetworkType = strings.ToLower(sim.NetworkType)
	sim.InitType = strings.ToLower(sim.InitType)
	sim.KDep = strings.ToLower(sim.KDep)
}

func (sim *Simulation) beta(k int, maxK int) float64 {
	if sim.KDep == "beta" || sim.KDep == "both" {
		return sim.Beta * float64(k) / float64(maxK)
	}
	return sim.Beta
}

func (sim *Simulation) gamma(k int, maxK int) float64 {
	if sim.KDep == "gamma" || sim.KDep == "both" {
		return sim.Gamma * float64(k) / float64(maxK)
	}
	return sim.Gamma
}

func main() {
	var sim Simulation
	sim.parseConfig()

	steps := int(1e3)

	if sim.N == 0 {
		panic("n must be greater than zero")
	}

	switch sim.KDep {
	case "both":
	case "beta":
	case "gamma":
	case "none":
	default:
		panic("invalid k dependence value")
	}

	var graph networks.Graph[State]
	switch sim.NetworkType {
	case "ba":
		graph = networks.BarabasiAlbert(sim.N, sim.M, Susceptible)
		if sim.M == 0 || sim.M < 2 {
			panic("m must be greater than 2")
		}
	case "er":
		graph = networks.ErdosRenyi(sim.N, sim.P, Susceptible)
		if sim.P == 0 {
			panic("p must be greater than 0")
		}
	default:
		panic("invalid network type")
	}

	switch sim.InitType {
	case "rnd":
		graph.Nodes[rand.Intn(sim.N)].Value = Infected
	case "max":
		graph.Nodes[graph.MaxDegreeNode()].Value = Infected
	case "min":
		graph.Nodes[graph.MinDegreeNode()].Value = Infected
	default:
		panic("invalid init type")
	}

	infected := 1
	maxK := graph.Nodes[graph.MaxDegreeNode()].K
	lambda := float64(sim.Beta) / float64(sim.Gamma)

	for step := 0; step < steps; step++ {
		fmt.Println(lambda, step, float64(infected)/float64(sim.N))

		for idx := 0; idx < sim.N; idx++ {
			node := &graph.Nodes[idx]

			if node.K > 0 && node.Value == Infected {
				for _, nnIdx := range graph.AdjList[idx] {
					nnNode := &graph.Nodes[nnIdx]
					if nnNode.Value == Infected {
						continue
					}

					if rand.Float64() < sim.beta(node.K, maxK) {
						nnNode.Value = Infected
						infected++
					}
				}

				if rand.Float64() < sim.gamma(node.K, maxK) {
					node.Value = Susceptible
					infected--
				}
			}
		}
	}
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"

	"sis-model/internal/networks"
)

type State byte

const (
	Susceptible State = iota
	Infected
)

type Simulation struct {
	Network string  `json:"network"`
	Init    string  `json:"init"`
	Steps   int     `json:"steps"`
	N       int     `json:"n"`
	M       int     `json:"m"`
	Beta    float32 `json:"beta"`
	Gamma   float32 `json:"gamma"`
	rng     *rand.Rand
}

func parseConfig() Simulation {
	var configFile string
	var sim Simulation

	flag.StringVar(&configFile, "config", "./config.json", "Path to config file")
	flag.Parse()

	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(file, &sim)
	if err != nil {
		log.Fatal("Could not parse config file: ", err)
	}

	sim.rng = rand.New(rand.NewSource(2001))
	return sim
}

func (sim *Simulation) infect(node1 *networks.Node[State], node2 *networks.Node[State]) {
	x := sim.rng.Float32()
	if sim.Beta == 0 && x < 1-1/float32(node2.K) {
		node1.Value = Infected
	} else if x < sim.Beta {
		node1.Value = Infected
	}
}

func (sim *Simulation) restore(node1 *networks.Node[State], node2 *networks.Node[State]) {
	x := sim.rng.Float32()
	if sim.Gamma == 0 && x < 1/float32(node2.K) {
		node1.Value = Susceptible
	} else if x < sim.Gamma {
		node1.Value = Susceptible
	}
}

func main() {
	sim := parseConfig()

	var graph networks.Graph[State]
	switch sim.Network {
	case "ba":
		graph = networks.BarabasiAlbert(sim.rng, sim.N, sim.M, Susceptible)
	case "er":
		// TODO erdos renyi
	default:
		log.Fatal("Invalid network type")
	}

	switch sim.Init {
	case "rand":
		graph.Nodes[sim.rng.Intn(sim.N)].Value = Infected
		break
	case "max":
		graph.Nodes[graph.MaxDegreeNode()].Value = Infected
		break
	case "min":
		graph.Nodes[graph.MinDegreeNode()].Value = Infected
		break
	default:
		log.Fatal("Invalid init type")
		break
	}
	infected := 1

	for step := 0; step < sim.Steps && infected >= 0 && infected <= sim.N; step++ {
		fmt.Println(step, sim.N-infected, infected)

		idx1 := sim.rng.Intn(sim.N)
		node1 := &graph.Nodes[idx1]

		idx2 := graph.AdjList[idx1][sim.rng.Intn(node1.K)]
		node2 := &graph.Nodes[idx2]

		switch {
		case node1.Value == Susceptible && node2.Value == Infected:
			sim.infect(node1, node2)
			infected++
			break
		case node1.Value == Infected && node2.Value == Susceptible:
			sim.infect(node2, node1)
			infected++
			break
		case node1.Value == Infected && node2.Value == Infected:
			sim.restore(node1, node2)
			infected--
			break
		default:
			break
		}
	}
}

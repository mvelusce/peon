package main

// thank you Igor

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yourbasic/graph"
)

type Modules struct {
	Name string
	Path string
	Deps []string
}

type Configuration struct {
	Modules []Modules
}

func readConfigs() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Modules)

	//m := []string{"module_a", "module_b", "module_c", "module_d", "module_e"}

	g := graph.New(len(configuration.Modules))

	indexes := make(map[string]int)
	for m := 0; m < len(configuration.Modules); m++ {
		indexes[configuration.Modules[m].Name] = m
	}

	for m := 0; m < len(configuration.Modules); m++ {
		for d := 0; d < len(configuration.Modules[m].Deps); d++ {
			g.Add(m, indexes[configuration.Modules[m].Deps[d]])
		}
	}

	/*g.Add(2, 3)
	g.Add(0, 1)
	g.Add(1, 4)
	g.Add(0, 2)
	g.Add(1, 3)
	g.Add(3, 4)*/

	fmt.Println(graph.Acyclic(g))

	//fmt.Println(graph.Sort(g))

	for v := 0; v < g.Order(); v++ {
		fmt.Println(v)
		fmt.Println("Install: ", configuration.Modules[v].Name)
	}
}

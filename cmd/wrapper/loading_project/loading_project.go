package loading_project

import (
	"github.com/yourbasic/graph"
	"log"
)

func loadProject() ([]PyModule, *graph.Mutable) {

	modules := LoadModules("testdata")

	g := graph.New(len(modules))
	indexes := make(map[string]int)
	for i, m := range modules {
		indexes[m.Name] = i
	}

	for i, m := range modules {
		for _, d := range m.Dependencies {
			g.Add(i, indexes[d])
		}
	}

	if !graph.Acyclic(g) {
		log.Fatalf("ERROR Cyrcular dependency detected")
	}

	return modules, g
}

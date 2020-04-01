package project

import (
	"fmt"
	"github.com/yourbasic/graph"
	"log"
)

type PyProject struct {
	modules      []PyModule
	dependencies *graph.Mutable
}

func LoadProject() PyProject {

	modules := loadModules("testdata")

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

	return PyProject{modules, g}
}

func (p *PyProject) Build() {

	for v := 0; v < p.dependencies.Order(); v++ {
		m := p.modules[v]
		fmt.Println(fmt.Sprintf("python %s/setup.py install", m.Path))
		/*err := runCommand(fmt.Sprintf("python %s/setup.py install", m.Path))// TODO pass cmd as dependency
		if err != nil {
			fmt.Println("Unable to build module {}", err)
		}
		fmt.Println("Install: ", p.modules[v].Name)*/
	}
}

func (p *PyProject) BuildModule(module string) {

	index := 0
	for v := 0; v < p.dependencies.Order(); v++ {
		m := p.modules[v]
		if m.Name == module {
			index = v
		}
	}

	b := func(w int, c int64) bool {
		m := p.modules[w]
		fmt.Println(fmt.Sprintf("python %s/setup.py install", m.Path))
		return false
	}

	res := p.dependencies.Visit(index, b)

	if !res {
		m := p.modules[index]
		fmt.Println(fmt.Sprintf("python %s/setup.py install", m.Path))
	}
}

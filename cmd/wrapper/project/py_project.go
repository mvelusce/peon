package project

import (
	"fmt"
	"github.com/skyveluscekm/setuptools.wrapper/cmd/wrapper/executor"
	"github.com/yourbasic/graph"
	"log"
)

type PyProject struct {
	modules      []PyModule
	dependencies *graph.Mutable
	executor     executor.PyExecutor
}

func LoadProject() PyProject {

	projectRoot := "testdata"
	pythonVersion := "python3.7"
	modules, g := loadModulesAndGraph(projectRoot)

	e := &executor.SetupPyExecutor{PyVersion: pythonVersion}
	return PyProject{modules, g, e}
}

func loadModulesAndGraph(projectRoot string) ([]PyModule, *graph.Mutable) {
	modules := loadModules(projectRoot)
	g := loadDependenciesGraph(modules)
	return modules, g
}

func loadDependenciesGraph(modules []PyModule) *graph.Mutable {
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
	return g
}

func (p *PyProject) Build() {

	order, ac := graph.TopSort(p.dependencies)
	if !ac {
		log.Fatalf("ERROR Cyrcular dependency detected")
	}

	for v := 0; v < len(order); v++ {
		i := len(order) - v - 1

		m := p.modules[order[i]]
		err := p.executor.Build(m.Path)

		if err != nil {
			log.Fatalf("Unable to build module %s. Error: %v", m.Path, err)
		}
		log.Printf("Install module %s successful", p.modules[v].Name)
	}
}

func (p *PyProject) BuildModule(module string) {

	index := p.findIndex(module)

	depsError := p.buildDependencies(index)

	p.buildModule(depsError, index)
}

func (p *PyProject) buildDependencies(index int) bool {

	// TODO use depth first search to build modules with no deps first

	b := func(w int, c int64) bool {
		println("QWEQWE ", w)
		m := p.modules[w]

		err := p.executor.Build(m.Path)
		if err != nil {
			fmt.Println("Unable to build module {}. Error: {}", m.Path, err)
		}
		fmt.Println("Install module {} successful", m.Name)
		return err != nil
	}
	depsError := p.dependencies.Visit(index, b)
	return depsError
}

func (p *PyProject) buildModule(depsError bool, index int) {
	if !depsError {
		m := p.modules[index]
		err := p.executor.Build(m.Path)
		if err != nil {
			fmt.Println("Unable to build module {}. Error: {}", m.Path, err)
		}
		fmt.Println("Install module {} successful", m.Name)
	}
}

func (p *PyProject) findIndex(module string) int {
	index := 0
	for v := 0; v < p.dependencies.Order(); v++ {
		m := p.modules[v]
		if m.Name == module {
			index = v
		}
	}
	println("ASDASD ", index)
	return index
}

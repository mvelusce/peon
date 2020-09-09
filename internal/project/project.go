package project

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/skyveluscekm/setuptools.wrapper/internal/executor"
	"github.com/yourbasic/graph"
)

type PyProject struct {
	modules      []PyModule
	dependencies *graph.Mutable
	executor     executor.PyExecutor
}

func LoadProject(projectRoot string, pythonVersion string) (PyProject, error) {

	if projectRoot == "" {
		projectRoot = "."
	}
	if pythonVersion == "" {
		pythonVersion = "python3.7"
	} // TODO save some where default configs ??

	modules, g, err := loadModulesAndGraph(projectRoot)
	e := &executor.SetupPyExecutor{PyVersion: pythonVersion}
	return PyProject{modules, g, e}, err
}

func loadModulesAndGraph(projectRoot string) ([]PyModule, *graph.Mutable, error) {
	modules, err := loadModules(projectRoot)
	if err != nil {
		log.Errorf("Unable to load modules. Error: %v", err)
		return nil, nil, err
	}
	g, err := loadDependenciesGraph(modules)
	if err != nil {
		return nil, nil, err
	}
	return modules, g, nil
}

func loadDependenciesGraph(modules []PyModule) (*graph.Mutable, error) {
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
		log.Errorf("ERROR Circular dependency detected")
		return nil, errors.New("ERROR Circular dependency detected")
	}
	return g, nil
}

func (p *PyProject) Build() error {

	order, ac := graph.TopSort(p.dependencies)
	if !ac {
		log.Errorf("ERROR Circular dependency detected")
		return errors.New("ERROR Circular dependency detected")
	}

	for v := 0; v < len(order); v++ {
		i := len(order) - v - 1

		m := p.modules[order[i]]
		err := p.executor.Build(m.Path) // TODO make parametric on the executor function ??

		if err != nil {
			log.Errorf("Unable to build module %s. Error: %v", m.Path, err)
			return err
		}
		log.Printf("Install module %s successful", p.modules[v].Name)
	}
	return nil
}

func (p *PyProject) BuildModule(module string) error {

	index := p.findIndex(module)

	visited := p.setupVisited()

	return p.buildDependencies(index, visited)
}

func (p *PyProject) Clean() {
	// TODO
}

func (p *PyProject) Test() {
	// TODO
}

func (p *PyProject) TestModule(module string) {
	// TODO
}

func (p *PyProject) setupVisited() []bool {
	visited := make([]bool, p.dependencies.Order())
	for v := 0; v < p.dependencies.Order(); v++ {
		visited[v] = false
	}
	return visited
}

func (p *PyProject) buildDependencies(index int, visited []bool) error {

	b := func(w int, c int64) (skip bool) {
		if !visited[w] {
			err := p.buildDependencies(w, visited)
			if err != nil {
				return
			}
		}
		return
	}
	p.dependencies.Visit(index, b)

	m := p.modules[index]
	err := p.executor.Build(m.Path)

	if err != nil {
		log.Errorf("Unable to build module %s. Error: %v", m.Path, err)
		return err
	}

	visited[index] = true
	log.Printf("Install module %s successful", m.Name)
	return nil
}

func (p *PyProject) buildModule(index int) error {
	m := p.modules[index]
	err := p.executor.Build(m.Path)
	if err != nil {
		log.Errorf("Unable to build module %s. Error: %v", m.Path, err)
		return err
	}
	log.Printf("Install module %s successful", m.Name)
	return nil
}

func (p *PyProject) findIndex(module string) int {
	index := 0
	for v := 0; v < p.dependencies.Order(); v++ {
		m := p.modules[v]
		if m.Name == module {
			index = v
		}
	}
	return index
}

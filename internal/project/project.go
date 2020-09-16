package project

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/skyveluscekm/peon/internal/executor"
	"github.com/yourbasic/graph"
)

type Project struct {
	modules      []Module
	dependencies *graph.Mutable
	executor     executor.Executor
}

func LoadProject(config *Config) (Project, error) {

	modules, g, err := loadModulesAndGraph(config.ProjectRoot)
	e := &executor.SetupPyExecutor{PyVersion: config.PythonExec}
	return Project{modules, g, e}, err
}

func loadModulesAndGraph(projectRoot string) ([]Module, *graph.Mutable, error) {
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

func loadDependenciesGraph(modules []Module) (*graph.Mutable, error) {
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

func (p *Project) Build() error {
	return p.executeOnAll(p.executor.Build)
}

func (p *Project) BuildModule(module string) error {

	index := p.findIndex(module)

	visited := p.setupVisited()

	return p.executeOnDependencies(index, visited, p.executor.Build)
}

func (p *Project) Clean() error {
	return p.executor.Clean()
}

func (p *Project) Test() error {
	err := p.Build()
	if err != nil {
		log.Errorf("Unable to test all. Error: %v", err)
		return err
	}
	return p.executeOnAll(p.executor.Test)
}

func (p *Project) TestModule(module string) error {

	err := p.BuildModule(module)
	if err != nil {
		log.Errorf("Unable to test module: %s. Error: %v", module, err)
		return err
	}

	index := p.findIndex(module)

	visited := p.setupVisited()

	return p.executeOnDependencies(index, visited, p.executor.Test)
}

func (p *Project) Exec(command string) error {

	exec := func(path string) error {
		return p.executor.Exec(command, path)
	}
	return p.executeOnAll(exec)
}

func (p *Project) ExecModule(command string, module string) error {

	index := p.findIndex(module)

	visited := p.setupVisited()

	exec := func(path string) error {
		return p.executor.Exec(command, path)
	}
	return p.executeOnDependencies(index, visited, exec)
}

func (p *Project) executeOnAll(action func(string) error) error {
	order, ac := graph.TopSort(p.dependencies)
	if !ac {
		log.Errorf("ERROR Circular dependency detected")
		return errors.New("ERROR Circular dependency detected")
	}

	for v := 0; v < len(order); v++ {
		i := len(order) - v - 1

		m := p.modules[order[i]]
		err := action(m.Path)

		if err != nil {
			log.Errorf("Unable to build module %s. Error: %v", m.Path, err)
			return err
		}
		log.Printf("Install module %s successful", p.modules[v].Name)
	}
	return nil
}

func (p *Project) setupVisited() []bool {
	visited := make([]bool, p.dependencies.Order())
	for v := 0; v < p.dependencies.Order(); v++ {
		visited[v] = false
	}
	return visited
}

func (p *Project) executeOnDependencies(index int, visited []bool, action func(string) error) error {

	b := func(w int, c int64) (skip bool) {
		if !visited[w] {
			err := p.executeOnDependencies(w, visited, action)
			if err != nil {
				return
			}
		}
		return
	}
	p.dependencies.Visit(index, b)

	m := p.modules[index]
	err := action(m.Path)

	if err != nil {
		log.Errorf("Unable to build module %s. Error: %v", m.Path, err)
		return err
	}

	visited[index] = true
	log.Printf("Install module %s successful", m.Name)
	return nil
}

func (p *Project) findIndex(module string) int {
	index := 0
	for v := 0; v < p.dependencies.Order(); v++ {
		m := p.modules[v]
		if m.Name == module {
			index = v
		}
	}
	return index
}

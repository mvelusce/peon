package project

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/yourbasic/graph"
)

type DependenciesGraph struct {
	modules     []Module
	graph       *graph.Mutable
	moduleIndex map[string]int
}

func loadModulesAndGraph(projectRoot string) (*DependenciesGraph, error) {
	modules, err := loadModules(projectRoot)
	if err != nil {
		log.Errorf("Unable to load modules. Error: %v", err)
		return nil, err
	}
	dg, err := loadDependenciesGraph(modules)
	if err != nil {
		return nil, err
	}
	return dg, nil
}

func loadDependenciesGraph(modules []Module) (*DependenciesGraph, error) {
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
	depGraph := &DependenciesGraph{
		modules:     modules,
		graph:       g,
		moduleIndex: indexes,
	}
	return depGraph, nil
}

func (dp *DependenciesGraph) executeOnAll(action func(string) error) error {
	order, ac := graph.TopSort(dp.graph)
	if !ac {
		log.Errorf("ERROR Circular dependency detected")
		return errors.New("ERROR Circular dependency detected")
	}

	for v := 0; v < len(order); v++ {
		i := len(order) - v - 1

		m := dp.modules[order[i]]
		err := action(m.Path)

		if err != nil {
			log.Errorf("Unable to build module %s. Error: %v", m.Path, err)
			return err
		}
		log.Printf("Install module %s successful", dp.modules[v].Name)
	}
	return nil
}

func (dp *DependenciesGraph) setupVisited() []bool {
	visited := make([]bool, dp.graph.Order())
	for v := 0; v < dp.graph.Order(); v++ {
		visited[v] = false
	}
	return visited
}

func (dp *DependenciesGraph) executeOnDependencies(module string, action func(string) error) error {

	visited := dp.setupVisited()

	moduleIndex := dp.moduleIndex[module]

	return dp.visit(moduleIndex, visited, action)
}

func (dp *DependenciesGraph) visit(index int, visited []bool, action func(string) error) error {

	b := func(w int, c int64) (skip bool) {
		if !visited[w] {
			err := dp.visit(w, visited, action)
			if err != nil {
				return
			}
		}
		return
	}
	dp.graph.Visit(index, b)

	m := dp.modules[index]
	err := action(m.Path)

	if err != nil {
		log.Errorf("Unable to build module %s. Error: %v", m.Path, err)
		return err
	}

	visited[index] = true
	log.Printf("Install module %s successful", m.Name)
	return nil
}

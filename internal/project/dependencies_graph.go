package project

import (
	"container/heap"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/yourbasic/graph"
)

type DependenciesGraph struct {
	modules     []*Module
	graph       *graph.Mutable
	moduleIndex map[string]int
	inverseDeps map[string][]*Module
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

func loadDependenciesGraph(modules []*Module) (*DependenciesGraph, error) {
	g := graph.New(len(modules))
	indexes := make(map[string]int)
	inverseDeps := make(map[string][]*Module)
	for i, m := range modules {
		indexes[m.Name] = i
	}
	for i, m := range modules {
		for _, d := range m.Dependencies {
			g.Add(i, indexes[d])

			if val, ok := inverseDeps[d]; ok {
				inverseDepsForM := append(val, m)
				inverseDeps[d] = inverseDepsForM
			} else {
				inverseDepsForM := []*Module{m}
				inverseDeps[d] = inverseDepsForM
			}
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
		inverseDeps: inverseDeps,
	}
	return depGraph, nil
}

func (dp *DependenciesGraph) executeOnAll(action func(string) error) error {

	pq := Init(dp.modules)

	for pq.Len() > 0 {
		mod := heap.Pop(pq).(*ModulePriority)
		fmt.Printf("priority %.2d:%s ", mod.priority, mod.module)

		err := action(mod.module.Path)

		if err != nil {
			log.Errorf("Unable to build module %s. Error: %v", mod.module.Path, err)
			return err
		}
		log.Printf("Install module %s successful", mod.module.Name)

		for _, d := range dp.inverseDeps[mod.module.Name] {
			pq.Decrease(d.Name)
		}
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

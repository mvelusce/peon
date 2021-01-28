package project

import (
	"container/heap"
	"errors"
	
	log "github.com/sirupsen/logrus"
	"github.com/yourbasic/graph"
)

type DependenciesGraph struct {
	modules     []*Module
	graph       *graph.Mutable
	moduleIndex map[string]int
	inverseDeps map[string][]*Module
}

type modActionStatus struct {
	module   *Module
	succeded bool
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

	c := make(chan *modActionStatus)// TODO use mutex and wait group instead

	for pq.Len() > 0 {
		
		for pq.Head() == 0 {
			mod := heap.Pop(pq).(*ModulePriority)// TODO use mutex to sync on changes
			log.Debugf("Priority %.2d: %s ", mod.priority, mod.module.Name)
			
			go executeAction(mod, c, action)
		}

		actionResult := <-c

		if !actionResult.succeded {
			return errors.New("Unable to execute action on module " + actionResult.module.Name)
		}
		for _, d := range dp.inverseDeps[actionResult.module.Name] {
			pq.Decrease(d.Name)// TODO move to go routine after build done.
		}
	}
	// TODO wait all spawned go routines
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

func executeAction(mod *ModulePriority, c chan *modActionStatus, action func(string) error) {
	err := action(mod.module.Path)

	if err != nil {
		log.Errorf("Unable to execute action on module %s. Error: %v", mod.module.Path, err)
		c <- &modActionStatus{module: mod.module, succeded: false}
	}
	//log.Printf("Execution on module %s successful", mod.module.Name)
	c <- &modActionStatus{module: mod.module, succeded: true}
}

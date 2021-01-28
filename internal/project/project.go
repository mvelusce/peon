package project

import (
	log "github.com/sirupsen/logrus"
	"github.com/mvelusce/peon/internal/executor"
	"github.com/yourbasic/graph"
)

type Project struct {
	dependencies *DependenciesGraph
	executor     executor.Executor
}

func LoadProject(config *Config, dryRun bool) (Project, error) {

	dependenciesGraph, err := loadModulesAndGraph(config.ProjectRoot)

	if !dryRun {
		e := &executor.SetupPyExecutor{PyVersion: config.PythonExec}
		return Project{dependenciesGraph, e}, err
	} else {
		e := &DryRunExecutor{}
		return Project{dependenciesGraph, e}, err
	}
}

func (p *Project) Build() error {
	return p.dependencies.executeOnAll(p.executor.Build)
}

func (p *Project) BuildModule(module string) error {
	return p.dependencies.executeOnDependencies(module, p.executor.Build)
}

func (p *Project) Clean() error {
	return p.executor.Clean()
}

func (p *Project) Test() error {
	return p.dependencies.executeOnAll(p.test)
}

func (p *Project) TestModule(module string) error {
	return p.dependencies.executeOnDependencies(module, p.test)
}

func (p *Project) test(path string) error {
	err := p.executor.Build(path)
	if err != nil {
		log.Errorf("Unable to build path %s during test all. Error: %v", path, err)
		return err
	}
	err = p.executor.Test(path)
	if err != nil {
		log.Errorf("Unable to test path %s during test all. Error: %v", path, err)
	}
	return err
}

func (p *Project) Exec(command string) error {

	exec := func(path string) error {
		return p.executor.Exec(command, path)
	}
	return p.dependencies.executeOnAll(exec)
}

func (p *Project) ExecModule(command string, module string) error {

	exec := func(path string) error {
		return p.executor.Exec(command, path)
	}
	return p.dependencies.executeOnDependencies(module, exec)
}

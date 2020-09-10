package project

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildProject(t *testing.T) {

	var modules = []PyModule{
		{"mod3", "mod3", []string{"mod2"}},
		{"mod5", "mod5", []string{"mod3"}},
		{"mod0", "mod0", []string{}},
		{"mod1", "mod1", []string{"mod0"}},
		{"mod4", "mod4", []string{"mod2", "mod1"}},
		{"mod2", "mod2", []string{"mod1", "mod0"}},
	}

	g, _ := loadDependenciesGraph(modules)

	var buildModules []string
	e := &MockExecutor{executedActions: buildModules}
	project := PyProject{modules, g, e}

	_ = project.Build()

	assert.Equal(t, 6, len(e.executedActions))
	assert.Equal(t, "mod0", e.executedActions[0])
	assert.Equal(t, "mod1", e.executedActions[1])
	assert.Equal(t, "mod2", e.executedActions[2])
	assert.Equal(t, "mod3", e.executedActions[3])
	assert.Equal(t, "mod4", e.executedActions[4])
	assert.Equal(t, "mod5", e.executedActions[5])
}

func TestBuildModule(t *testing.T) {

	var modules = []PyModule{
		{"mod0", "mod0", []string{}},
		{"mod1", "mod1", []string{"mod0"}},
		{"mod2", "mod2", []string{"mod1", "mod0"}},
		{"mod4", "mod4", []string{"mod2", "mod1"}},
	}

	g, _ := loadDependenciesGraph(modules)

	var buildModules []string
	e := &MockExecutor{executedActions: buildModules}
	project := PyProject{modules, g, e}

	_ = project.BuildModule("mod2")

	assert.Equal(t, 3, len(e.executedActions))
	assert.Equal(t, "mod0", e.executedActions[0])
	assert.Equal(t, "mod1", e.executedActions[1])
	assert.Equal(t, "mod2", e.executedActions[2])
}

func TestBuildModule1(t *testing.T) {

	var modules = []PyModule{
		{"mod0", "mod0", []string{}},
		{"mod1", "mod1", []string{"mod0"}},
		{"mod2", "mod2", []string{"mod1", "mod0"}},
		{"mod4", "mod4", []string{"mod2", "mod1"}},
	}

	g, _ := loadDependenciesGraph(modules)

	var buildModules []string
	e := &MockExecutor{executedActions: buildModules}
	project := PyProject{modules, g, e}

	_ = project.BuildModule("mod4")

	assert.Equal(t, 4, len(e.executedActions))
	assert.Equal(t, "mod0", e.executedActions[0])
	assert.Equal(t, "mod1", e.executedActions[1])
	assert.Equal(t, "mod2", e.executedActions[2])
	assert.Equal(t, "mod4", e.executedActions[3])
}

func TestLoadDependenciesGraph(t *testing.T) {

	var modules = []PyModule{
		{"mod0", "mod0", []string{}},
		{"mod1", "mod1", []string{"mod0"}},
		{"mod2", "mod2", []string{"mod1", "mod0"}},
	}

	res, _ := loadDependenciesGraph(modules)

	assert.Equal(t, "3 [(1 0) (2 0) (2 1)]", res.String())
}

type MockExecutor struct {
	executedActions []string
}

func (e *MockExecutor) Build(path string) error {
	e.executedActions = append(e.executedActions, path)
	return nil
}

func (e *MockExecutor) Run(path string) error {
	return nil
}
func (e *MockExecutor) Clean() error {
	return nil
}
func (e *MockExecutor) Test(path string) error {
	return nil
}
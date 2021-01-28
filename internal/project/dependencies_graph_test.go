package project

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDependenciesGraph(t *testing.T) {

	var modules = []*Module{
		{"mod3", "mod3", []string{"mod2"}},
		{"mod5", "mod5", []string{"mod3"}},
		{"mod0", "mod0", []string{}},
		{"mod1", "mod1", []string{"mod0"}},
		{"mod4", "mod4", []string{"mod2", "mod1"}},
		{"mod2", "mod2", []string{"mod1", "mod0"}},
	}

	res, _ := loadDependenciesGraph(modules)

	assert.Equal(t, "6 [(0 5) (1 0) (3 2) (4 3) (4 5) (5 2) (5 3)]", res.graph.String())
	assert.Equal(t, map[string]int{
		"mod0": 2, "mod1": 3, "mod2": 5, "mod3": 0, "mod4": 4, "mod5": 1,
	}, res.moduleIndex)
	assert.Equal(t, map[string][]*Module{
		"mod0": {modules[3], modules[5]},
		"mod1": {modules[4], modules[5]},
		"mod2": {modules[0], modules[4]},
		"mod3": {modules[1]},
	}, res.inverseDeps)
}

func TestLoadDependenciesGraphWithCyrcularDeps(t *testing.T) {

	var modules = []*Module{
		{"mod3", "mod3", []string{"mod2"}},
		{"mod5", "mod5", []string{"mod3"}},
		{"mod0", "mod0", []string{"mod3"}},
		{"mod1", "mod1", []string{"mod0"}},
		{"mod4", "mod4", []string{"mod2", "mod1"}},
		{"mod2", "mod2", []string{"mod1", "mod0"}},
	}

	res, err := loadDependenciesGraph(modules)

	assert.Nil(t, res)
	assert.Equal(t, "ERROR Circular dependency detected", err.Error())
}

func TestExecuteOnAll(t *testing.T) {

	var modules = []*Module{
		{"mod3", "mod3", []string{"mod2"}},
		{"mod5", "mod5", []string{"mod3"}},
		{"mod0", "mod0", []string{}},
		{"mod1", "mod1", []string{"mod0"}},
		{"mod4", "mod4", []string{"mod2", "mod1", "mod3"}},
		{"mod2", "mod2", []string{"mod1", "mod0"}},
	}
	g, _ := loadDependenciesGraph(modules)

	res := []string{}
	g.executeOnAll(func(s string) error {
		res = append(res, s)
		return nil
	})
	assert.Equal(t, "mod0", res[0])
	assert.Equal(t, "mod1", res[1])
	assert.Equal(t, "mod2", res[2])
	assert.Equal(t, "mod3", res[3])
	assert.Equal(t, "mod5", res[4])
	assert.Equal(t, "mod4", res[5])
}

func TestExecuteOnAllStressTest(t *testing.T) {
	if os.Getenv("STRESS_TEST") != "" {
		t.Skip("Skipping stress test. Set STRESS_TEST env variable to run the test.")
	}

	numsParallelDeps := 100000
	modules := generateLargeNumberOfModules(numsParallelDeps)

	g, _ := loadDependenciesGraph(modules)

	res := []string{}
	g.executeOnAll(func(s string) error {
		res = append(res, s)
		return nil
	})
	assert.Equal(t, numsParallelDeps+2, len(res))
}

func TestExecuteOnDependencies(t *testing.T) {
	
	var modules = []*Module{
		{"mod3", "mod3", []string{"mod2"}},
		{"mod5", "mod5", []string{"mod3"}},
		{"mod0", "mod0", []string{}},
		{"mod1", "mod1", []string{"mod0"}},
		{"mod4", "mod4", []string{"mod2", "mod1", "mod3"}},
		{"mod2", "mod2", []string{"mod1", "mod0"}},
	}
	g, _ := loadDependenciesGraph(modules)
	
	res := []string{}
	g.executeOnDependencies("mod3", func(s string) error {
		res = append(res, s)
		return nil
	})
	assert.Equal(t, "mod0", res[0])
	assert.Equal(t, "mod1", res[1])
	assert.Equal(t, "mod2", res[2])
	assert.Equal(t, "mod3", res[3])
}

func generateLargeNumberOfModules(numsParallelDeps int) []*Module {
	var modules = []*Module{
		{"mod0", "mod0", []string{}},
	}
	lastModDeps := []string{}
	for i:=0; i<numsParallelDeps; i++ {
		modName := fmt.Sprintf("mod%d", i+1)
		modules = append(modules, &Module{modName, modName, []string{"mod0"}})
		lastModDeps = append(lastModDeps, modName)
	}
	lastModName := fmt.Sprintf("mod%d", numsParallelDeps+1)
	return append(modules, &Module{lastModName, lastModName, lastModDeps})
}

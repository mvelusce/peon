package project

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
}

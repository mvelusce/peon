package project

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestLoadModules(t *testing.T) {

	res, err := loadModules("../../test/data/project")

	assert.Empty(t, err, "Error must be nil")

	assert.Equal(t, 4, len(res))

	modA := res[0]
	expectedA := &Module{
		Name: "module_a",
		Path: "../../test/data/project/module_a",
	}
	assert.True(t, reflect.DeepEqual(modA, expectedA), "expected %s, got %v", expectedA, modA)

	modB := res[1]
	expectedB := &Module{
		Name:         "module_b",
		Path:         "../../test/data/project/module_b",
		Dependencies: []string{"module_a"},
	}
	assert.True(t, reflect.DeepEqual(modB, expectedB), "expected %v, got %v", expectedB, modB)
}

func TestLoadingModulesYaml(t *testing.T) {
	path := "../../test/data/project"
	res, err := loadYamlModules(path)

	assert.Empty(t, err, "Error must be nil")

	assert.Equal(t, 4, len(res))
	assert.Equal(t, "module_a", res[0].Name)
	assert.Equal(t, "module_b", res[1].Name)
}

func TestLoadingModulesJson(t *testing.T) {
	path := "../../test/data/project"
	res, err := loadJsonModules(path)

	assert.Empty(t, err, "Error must be nil")

	assert.Equal(t, 5, len(res))
	assert.Equal(t, "module_a", res[0].Name)
	assert.Equal(t, "module_b", res[1].Name)
}

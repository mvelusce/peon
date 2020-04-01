package project

import (
	"reflect"
	"testing"
)

func TestLoadModules(t *testing.T) {

	res := loadModules("testdata")

	if len(res) != 4 {
		t.Errorf("expected 4, got %d", len(res))
	}

	modA := res[0]
	expectedA := PyModule{
		Name: "module_a",
		Path: "testdata/module_a",
	}
	if !reflect.DeepEqual(modA, expectedA) {
		t.Errorf("expected %v, got %v", expectedA, modA)
	}

	modB := res[1]
	expectedB := PyModule{
		Name:         "module_b",
		Path:         "testdata/module_b",
		Dependencies: []string{"module_a"},
	}
	if !reflect.DeepEqual(modB, expectedB) {
		t.Errorf("expected %v, got %v", expectedB, modB)
	}
}

func TestLoadingModulesYaml(t *testing.T) {
	path := "testdata"
	res := loadYamlModules(path)

	if len(res) != 4 {
		t.Errorf("In %s expected 4 projects, got %d", path, len(res))
	}
	if res[0].Name != "module_a" {
		t.Errorf("expected module_a, got %s", res[0].Name)
	}
	if res[1].Name != "module_b" {
		t.Errorf("expected module_b, got %s", res[1].Name)
	}
}

func TestLoadingModulesJson(t *testing.T) {
	path := "testdata"
	res := loadJsonModules(path)

	if len(res) != 5 {
		t.Errorf("In %s expected 2 projects, got %d", path, len(res))
	}
	if res[0].Name != "module_a" {
		t.Errorf("expected module_a, got %s", res[0].Name)
	}
	if res[1].Name != "module_b" {
		t.Errorf("expected module_b, got %s", res[1].Name)
	}
}

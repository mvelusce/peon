package project

import (
	"reflect"
	"testing"
)

func TestParsingSetupPy(t *testing.T) {

	modules := []PyModule{
		{Name: "module_a"},
		{Name: "module_b"},
	}

	res, err := parseSetupPyFile("../../test/data/project", modules)

	if err != nil {
		t.Errorf("Parse return error %v", err)
	}
	moduleADeps := res.Dependencies
	if len(moduleADeps) != 1 {
		t.Errorf("expected 1, got %d", len(moduleADeps))
	}

}

func TestParsingSetupPyFiles(t *testing.T) {

	modules := []PyModule{
		{Name: "module_a", Path: "../../test/data/project/module_a"},
		{Name: "module_b", Path: "../../test/data/project/module_b"},
	}

	res := parseSetupPyFiles(modules)

	if len(res) != 2 {
		t.Errorf("expected 2, got %d", len(res))
	}

	modA := res[0]
	expectedA := PyModule{
		Name: "module_a",
		Path: "../../test/data/project/module_a",
	}
	if !reflect.DeepEqual(modA, expectedA) {
		t.Errorf("expected %v, got %v", expectedA, modA)
	}

	modB := res[1]
	expectedB := PyModule{
		Name:         "module_b",
		Path:         "../../test/data/project/module_b",
		Dependencies: []string{"module_a"},
	}
	if !reflect.DeepEqual(modB, expectedB) {
		t.Errorf("expected %v, got %v", expectedB, modB)
	}
}

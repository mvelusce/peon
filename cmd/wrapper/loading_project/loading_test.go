package loading_project

import (
	"testing"
)

func TestLoadingModulesYaml(t *testing.T) {
	path := "testdata/modules.yml"
	res := loadYamlModules(path)

	if len(res) != 2 {
		t.Errorf("In %s expected 2 projects, got %d", path, len(res))
	}
	if res[0].Module != "module_a" {
		t.Errorf("expected module_a, got %s", res[0].Module)
	}
	if res[1].Module != "module_b" {
		t.Errorf("expected module_b, got %s", res[1].Module)
	}
}

func TestLoadingModulesJson(t *testing.T) {
	path := "testdata/modules.json"
	res := loadJsonModules(path)

	if len(res) != 5 {
		t.Errorf("In %s expected 2 projects, got %d", path, len(res))
	}
	if res[0].Module != "module_a" {
		t.Errorf("expected module_a, got %s", res[0].Module)
	}
	if res[1].Module != "module_b" {
		t.Errorf("expected module_b, got %s", res[1].Module)
	}
}

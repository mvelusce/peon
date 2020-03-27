package loading_project

import (
	"testing"
)

func TestParsingSetupPy(t *testing.T) {

	modules := []Modules{
		{Module: "module_a"},
		{Module: "module_b"},
	}

	module, _ := parseSetupPyFile("testdata", modules)
	if module.name != "module_a" {
		t.Errorf("expected name to be module_a, got %s", module.name)
	}

	if len(module.dependencies) != 2 {
		t.Errorf("expected 1, got %d", len(module.dependencies))
	}

}

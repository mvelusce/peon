package loading_project

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type setupPyModule struct {
	name         string
	dependencies []string
}

const setupPy = "setup.py"

func parseSetupPyFile(path string, modules []Modules) (setupPyModule, error) {

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, setupPy))
	if err != nil {
		log.Fatalf("Failed to read setup.py file in %s. Error: %v", path, err)
	}
	content := string(file)

	module := setupPyModule{}

	parseName(content, &module)

	parseDependencies(modules, content, &module)

	return module, nil
}

func parseDependencies(modules []Modules, content string, module *setupPyModule) {
	for _, m := range modules {

		if strings.Contains(content, m.Module) {
			module.dependencies = append(module.dependencies, m.Module)
		}
	}
}

func parseName(content string, module *setupPyModule) {
	r := regexp.MustCompile(`name='(.+)'`)
	nameRes := r.FindStringSubmatch(content)
	if len(nameRes) == 2 {
		module.name = nameRes[1]
	}
}

package project

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"regexp"
	"strings"
)

const setupPy = "setup.py"

func parseSetupPyFiles(modules []Module) []Module {

	var modulesWithDeps []Module
	for _, m := range modules {
		pyMod, err := parseSetupPyFile(m.Path, modules)
		if err == nil {
			modulesWithDeps = append(modulesWithDeps, pyMod)
		} else {
			modulesWithDeps = append(modulesWithDeps, m)
		}
	}
	return modulesWithDeps
}

func parseSetupPyFile(path string, modules []Module) (Module, error) {

	p := TrimSuffix(path, "/")
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", p, setupPy))
	if err != nil {
		log.Errorf("Failed to read setup.py file in %s. Error: %v", p, err)
		return Module{}, err
	}
	content := string(file)

	name := parseName(content)

	deps := parseDependencies(modules, content, name)

	for _, m := range modules {
		if m.Name == name {
			m.Dependencies = deps
			return m, nil
		}
	}
	return Module{}, errors.New("Py module not found in " + p)
}

func parseDependencies(modules []Module, content string, nameToExclude string) []string {

	c := strings.Replace(content, "\n", "", -1)
	c = strings.Replace(c, nameToExclude, "", -1)

	var deps []string
	for _, m := range modules {
		r := regexp.MustCompile(`install_requires=\[.+` + m.Name + `.+\]`)
		if r.MatchString(c) {
			deps = append(deps, m.Name)
		}
	}
	return deps
}

func parseName(content string) string {
	r := regexp.MustCompile(`name='(.+)'`)
	nameRes := r.FindStringSubmatch(content)
	if len(nameRes) == 2 {
		return nameRes[1]
	}
	return ""
}

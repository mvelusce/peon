package loading_project

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Modules struct {
	Module          string `json:"module" yaml:"module"`
	Path            string `json:"path" yaml:"path"`
	ModuleDeps      []string
	ModuleLocalDeps []string
}

type modulesRoot struct {
	Modules []Modules
}

const modulesYaml = "modules.yml"
const modulesJson = "modules.json"

func LoadModules() ([]Modules, error) {

	if _, err := os.Stat(modulesYaml); err == nil {
		return loadYamlModules(modulesYaml), nil
	} else {

		if _, err := os.Stat(modulesJson); err == nil {
			return loadJsonModules(modulesJson), nil
		} else {
			log.Fatalf("Failed to load modules. Error: %v", err)
			return nil, err
		}
	}
}

func loadYamlModules(path string) []Modules { // TODO return error
	var c []Modules
	modules, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read yaml modules. Error: %v", err)
	}
	err = yaml.Unmarshal(modules, &c)
	if err != nil {
		log.Fatalf("Failed to unmarshal yaml modules. Erroro: %v", err)
	}
	return c
}

func loadJsonModules(path string) []Modules { // TODO return error
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read json modules. Error: %v ", err)
	}
	decoder := json.NewDecoder(file)
	rootModules := modulesRoot{}
	err = decoder.Decode(&rootModules)
	if err != nil {
		log.Fatalf("Failed to decode json modules. Error: %v", err)
	}
	return rootModules.Modules
}

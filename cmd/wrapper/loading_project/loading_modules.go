package loading_project

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type PyModule struct {
	Name         string `json:"module" yaml:"module"`
	Path         string `json:"path" yaml:"path"`
	Dependencies []string
}

type modulesRoot struct {
	Modules []PyModule
}

const modulesYaml = "modules.yml"
const modulesJson = "modules.json"

func loadModules(root string) []PyModule {

	var modules []PyModule

	modules = loadYamlModules(root)

	if len(modules) == 0 {
		modules = loadJsonModules(root)
	}

	return parseSetupPyFiles(modules)
}

func loadYamlModules(root string) []PyModule {

	r := TrimSuffix(root, "/")
	var c []PyModule
	modules, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", r, modulesYaml))
	if err != nil {
		log.Fatalf("Failed to read yaml modules. Error: %v", err)
	}
	err = yaml.Unmarshal(modules, &c)
	if err != nil {
		log.Fatalf("Failed to unmarshal yaml modules. Erroro: %v", err)
	}
	return appendRoot(r, c)
}

func loadJsonModules(root string) []PyModule {

	r := TrimSuffix(root, "/")
	file, err := os.Open(fmt.Sprintf("%s/%s", r, modulesJson))
	if err != nil {
		log.Fatalf("Failed to read json modules. Error: %v ", err)
	}
	decoder := json.NewDecoder(file)
	rootModules := modulesRoot{}
	err = decoder.Decode(&rootModules)
	if err != nil {
		log.Fatalf("Failed to decode json modules. Error: %v", err)
	}
	return appendRoot(r, rootModules.Modules)
}

func appendRoot(root string, modules []PyModule) []PyModule {
	r := TrimSuffix(root, "/")

	var mods []PyModule
	for _, m := range modules {
		path := TrimPrefix(TrimPrefix(m.Path, "."), "/")
		p := fmt.Sprintf("%s/%s", r, path)
		m.Path = p
		mods = append(mods, m)
	}
	return mods
}

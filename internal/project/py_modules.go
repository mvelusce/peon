package project

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

const modulesYaml = "peon-modules.yml" // TODO move this in project and allow a custom file name
const modulesJson = "peon-modules.json"

func loadModules(root string) ([]PyModule, error) {

	var modules []PyModule

	modules, err := loadYamlModules(root)

	if err != nil || len(modules) == 0 {
		modules, err = loadJsonModules(root)
	}

	return parseSetupPyFiles(modules), err
}

func loadYamlModules(root string) ([]PyModule, error) {

	r := TrimSuffix(root, "/")
	var c []PyModule
	modules, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", r, modulesYaml))
	if err != nil {
		log.Printf("Failed to read yaml modules. Error: %v", err)
		return nil, err
	}
	err = yaml.Unmarshal(modules, &c)
	if err != nil {
		log.Printf("Failed to unmarshal yaml modules. Erroro: %v", err)
		return nil, err
	}
	return appendRoot(r, c), nil
}

func loadJsonModules(root string) ([]PyModule, error) {

	r := TrimSuffix(root, "/")
	file, err := os.Open(fmt.Sprintf("%s/%s", r, modulesJson))
	if err != nil {
		log.Printf("Failed to read json modules. Error: %v ", err)
		return nil, err
	}
	decoder := json.NewDecoder(file)
	rootModules := modulesRoot{}
	err = decoder.Decode(&rootModules)
	if err != nil {
		log.Printf("Failed to decode json modules. Error: %v", err)
		return nil, err
	}
	return appendRoot(r, rootModules.Modules), nil
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

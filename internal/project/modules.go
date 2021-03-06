package project

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Module struct {
	Name         string `json:"module" yaml:"module"`
	Path         string `json:"path" yaml:"path"`
	Dependencies []string
}

type modulesRoot struct {
	Modules []Module
}

const modulesYaml = "peon-modules.yml"
const modulesJson = "peon-modules.json"

func loadModules(root string) ([]Module, error) {

	var modules []Module

	modules, err := loadYamlModules(root)

	if err != nil || len(modules) == 0 {
		log.Infof("Unable to read yaml modules. Falling back to json.")
		modules, err = loadJsonModules(root)
	}

	return parseSetupPyFiles(modules), err
}

func loadYamlModules(root string) ([]Module, error) {

	r := TrimSuffix(root, "/")
	var c []Module
	modules, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", r, modulesYaml))
	if err != nil {
		log.Debugf("Failed to read yaml modules. Error: %v", err)
		return nil, err
	}
	err = yaml.Unmarshal(modules, &c)
	if err != nil {
		log.Debugf("Failed to unmarshal yaml modules. Error: %v", err)
		return nil, err
	}
	return appendRoot(r, c), nil
}

func loadJsonModules(root string) ([]Module, error) {

	r := TrimSuffix(root, "/")
	file, err := os.Open(fmt.Sprintf("%s/%s", r, modulesJson))
	if err != nil {
		log.Infof("Failed to read json modules. Error: %v ", err)
		return nil, err
	}
	decoder := json.NewDecoder(file)
	rootModules := modulesRoot{}
	err = decoder.Decode(&rootModules)
	if err != nil {
		log.Infof("Failed to decode json modules. Error: %v", err)
		return nil, err
	}
	return appendRoot(r, rootModules.Modules), nil
}

func appendRoot(root string, modules []Module) []Module {
	r := TrimSuffix(root, "/")

	var mods []Module
	for _, m := range modules {
		path := TrimPrefix(TrimPrefix(m.Path, "."), "/")
		p := fmt.Sprintf("%s/%s", r, path)
		m.Path = p
		mods = append(mods, m)
	}
	return mods
}

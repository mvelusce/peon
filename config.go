package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ModulesY struct {
	Module          string `json:"module" yaml:"module"`
	Path            string `json:"path" yaml:"path"`
	ModuleDeps      []string
	ModuleLocalDeps []string
}

func GetModules() []ModulesY {

	var c []ModulesY
	modules, err := ioutil.ReadFile("example/modules.yml")
	if err != nil {
		log.Printf("modules.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(modules, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

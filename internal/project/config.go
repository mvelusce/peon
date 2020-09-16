package project

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type Config struct {
	ProjectRoot string `json:"project-root" yaml:"project-root"`
	PythonExec  string `json:"python-exec" yaml:"python-exec"`
}

const configFile = "./.peon-config.json"

const defaultProjectRoot = "."
const defaultPythonExec = "python3.7"

func NewConfig(projectRoot string, pythonExec string) *Config {

	config := loadConfig(configFile, projectRoot, pythonExec)

	_ = writeConfig(configFile, config)

	return config
}

func loadConfig(configFile string, projectRoot string, pythonExec string) *Config {

	config, err := loadConfigFile(configFile)
	if err != nil {
		config = &Config{
			ProjectRoot: defaultProjectRoot,
			PythonExec:  defaultPythonExec,
		}
	}
	if projectRoot != "" {
		config.ProjectRoot = projectRoot
	}
	if pythonExec != "" {
		config.PythonExec = pythonExec
	}

	return config
}

func loadConfigFile(configFile string) (*Config, error) {

	file, err := os.Open(configFile)
	if err != nil {
		log.Debugf("Failed to read configs. Error: %v ", err)
		return nil, err
	}
	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		log.Debugf("Failed to decode json configs. Error: %v", err)
		return nil, err
	}
	return config, nil
}

func writeConfig(configFile string, config *Config) error {
	file, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		log.Debugf("Failed to marshal configs. Error: %v", err)
		return err
	}
	err = ioutil.WriteFile(configFile, file, 0644)
	if err != nil {
		log.Debugf("Failed to write file. Error: %v", err)
	}
	return err
}

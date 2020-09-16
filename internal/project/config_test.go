package project

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfigFromFile(t *testing.T) {

	res := loadConfig("../../test/data/project/.test-peon-config.json", "", "")

	assert.Equal(t, "test-root", res.ProjectRoot)
	assert.Equal(t, "test-python-exec", res.PythonExec)
}

func TestLoadConfigFromArgs(t *testing.T) {

	res := loadConfig("../../test/data/project/.test-peon-config.json", "args1", "args2")

	assert.Equal(t, "args1", res.ProjectRoot)
	assert.Equal(t, "args2", res.PythonExec)
}

func TestLoadConfigBothFromArgsAndFile(t *testing.T) {

	res := loadConfig("../../test/data/project/.test-peon-config.json", "args1", "")

	assert.Equal(t, "args1", res.ProjectRoot)
	assert.Equal(t, "test-python-exec", res.PythonExec)
}

func TestWriteConfigFile(t *testing.T) {

	config := &Config{
		ProjectRoot: "asd",
		PythonExec:  "qwe",
	}
	path := "../../test/data/project/write-test.json"
	err := writeConfig(path, config)

	assert.Empty(t, err, "Error must be nil")

	res := loadConfig(path, "", "")
	assert.Equal(t, "asd", res.ProjectRoot)
	assert.Equal(t, "qwe", res.PythonExec)

	err = os.Remove(path)
	assert.Empty(t, err, "Error must be nil")
}

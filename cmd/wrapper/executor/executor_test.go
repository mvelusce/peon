package executor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecutorBuild(t *testing.T) {

	e := &SetupPyExecutor{PyVersion: "python3.7"}

	err := e.Build("testdata/module_a")

	assert.Empty(t, err, "Error must be nil")
}

func TestExecutorClean(t *testing.T) {

	e := &SetupPyExecutor{PyVersion: "python3.7"}

	err := e.Clean()

	assert.Empty(t, err, "Error must be nil")
}

func TestExecutorTest(t *testing.T) {

	e := &SetupPyExecutor{PyVersion: "python3.7"}

	err := e.Test("testdata/module_a")

	assert.Empty(t, err, "Error must be nil")
}

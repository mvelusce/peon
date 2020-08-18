package executor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecutorBuild(t *testing.T) {

	e := &SetupPyExecutor{PyVersion: "python3"}

	err := e.Build("../../test/data/executor/module_a")

	assert.Empty(t, err, "Error must be nil")
}

func TestExecutorClean(t *testing.T) {

	e := &SetupPyExecutor{PyVersion: "python3"}

	err := e.Clean()

	assert.Empty(t, err, "Error must be nil")
}

func TestExecutorTest(t *testing.T) {

	e := &SetupPyExecutor{PyVersion: "python3"}

	err := e.Test("../../test/data/executor/module_a")

	assert.Empty(t, err, "Error must be nil")
}

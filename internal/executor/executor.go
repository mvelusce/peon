package executor

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

type Executor interface {
	Build(path string) error
	Clean() error
	Test(path string) error
	Exec(command string, path string) error
}

type SetupPyExecutor struct {
	PyVersion string
}

const venv = "peonvenv"

func (e *SetupPyExecutor) Build(path string) error {
	log.Infof("Building %s", path)
	return e.runInVenv("python setup.py install", path)
}

func (e *SetupPyExecutor) Clean() error {
	log.Infof("Cleaning %s", venv)
	return runCommand("rm", "-r", venv)
}

func (e *SetupPyExecutor) Test(path string) error {
	log.Infof("Testing %s", path)
	return e.runInVenv("python setup.py test", path)
}

func (e *SetupPyExecutor) Exec(command string, path string) error {
	log.Infof("Executing command %s in %s", command, path)
	return e.runInVenv(command, path)
}

func (e *SetupPyExecutor) createVenv() error {
	command := fmt.Sprintf("%s", e.PyVersion)
	return runCommand(command, "-m", "venv", venv)
}

func (e *SetupPyExecutor) runInVenv(command string, path string) error {
	err := e.createVenv()
	if err == nil {
		activateVenv := fmt.Sprintf("source %s/bin/activate", venv)
		cd := fmt.Sprintf("cd %s", path)
		command := fmt.Sprintf("%s; %s; %s", activateVenv, cd, command)
		return runCommand("bash", "-c", command)
	}
	log.Errorf("Unable to init virtual env. Error: %v", err)
	return err
}

func runCommand(command string, arg ...string) error {
	return runCommandInPath(".", command, arg...)
}

func runCommandInPath(path string, command string, arg ...string) error {
	cmd := exec.Command(command, arg...)
	cmd.Dir = path

	stdout, err := cmd.CombinedOutput()

	log.Infof("Command output: \n%s\n", string(stdout))

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

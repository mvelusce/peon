package executor

import (
	"fmt"
	"log"
	"os/exec"
)

type PyExecutor interface {
	Build(path string) error
	Run(path string) error
	Clean() error
	Test(path string) error
}

type SetupPyExecutor struct {
	PyVersion string
}

func (e *SetupPyExecutor) Build(path string) error {
	log.Printf("Building %s", path)
	err := e.createVenv()
	if err == nil {
		activateVenv := "source venv/bin/activate"
		cd := fmt.Sprintf("cd %s", path)
		install := "python setup.py install"
		command := fmt.Sprintf("%s; %s; %s", activateVenv, cd, install)
		return runCommand("bash", "-c", command)
	}
	log.Fatalf("Unable to init project. Error: %v", err)
	return err
}

func (e *SetupPyExecutor) Run(path string) error {
	// TODO does it make sense ??
	return nil
}

func (e *SetupPyExecutor) Clean() error {
	return runCommand("rm", "-r", "venv")
}

func (e *SetupPyExecutor) Test(path string) error {
	log.Printf("Testing %s", path)
	err := e.createVenv()
	if err == nil {
		activateVenv := "source venv/bin/activate"
		cd := fmt.Sprintf("cd %s", path)
		install := "python setup.py test"
		command := fmt.Sprintf("%s; %s; %s", activateVenv, cd, install)
		return runCommand("bash", "-c", command)
	}
	log.Fatalf("Unable to init project. Error: %v", err)
	return err
}

func (e *SetupPyExecutor) createVenv() error {
	command := fmt.Sprintf("%s", e.PyVersion)
	return runCommand(command, "-m", "venv", "venv")
}

func runCommand(command string, arg ...string) error {
	return runCommandInPath(".", command, arg...)
}

func runCommandInPath(path string, command string, arg ...string) error {
	cmd := exec.Command(command, arg...)
	cmd.Dir = path

	stdout, err := cmd.CombinedOutput()

	log.Println(string(stdout))

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

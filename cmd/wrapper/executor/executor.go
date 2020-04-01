package executor

import "fmt"

type PyExecutor interface {
	Build(path string) error
	Clean() error
	Test(path string) error
}

type SetupPyExecutor struct{}

func (e *SetupPyExecutor) Build(path string) error {
	fmt.Println("BUILD")
	return nil
}

func (e *SetupPyExecutor) Clean() error {
	return nil
}

func (e *SetupPyExecutor) Test(path string) error {
	return nil
}

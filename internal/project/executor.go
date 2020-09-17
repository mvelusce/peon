package project

import "github.com/sirupsen/logrus"

type DryRunExecutor struct{}

func (e *DryRunExecutor) Build(path string) error {
	logrus.Infof("Building %s", path)
	return nil
}

func (e *DryRunExecutor) Clean() error {
	logrus.Infof("Cleaning")
	return nil
}
func (e *DryRunExecutor) Test(path string) error {
	logrus.Infof("Testing %s", path)
	return nil
}

func (e *DryRunExecutor) Exec(command string, path string) error {
	logrus.Infof("Executing command %s in %s", command, path)
	return nil
}

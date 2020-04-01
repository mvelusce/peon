package main

import (
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	&commandBuild,
}

var commandBuild = cli.Command{
	Name:        "build",
	Usage:       "",
	Description: `Prints modified files.`,
	Action:      build,
}

func build(c *cli.Context) error {

	// TODO build all according to graph
	return runCommand("python setup.py install")
}

func activateEnv() error {
	return runCommand("source venv/bin/activate") // TODO check if already active
}

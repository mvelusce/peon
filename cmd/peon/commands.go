package main

import (
	"github.com/skyveluscekm/setuptools.wrapper/internal/project"
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	&commandBuild,
}

var commandBuild = cli.Command{
	Name:        "build",
	Aliases:     []string{"b"},
	Usage:       "Build all modules",
	Description: "Build all modules",
	Action:      build,
}

func build(c *cli.Context) error {

	projectRoot := c.String("project-root")
	pythonVersion := c.String("py-version")
	// TODO not supported parameter
	c.String("modules-file")

	println(pythonVersion)
	mod := c.Args().First()
	println(mod)

	p, err := project.LoadProject(projectRoot, pythonVersion)
	if err != nil {
		return err
	}
	if mod == "" {
		return p.Build()
	} else {
		return p.BuildModule(mod)
	}
}

func release(c *cli.Context) error {

	// TODO release all python modules
	return nil
}

func activateEnv() error {
	return runCommand("source venv/bin/activate") // TODO check if already active
}

package main

import (
	"github.com/skyveluscekm/peon/internal/project"
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	&commandBuild,
	&testBuild,
	&cleanBuild,
}

var commandBuild = cli.Command{
	Name:        "build",
	Aliases:     []string{"b"},
	Usage:       "Build all modules. Pass module name to build the module (and it's dependencies).",
	Description: "Build all modules",
	Action:      buildCommand,
}

var testBuild = cli.Command{
	Name:        "test",
	Aliases:     []string{"t"},
	Usage:       "Test all modules. Pass module name to test the module (and it's dependencies).",
	Description: "Test all modules",
	Action:      testCommand,
}

var cleanBuild = cli.Command{
	Name:        "clean",
	Aliases:     []string{"c"},
	Usage:       "Clean all modules by deleting the virtual env.",
	Description: "Clean all modules",
	Action:      cleanCommand,
}

func buildCommand(c *cli.Context) error {

	return executeCommand(c, build)
}

func build(p *project.Project, module string) error {
	if module == "" {
		return p.Build()
	} else {
		return p.BuildModule(module)
	}
}

func testCommand(c *cli.Context) error {

	return executeCommand(c, test)
}

func test(p *project.Project, module string) error {
	if module == "" {
		return p.Test()
	} else {
		return p.TestModule(module)
	}
}

func cleanCommand(c *cli.Context) error {
	p, err := loadProject(c)
	if err != nil {
		return err
	}
	return p.Clean()
}

func release(c *cli.Context) error {

	// TODO release all python modules
	return nil
}

func executeCommand(c *cli.Context, runCommand func(*project.Project, string) error) error {

	module := c.Args().First()

	p, err := loadProject(c)
	if err != nil {
		return err
	}
	return runCommand(&p, module)
}

func loadProject(c *cli.Context) (project.Project, error) {

	projectRoot := c.String("project-root")
	pythonVersion := c.String("py-version")
	// TODO not supported parameter
	c.String("modules-file")

	p, err := project.LoadProject(projectRoot, pythonVersion)

	return p, err
}

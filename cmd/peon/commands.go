package main

import (
	"errors"
	"github.com/skyveluscekm/peon/internal/project"
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	&commandBuild,
	&commandTest,
	&commandClean,
	&commandExec,
}

var commandBuild = cli.Command{
	Name:        "build",
	Aliases:     []string{"b"},
	Usage:       "Build all modules. Pass module name to build the module (and it's dependencies).",
	Description: "Build all modules",
	Action:      buildCommand,
}

var commandTest = cli.Command{
	Name:        "test",
	Aliases:     []string{"t"},
	Usage:       "Test all modules. Pass module name to test the module (and it's dependencies).",
	Description: "Test all modules",
	Action:      testCommand,
}

var commandClean = cli.Command{
	Name:        "clean",
	Aliases:     []string{"c"},
	Usage:       "Clean all modules by deleting the virtual env.",
	Description: "Clean all modules",
	Action:      cleanCommand,
}

var commandExec = cli.Command{
	Name:        "exec",
	Aliases:     []string{"e"},
	Usage:       "Exec command on modules. Examples:\n\t Run command on all modules: peon exec 'my custom command'\n\t Run command on a module and its dependencies: peon exec my-module 'my custom command'",
	Description: "Exec command on modules",
	Action:      execCommand,
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

func execCommand(c *cli.Context) error {

	module := ""
	command := ""
	if c.Args().Len() > 1 {
		module = c.Args().First()
		command = c.Args().Get(1)
	} else {
		command = c.Args().First()
	}

	p, err := loadProject(c)
	if err != nil {
		return err
	}

	return exec(&p, module, command)
}

func exec(p *project.Project, module string, command string) error {
	if command == "" {
		return errors.New("empty command")
	}
	if module == "" {
		return p.Exec(command)
	} else {
		return p.ExecModule(command, module)
	}
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

	config := project.NewConfig(projectRoot, pythonVersion)

	p, err := project.LoadProject(config)

	return p, err
}

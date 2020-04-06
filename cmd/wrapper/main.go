package main

import (
	"fmt"
	"github.com/skyveluscekm/setuptools.wrapper/cmd/wrapper/executor"
	"github.com/skyveluscekm/setuptools.wrapper/cmd/wrapper/project"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {

	e := &executor.SetupPyExecutor{PyVersion: "python3.7"}

	asd := e.Build("testdata/module_a")
	fmt.Println(asd)

	p := project.LoadProject()

	p.Build()

	app := cli.NewApp()
	app.Name = "gip"
	app.Version = "0.0.0"
	app.Usage = "Manage your setup.py ptojects"
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "file", Aliases: []string{"f"}, Usage: ""},
		&cli.BoolFlag{Name: "debug", Aliases: []string{"d"}, Usage: ""},
		&cli.BoolFlag{Name: "quiet", Aliases: []string{"q"}, Usage: ""},
		&cli.BoolFlag{Name: "ignore-missing", Aliases: []string{"m"}, Usage: ""},
	}
	app.EnableBashCompletion = true

	//app.Commands = commands

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// read modules root
	// read single modules dependencies and add it to list of modules <-- READ SETUP.PY
	// create graph of dependencies

	// build module == add folder to python path
	// or build with setup.py --> need to parse setup.py

	// start commands
	// init create venv
	// build project: create venv, pip install dependencies by module in graph order
	// clean: delete venv
	// build single module: create venv, pip install dependecies of module starting from module in graph

	// run tests: run python -m unittest with all files test_*.py
	// run module: python bin/run_module_name.py

	cmd := exec.Command("ls")
	cmd.Dir = "cmd/wrapper/loading_project/testdata"
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}

package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "peon"
	app.Version = "0.0.0"
	app.Usage = "One script to rule them all!"
	app.Flags = []cli.Flag{
		//&cli.StringFlag{Name: "modules-file", Aliases: []string{"f"}, Usage: ""},
		&cli.StringFlag{Name: "project-root", Aliases: []string{"r"}, Usage: ""},
		&cli.StringFlag{Name: "py-version", Aliases: []string{"p"}, Usage: ""},
	}
	app.EnableBashCompletion = true

	app.Commands = commands

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
}

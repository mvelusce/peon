package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {

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

	app := &cli.App{
		Action: func(c *cli.Context) error {
			fmt.Printf("Hello %q", c.Args().Get(0))
			return nil
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package wrapper

import (
	"fmt"
	"github.com/skyveluscekm/setuptools.wrapper/cmd/wrapper/loading_project"
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

	var modules = loading_project.LoadModules("")
	fmt.Println(modules[1].Name)
}

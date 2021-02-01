# peon

Work work...

Peon is a simple script to manage multiple python applications and libraries in a monorepo. Peon uses Setuptools to build and test python applications, modules and libraries.

If you have many python modules in several repositories with complex intra-dependencies, you can use Peon to move all your code to a monorepo and simplify your development process.

## Why

There are many other tools that aim to help manage monorepos containing python code. We found it difficult to adapt them to our existing code base. Peon uses Setuptools configuration to discover python modules and their dependencies. If you already use Setuptools to manage your python code, you can use Peon to build and test multiple modules and libraries with one command. There is no need to migrate your code base to another build system.

## Usage

Create a `peon-modules.yml` or `peon-modules.josn` file listing all python modules you want to manage with Peon.

Build all modules with:
```bash
./peon build
```

Build a single module and all its dependencies:
```bash
./peon build [module_name]
```

Test all modules or a single module (and its dependencies) with:
```bash
./peon test [module_name]
```

You can select the project root and the python version as follows:
```bash
./peon --project-root ../../test/data/project --py-version python3.5 build
```
Peon saves the last valid parameters in the file `.peon-config.json` so you don't have to type them every time.

For a list of all possible parameters:
```bash
./peon --help
```

You can execute any command with `exec`. Peon will execute the command on all modules and libraries, maintaining dependencies.
```bash
./peon exec module_d 'python setup.py test'
```

See `test/data` for examples.

## Development

The code is written in go 1.15.

Init Go module
```bash
go mod init
```

Get all dependencies
```bash
go get -d ./...
```

Build program executable
```bash
cd cmd/peon
go build
```

Run all tests
```bash
go test ./...
```

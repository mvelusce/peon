package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Name = "peon"
	app.Version = "0.0.0"
	app.Usage = "Something needs doing?"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{Name: "dry-run", Aliases: []string{"d"}, Usage: ""},
		&cli.StringFlag{Name: "project-root", Aliases: []string{"r"}, Usage: ""},
		&cli.StringFlag{Name: "py-version", Aliases: []string{"p"}, Usage: ""},
	}
	app.EnableBashCompletion = true

	app.Commands = commands

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

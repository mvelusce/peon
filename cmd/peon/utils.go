package main

import (
	"fmt"
	"os/exec"
)

func runCommand(command string) error {
	cmd := exec.Command(command)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Print(string(stdout))
	return nil
}

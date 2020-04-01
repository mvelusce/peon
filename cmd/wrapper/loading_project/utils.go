package loading_project

import (
	"fmt"
	"os/exec"
	"strings"
)

func TrimSuffix(s string, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func TrimPrefix(s string, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		s = s[0+len(prefix):]
	}
	return s
}

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

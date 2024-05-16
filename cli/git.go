package main

import (
	"os/exec"
)

func runGitStatus() (string, error) {
	cmd := exec.Command("git", "status")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(output), nil
}

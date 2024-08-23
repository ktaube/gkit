package main

import (
	"os/exec"
	"strings"
)

// GetDiff executes the git diff command and returns the output as a string
func GetDiff() (string, error) {
	cmd := exec.Command("git", "diff")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

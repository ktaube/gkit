package main

import (
	"os/exec"
	"time"
)

type git struct {
	status   string
	statusTs time.Time
}

func (g *git) runGitStatus() (string, error) {
	cmd := exec.Command("git", "status")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	g.status = string(output)
	g.statusTs = time.Now()

	return g.status, nil
}

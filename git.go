package main

import (
	"os/exec"
	"time"
)

type git struct {
	dirpath  string
	status   string
	statusTs time.Time
}

func (g *git) runGitStatus() (string, error) {
	cmd := exec.Command("git", "-C", g.dirpath, "status")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	g.status = string(output)
	g.statusTs = time.Now()

	return g.status, nil
}

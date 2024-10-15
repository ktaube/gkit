package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	ctx := &ctx{
		theme: getTheme(),
	}

	diff, err := GetDiff()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(model{ctx: ctx, diff: diff, commitMessage: ""}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}


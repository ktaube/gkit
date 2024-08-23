package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	diff, err := GetDiff()
	if err != nil {
		log.Fatal(err)
	}

	ctx := &ctx{
		theme: getTheme(),
		diff:  diff,
	}

	p := tea.NewProgram(model{ctx: ctx}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

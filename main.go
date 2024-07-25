package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	ctx := &ctx{
		theme: getTheme(),
	}

	p := tea.NewProgram(model{ctx: ctx}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

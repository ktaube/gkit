package main

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model int

type tickMsg string

func main() {
	p := tea.NewProgram(model(5), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		status, err := runGitStatus()
		if err != nil {
			log.Fatal(err)
		}
		return tickMsg(status)
	})
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg:
		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprint("\n\n     Hi. This this is git status\n", m)
}

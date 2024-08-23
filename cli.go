package main

import (
	"log"
	"os"
	"time"

	tree "github.com/charmbracelet/lipgloss/tree"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

type model struct {
	ctx *ctx
}
type tickMsg string

func (m model) tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg("")
	})
}

func (m model) Init() tea.Cmd {
	return m.tick()
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg:
		return m, m.tick()
	}

	return m, nil
}

func (m model) View() string {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	style := getRootStyle(width, height, m.ctx.theme)

	t := tree.New().
		EnumeratorStyle(enumeratorStyle).
		Root("# Table of Contents").
		Child(
			tree.New().
				Root("## Chapter 1").
				Child("Chapter 1.1").
				Child("Chapter 1.2"),
		).
		Child(
			tree.New().
				Root("## Chapter 2").
				Child("Chapter 2.1").
				Child("Chapter 2.2"),
		)

	return style.Render(t.String())
}

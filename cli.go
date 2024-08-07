package main

import (
	"log"
	"os"
	"time"

	tree "github.com/charmbracelet/lipgloss/tree"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

	var style = lipgloss.NewStyle().
		Width(width).                                            // Set the width to terminal width
		Height(height).                                          // Set the height to terminal height
		Background(lipgloss.Color(m.ctx.theme.backgroundColor)). // Set your desired background color
		Padding(1)

	// Prepare the content with "hello" at the bottom
	// content := fmt.Sprintf("%s\n%s", m.git.status, m.git.statusTs.Format(time.ANSIC))

	enumeratorStyle := lipgloss.NewStyle().
		Padding(0, 1)

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

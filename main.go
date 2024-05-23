package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type model struct {
	git *git
}
type tickMsg string

func main() {
	dirpath := flag.String("dir", ".", "directory to watch")
	flag.Parse()

	git := &git{dirpath: *dirpath}

	_, err := git.runGitStatus()
	if err != nil {
		log.Fatal(err)
	}
	p := tea.NewProgram(model{git: git}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m model) tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		_, err := m.git.runGitStatus()
		if err != nil {
			log.Fatal(err)
		}
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
	var style = lg.NewStyle().
		Border(lg.RoundedBorder()).
		PaddingTop(2).
		PaddingLeft(4).
		Height(40).
		Width(120)

	return style.Render(fmt.Sprint("", m.git.status, "\n", m.git.statusTs.Format(time.ANSIC)))
}

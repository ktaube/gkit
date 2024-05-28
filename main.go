package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
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
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	var style = lipgloss.NewStyle().
		Width(width).                          // Set the width to terminal width
		Height(height).                        // Set the height to terminal height
		Background(lipgloss.Color("#000000")). // Set your desired background color
		Padding(1)

	// Prepare the content with "hello" at the bottom
	content := fmt.Sprintf("%s\n%s", m.git.status, m.git.statusTs.Format(time.ANSIC))

	return style.Render(content)
}

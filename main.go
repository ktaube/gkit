package main

import (
	"flag"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	dirpath := flag.String("dir", ".", "directory to watch")
	flag.Parse()

	git := &git{dirpath: *dirpath}

	_, err := git.runGitStatus()
	if err != nil {
		log.Fatal(err)
	}

	ctx := &ctx{
		theme: getTheme(),
	}

	p := tea.NewProgram(model{git: git, ctx: ctx}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

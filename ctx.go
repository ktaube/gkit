package main

import "github.com/charmbracelet/lipgloss"

type theme struct {
	backgroundColor lipgloss.Color
}

type ctx struct {
	theme theme
	diff  string
}

func getTheme() theme {
	return theme{
		backgroundColor: lipgloss.Color("#000000"),
	}
}

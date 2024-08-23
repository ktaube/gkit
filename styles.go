package main

import (
	"github.com/charmbracelet/lipgloss"
)

var enumeratorStyle = lipgloss.NewStyle().
	Padding(0, 1)

func getRootStyle(width, height int, theme theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Background(lipgloss.Color(theme.backgroundColor)).
		Padding(1)
}

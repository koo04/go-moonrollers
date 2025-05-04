package main

import "github.com/charmbracelet/lipgloss"

func stylePool() lipgloss.Style {
	return lipgloss.NewStyle().
		Width(36).
		Height(10).
		Align(lipgloss.Center)
}

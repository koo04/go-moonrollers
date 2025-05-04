package main

import "github.com/charmbracelet/lipgloss"

func stylePool(w, h int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(w).
		Height(h).
		Align(lipgloss.Center)
}

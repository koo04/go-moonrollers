package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(newMenuModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
}

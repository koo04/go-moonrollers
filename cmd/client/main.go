package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	debug = false
)

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.Parse()
}

func main() {
	p := tea.NewProgram(newMenuModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
}

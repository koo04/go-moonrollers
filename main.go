package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	board *board

	height int
	width  int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	return m, nil
}

func (m model) View() string {
	// style := lipgloss.NewStyle().Width(m.width).Height(m.height).Align(lipgloss.Center)

	// cardsPane := stylePool().Render(m.board.Render())

	// panes := lipgloss.JoinVertical(
	// 	lipgloss.Center,
	// 	cardsPane,
	// )
	return m.board.Render()
}

func main() {
	b := &board{
		Deck: deck,
		Out:  []*card{},
	}

	b.ShuffleDeck()
	b.DrawCards(5)

	if _, err := tea.NewProgram(model{
		board: b,
	}, tea.WithAltScreen()).Run(); err != nil {
		panic(err)
	}
}

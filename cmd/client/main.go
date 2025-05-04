package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	mr "github.com/koo04/go-moonrollers"
)

type model struct {
	board *mr.Board

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
	style := lipgloss.NewStyle().Width(m.width).Height(m.height).Align(lipgloss.Center)

	cardsPane := style.Render(m.board.Render())

	panes := lipgloss.JoinVertical(
		lipgloss.Center,
		cardsPane,
	)

	return panes
}

func main() {
	b := &mr.Board{
		Deck: mr.GetDeck(),
		Out:  []*mr.Crew{},
	}

	b.ShuffleDeck()
	b.DrawCards(5)

	if _, err := tea.NewProgram(model{
		board: b,
	}, tea.WithAltScreen()).Run(); err != nil {
		panic(err)
	}
}

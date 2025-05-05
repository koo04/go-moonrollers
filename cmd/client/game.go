package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	mr "github.com/koo04/go-moonrollers"
)

type gameModel struct {
	game *mr.Game

	height int
	width  int
}

func (m gameModel) Init() tea.Cmd {
	return nil
}

func (m gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyRunes:
			switch string(msg.Runes) {
			case "q":
				return m, tea.Quit
			}
		case tea.KeyUp, tea.KeyDown, tea.KeyLeft, tea.KeyRight:

		default:
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	return m, nil
}

func (m gameModel) View() string {
	cardsBoardStyle := lipgloss.NewStyle().Align(lipgloss.Left)
	pointsBoardStyle := lipgloss.NewStyle().Width(15).Align(lipgloss.Left).AlignVertical(lipgloss.Center).Border(lipgloss.NormalBorder()).Margin(1)

	cardsPane := cardsBoardStyle.Render(m.game.RenderCards())
	pointsBoard := pointsBoardStyle.Render(m.game.RenderPoints())

	v := lipgloss.JoinHorizontal(
		lipgloss.Left,
		pointsBoard,
		cardsPane,
	)
	// tempDice := style.Align(lipgloss.Left).Render(lipgloss.JoinHorizontal(
	// 	lipgloss.Top,
	// 	mr.Die["damage"],
	// 	mr.Die["thruster"],
	// 	mr.Die["shield"],
	// 	mr.Die["add"],
	// 	mr.Die["wild"],
	// 	mr.Die["reactor"],
	// 	mr.Die["damage"],
	// 	mr.Die["thruster"],
	// 	mr.Die["shield"],
	// 	mr.Die["add"],
	// 	mr.Die["wild"],
	// 	mr.Die["reactor"],
	// ))

	panes := lipgloss.JoinVertical(
		lipgloss.Center,
		v,
		// tempDice,
	)

	return panes
}

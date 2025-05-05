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
		return m, nil
	}

	return m, nil
}

func (m gameModel) View() string {
	cardsPaneStyle := lipgloss.NewStyle().Align(lipgloss.Left).Margin(1)
	pointsPaneStyle := lipgloss.NewStyle().Width(15).Align(lipgloss.Left).AlignVertical(lipgloss.Center).Border(lipgloss.NormalBorder()).Margin(1)
	dicePaneStyle := lipgloss.NewStyle().Align(lipgloss.Left).Border(lipgloss.NormalBorder()).Margin(1)

	if debug {
		cardsPaneStyle = cardsPaneStyle.Border(lipgloss.ASCIIBorder()).BorderForeground(lipgloss.Color("205")).Margin(0)
		pointsPaneStyle = pointsPaneStyle.Border(lipgloss.ASCIIBorder()).BorderForeground(lipgloss.Color("205")).Margin(0)
		dicePaneStyle = dicePaneStyle.Border(lipgloss.ASCIIBorder()).BorderForeground(lipgloss.Color("205")).Margin(0)
	}

	cardsPane := cardsPaneStyle.Render(m.game.RenderCards())
	pointsPane := pointsPaneStyle.Render(m.game.RenderPoints())
	dicePane := dicePaneStyle.Render(m.game.RenderDice())

	v := lipgloss.JoinHorizontal(
		lipgloss.Left,
		pointsPane,
		cardsPane,
		dicePane,
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

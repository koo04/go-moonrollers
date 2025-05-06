package main

import (
	"fmt"
	"math/rand"

	mr "github.com/ascii-arcade/moonrollers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	dicePaneStyle := lipgloss.NewStyle().Align(lipgloss.Left).Border(lipgloss.NormalBorder()).Margin(1)

	fullPane := lipgloss.NewStyle().Width(m.width).Height(m.height).Align(lipgloss.Center).Margin(1)

	if debug {
		cardsPaneStyle = cardsPaneStyle.Border(lipgloss.ASCIIBorder()).BorderForeground(lipgloss.Color(randomColorHash())).Margin(0)
		dicePaneStyle = dicePaneStyle.Border(lipgloss.ASCIIBorder()).BorderForeground(lipgloss.Color(randomColorHash())).Margin(0)

		fullPane = fullPane.Border(lipgloss.ASCIIBorder()).BorderForeground(lipgloss.Color(randomColorHash())).Margin(0)
	}

	cardsPane := cardsPaneStyle.Render(m.game.RenderCards())
	pointsPane := m.game.RenderPoints(debug)
	dicePane := dicePaneStyle.Render(m.game.RenderDice())

	boardPanes := lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			pointsPane,
			cardsPane,
			dicePane,
		),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		fullPane.Render(boardPanes),
		lipgloss.NewStyle().Width(m.width).AlignHorizontal(lipgloss.Left).Foreground(lipgloss.Color("#999999")).Render("Press q to quit"),
	)
}

func randomColorHash() string {
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)

	return fmt.Sprintf("#%02X%02X%02X", r, g, b) // Format as hex color
}

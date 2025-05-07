package main

import (
	"fmt"
	"math/rand"
	"strings"

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
	if m.height < 23 || m.width < 102 {
		return "Window too small, please resize to something larger."
	}

	cardsPaneStyle := lipgloss.NewStyle().Align(lipgloss.Left).Margin(1)
	dicePaneStyle := lipgloss.NewStyle().Align(lipgloss.Left).Border(lipgloss.NormalBorder()).Margin(1)

	if debug {
		cardsPaneStyle = cardsPaneStyle.Border(lipgloss.ASCIIBorder()).BorderForeground(lipgloss.Color(randomColorHash())).Margin(0)
		dicePaneStyle = dicePaneStyle.Border(lipgloss.ASCIIBorder()).BorderForeground(lipgloss.Color(randomColorHash())).Margin(0)
	}

	cardsPane := cardsPaneStyle.Render(m.game.RenderCards())
	smallPointsPane := (m.height > 0 && m.height < 40) || m.width < 125
	pointsPane := m.game.RenderPoints(smallPointsPane, debug)
	dicePane := dicePaneStyle.Render(m.game.RenderDice())

	content := lipgloss.JoinHorizontal(
		lipgloss.Left,
		pointsPane,
		cardsPane,
		dicePane,
	)

	boardPanes := lipgloss.JoinVertical(
		lipgloss.Top,
		content,
	)
	fullPaneStyle := lipgloss.NewStyle().Width(m.width-1).Height(m.height-1).Align(lipgloss.Center, lipgloss.Center).Margin(1)
	if debug {
		fullPaneStyle = fullPaneStyle.Width(m.width - 2).Height(m.height - 3).Border(lipgloss.ASCIIBorder()).BorderForeground(lipgloss.Color(randomColorHash())).Margin(0)
	}

	controlsPaneStyle := lipgloss.NewStyle().AlignHorizontal(lipgloss.Left).MarginRight(1).MarginLeft(1).Foreground(lipgloss.Color("#999999"))
	controls := []string{
		controlsPaneStyle.Render("Press q to quit"),
		controlsPaneStyle.Render("1=X, 2=#, 3=+, 4=↟, 6=%"),
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		fullPaneStyle.Render(boardPanes),
		strings.Join(controls, controlsPaneStyle.Render("|")),
	)
}

func randomColorHash() string {
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)

	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

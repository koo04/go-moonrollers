package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	mr "github.com/ascii-arcade/moonrollers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

type menuChoice struct {
	s         string
	shortKeys []string
	action    func(m menuModel) (tea.Model, tea.Cmd)
}

type menuModel struct {
	width           int
	height          int
	cursor          int
	numberOfPlayers int
	choices         []menuChoice
	splashScreen    bool
}

func (m menuModel) Init() tea.Cmd {
	w, h, _ := term.GetSize(os.Stdout.Fd())
	m.height = h
	m.width = w
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tMsg(t)
	})
}

/*
not happy with the idea of choices. i thought maybe they
should control their own actions, but it feels convoluted.
just wanted a way to have both a selectable list of options
and a way to have a short key for each option. without repeating
code.
*/
func newMenuModel() *menuModel {
	return &menuModel{
		cursor:          1,
		numberOfPlayers: 3,
		splashScreen:    true,
		choices: []menuChoice{
			{
				s:         "Number of Players",
				shortKeys: []string{"←", "→"},
				action: func(m menuModel) (tea.Model, tea.Cmd) {
					return m, nil
				},
			},
			{
				s:         "New Game",
				shortKeys: []string{"n"},
				action: func(m menuModel) (tea.Model, tea.Cmd) {
					g := mr.NewGame(m.numberOfPlayers)
					w, h, _ := term.GetSize(os.Stdout.Fd())
					_, _ = tea.NewProgram(gameModel{
						game:   g,
						height: h,
						width:  w,
					}, tea.WithAltScreen()).Run()
					return m, nil
				},
			},
			{
				s:         "Join Game",
				shortKeys: []string{"j"},
				action: func(m menuModel) (tea.Model, tea.Cmd) {
					return m, nil
				},
			},
			{
				s:         "Exit",
				shortKeys: []string{"e"},
				action: func(m menuModel) (tea.Model, tea.Cmd) {
					return m, tea.Quit
				},
			},
		},
	}
}

type tMsg time.Time

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.splashScreen {
			return m, nil
		}
		switch msg.Type {
		case tea.KeyEnter:
			return m.choices[m.cursor].action(m)
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyRunes:
			switch string(msg.Runes) {
			case "e":
				return m, tea.Quit
			default:
				for _, choice := range m.choices {
					if strings.Contains(choice.s, string(msg.Runes)) {
						return choice.action(m)
					}
				}
			}
		case tea.KeyUp:
			if m.cursor > 1 {
				m.cursor--
			}
		case tea.KeyDown:
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case tea.KeyLeft:
			if m.numberOfPlayers > 2 {
				m.numberOfPlayers--
			}
		case tea.KeyRight:
			if m.numberOfPlayers < 5 {
				m.numberOfPlayers++
			}
		default:
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tMsg:
		m.splashScreen = false
		return m, nil
	}

	return m, nil
}

func (m menuModel) View() string {
	if m.height < 15 || m.width < 100 {
		return "Window too small, please resize to something larger."
	}

	if m.splashScreen {
		return lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Width(m.width).Height(m.height).Foreground(lipgloss.Color("#fff")).Render(logo)
	}

	logoColor := lipgloss.Color("#eb5761ff")

	menuBaseStyle := lipgloss.NewStyle().Foreground(logoColor).BorderForeground(logoColor).Align(lipgloss.Center)

	logo := menuBaseStyle.Width(m.width / 3).AlignVertical(lipgloss.Center).Render(moonrollersLogo)
	title := menuBaseStyle.Border(lipgloss.NormalBorder()).Margin(1).Padding(1, 2).Align(lipgloss.Center, lipgloss.Center).Render("Moon Rollers")
	menu := ""
	for i, choice := range m.choices {
		sk := strings.Join(choice.shortKeys, "/")
		if i == 0 {
			menu += lipgloss.NewStyle().Foreground(lipgloss.Color("#fff")).Render(fmt.Sprintf("Number of Players: %d (%s)", m.numberOfPlayers, sk)) + "\n"
			continue
		}
		if i == m.cursor {
			menu += lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00")).Render("-> "+fmt.Sprintf("%s (%s)", choice.s, sk)) + "\n"
		} else {
			menu += lipgloss.NewStyle().Foreground(lipgloss.Color("#fff")).Render("   "+fmt.Sprintf("%s (%s)", choice.s, sk)) + "\n"
		}
	}

	menuJoin := lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		menuBaseStyle.AlignHorizontal(lipgloss.Left).Render(menu),
	)

	menuJoin = lipgloss.NewStyle().Width((m.width / 3) * 2).Height(m.height).AlignVertical(lipgloss.Center).Render(menuJoin)

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		logo,
		menuJoin,
	)
}

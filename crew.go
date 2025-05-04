package moonrollers

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Crew struct {
	Name       string      `json:"name"`
	Objectives []objective `json:"objectives"`
	Selected   int         `json:"selected"`
	Ability    ability     `json:"ability"`
	Starter    bool        `json:"starter"`
	Faction    faction     `json:"faction"`
}

func (c *Crew) RenderName() string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(factionColors[c.Faction])).
		Italic(true).
		Bold(true)

	return style.Render(c.Name)
}

func (c *Crew) Render(selectedObj int) string {
	bPadding := 0
	if len(c.Objectives) < 4 {
		bPadding = 1
	}
	pane := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(factionColors[c.Faction])).
		Padding(0, max(len(c.Name)-10, len(c.Objectives)+3), bPadding, 0)

	content := make([]string, 0)
	content = append(content, c.RenderName()+"\n")

	for _, objective := range c.Objectives {
		line := ""
		line += objective.Type.Abbr()
		if objective.Hazard {
			line += lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("*")
		} else {
			line += " "
		}
		for range objective.Amount {
			line += "o"
		}

		content = append(content, line)
	}

	return pane.Render(strings.Join(content, "\n"))
}

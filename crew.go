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
		Foreground(factionColors[c.Faction]).
		Italic(true).
		Bold(true)

	return style.Render(c.Name)
}

func (c *Crew) Render(selectedObj int) string {
	content := make([]string, 0)
	content = append(content, c.RenderName()+"\n")

	for _, objective := range c.Objectives {
		line := ""
		line += objective.Type.Abbr()
		if objective.Hazard {
			line += lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("!")
		} else {
			line += " "
		}
		for range objective.Amount {
			line += "o"
		}

		content = append(content, line)
	}

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(11).
		Height(7).
		Padding(0, 0, 0, 1).
		BorderForeground(lipgloss.Color(factionColors[c.Faction])).
		Render(strings.Join(content, "\n"))
}

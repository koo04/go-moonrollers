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
	content = append(content, c.renderObjectives())

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(11).
		Height(7).
		BorderForeground(lipgloss.Color(factionColors[c.Faction])).
		Render(strings.Join(content, "\n"))
}

func (c *Crew) renderObjectives() string {
	objectives := make([]string, 0)
	for _, objective := range c.Objectives {
		line := " " + objective.Type.Abbr() + " "
		if objective.Hazard {
			line += lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("!")
		} else {
			line += " "
		}

		for range objective.Amount {
			if objective.CompletedBy != nil {
				line += lipgloss.NewStyle().Render("■")
			} else {
				line += "o"
			}
		}

		objectives = append(objectives, line)
	}
	return lipgloss.NewStyle().Render(strings.Join(objectives, "\n"))
}

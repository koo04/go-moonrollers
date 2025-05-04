package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type card struct {
	Name       string      `json:"name"`
	Objectives []objective `json:"objectives"`
	Selected   int         `json:"selected"`
	Ability    ability     `json:"ability"`
	Starter    bool        `json:"starter"`
	Faction    faction     `json:"faction"`
}

var deck = []*card{
	{
		Name:    "Aponi",
		Faction: FactionBlue,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeReactor,
				Amount: 4,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeWild,
				Amount: 2,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Vila",
		Faction: FactionBlue,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeReactor,
				Amount: 4,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Salatar",
		Faction: FactionBlue,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeReactor,
				Amount: 4,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Ada",
		Faction: FactionBlue,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeReactor,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
		Starter: true,
	},
	{
		Name:    "Lee",
		Faction: FactionBlue,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeReactor,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 2,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Lila",
		Faction: FactionBlue,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeReactor,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "[REDACTED]",
		Faction: FactionGreen,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeShield,
				Amount: 4,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeWild,
				Amount: 2,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Imdar",
		Faction: FactionGreen,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeShield,
				Amount: 4,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Namari",
		Faction: FactionGreen,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeShield,
				Amount: 4,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Ryle",
		Faction: FactionGreen,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeShield,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 1,
			},
		},
		Ability: ability{},
		Starter: true,
	},
	{
		Name:    "Bill",
		Faction: FactionGreen,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeShield,
				Amount: 2,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "AT-OK",
		Faction: FactionGreen,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeShield,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Dr.Umbrage",
		Faction: FactionOrange,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeDamage,
				Amount: 4,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeWild,
				Amount: 2,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Saghari",
		Faction: FactionOrange,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeDamage,
				Amount: 4,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Kary",
		Faction: FactionOrange,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeDamage,
				Amount: 4,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Dana",
		Faction: FactionOrange,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeDamage,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
		Starter: true,
	},
	{
		Name:    "Tantin",
		Faction: FactionOrange,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeDamage,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 2,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Ryan",
		Faction: FactionOrange,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeDamage,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Moro",
		Faction: FactionPurple,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeReactor,
				Amount: 4,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 2,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Vanta",
		Faction: FactionPurple,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeWild,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeWild,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeWild,
				Amount: 1,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Meg",
		Faction: FactionPurple,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeThruster,
				Amount: 4,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Sella",
		Faction: FactionPurple,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeThruster,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 1,
			},
		},
		Ability: ability{},
		Starter: true,
	},
	{
		Name:    "FT-1000",
		Faction: FactionPurple,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeShield,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 2,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 2,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 2,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Avari",
		Faction: FactionPurple,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeDamage,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 2,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Sol",
		Faction: FactionYellow,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeThruster,
				Amount: 4,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeWild,
				Amount: 2,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "B3-AR",
		Faction: FactionYellow,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeThruster,
				Amount: 4,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeThruster,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Kal",
		Faction: FactionYellow,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeThruster,
				Amount: 4,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Nella",
		Faction: FactionYellow,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeThruster,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
		Starter: true,
	},
	{
		Name:    "Zek",
		Faction: FactionYellow,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeThruster,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 2,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
	{
		Name:    "Myla",
		Faction: FactionYellow,
		Objectives: []objective{
			{
				Type:   ObjectiveTypeThruster,
				Amount: 3,
				Hazard: true,
			},
			{
				Type:   ObjectiveTypeReactor,
				Amount: 3,
			},
			{
				Type:   ObjectiveTypeDamage,
				Amount: 2,
			},
			{
				Type:   ObjectiveTypeShield,
				Amount: 1,
				Hazard: true,
			},
		},
		Ability: ability{},
	},
}

func (c *card) RenderName() string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(factionColors[c.Faction])).
		Italic(true).
		Bold(true)

	return style.Render(c.Name)
}

func (c *card) Render(selectedObj int) string {
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

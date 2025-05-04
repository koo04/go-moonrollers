package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type card struct {
	Name       string
	Objectives []objective
	Selected   int
	Ability    ability
	Starter    bool
	Faction    faction
}

var deck = []*card{
	{
		Name:    "Aponi",
		Faction: FactionTypeBlue,
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
		Faction: FactionTypeBlue,
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
		Faction: FactionTypeBlue,
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
		Faction: FactionTypeBlue,
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
		Faction: FactionTypeBlue,
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
		Faction: FactionTypeBlue,
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
		Faction: FactionTypeGreen,
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
		Faction: FactionTypeGreen,
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
		Faction: FactionTypeGreen,
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
		Faction: FactionTypeGreen,
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
		Faction: FactionTypeGreen,
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
		Faction: FactionTypeGreen,
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
		Faction: FactionTypeOrange,
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
		Faction: FactionTypeOrange,
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
		Faction: FactionTypeOrange,
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
		Faction: FactionTypeOrange,
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
		Faction: FactionTypeOrange,
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
		Faction: FactionTypeOrange,
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
		Faction: FactionTypePurple,
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
		Faction: FactionTypePurple,
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
		Faction: FactionTypePurple,
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
		Faction: FactionTypePurple,
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
		Faction: FactionTypePurple,
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
		Faction: FactionTypePurple,
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
		Faction: FactionTypeYellow,
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
		Faction: FactionTypeYellow,
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
		Faction: FactionTypeYellow,
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
		Faction: FactionTypeYellow,
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
		Faction: FactionTypeYellow,
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
		Faction: FactionTypeYellow,
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
		Foreground(lipgloss.Color(c.Faction.color)).
		Italic(true).
		Bold(true)

	return style.Render(c.Name)
}

func (c *card) Render(selectedObj int) string {
	pane := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(c.Faction.color)).
		Padding(0, 1)

	content := make([]string, 0)
	content = append(content, c.RenderName())

	for _, objective := range c.Objectives {
		line := make([]string, 0)
		line = append(line, objective.Type.Abbr())
		if objective.Hazard {
			line = append(line, lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("*"))
		} else {
			line = append(line, " ")
		}
		for range objective.Amount {
			line = append(line, "o")
		}

		content = append(content, strings.Join(line, ""))
	}

	return pane.Render(strings.Join(content, "\n"))
}

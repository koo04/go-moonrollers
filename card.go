package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type card struct {
	Name       string
	Objectives []objective
	Selected   int
	Ability    ability
	Starter    bool
}

var deck = []*card{
	{
		Name: "Aponi",
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
		Name: "Vila",
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
		Name: "Salatar",
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
		Name: "Ada",
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
		Name: "Lee",
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
		Name: "Lila",
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
		Name: "[REDACTED]",
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
		Name: "Imdar",
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
		Name: "Namari",
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
		Name: "Ryle",
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
		Name: "Bill",
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
		Name: "AT-OK",
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
		Name: "Dr.Umbrage",
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
		Name: "Saghari",
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
		Name: "Kary",
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
		Name: "Dana",
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
		Name: "Tantin",
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
		Name: "Ryan",
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
		Name: "Moro",
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
		Name: "Vanta",
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
		Name: "Meg",
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
		Name: "Sella",
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
		Name: "FT-1000",
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
		Name: "Avari",
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
		Name: "Sol",
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
		Name: "B3-AR",
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
		Name: "Kal",
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
		Name: "Nella",
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
		Name: "Zek",
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
		Name: "Myla",
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

func (c *card) Render(selectedObj int) string {
	extraWidth := 0
	if len(c.Name) < 10 {
		extraWidth = 10 - len(c.Name)
	}

	top := "┌──────────┐\n"
	name := fmt.Sprintf("│%s│\n", c.Name+strings.Repeat(" ", extraWidth))
	midTop := "│          │\n"
	midBot := "\n│          │\n"
	bottom := "└──────────┘"

	objectives := make([]string, len(c.Objectives))
	for i, o := range c.Objectives {
		h := fmt.Sprintf("%s%d", o.Type.Abbr(), o.Amount)
		if o.Hazard {
			h += color.RedString("*")
		} else {
			h += " "
		}
		if selectedObj == i {
			h += " <"
		}
		length := len(h)
		if o.Hazard {
			length -= 9
		}
		leftOver := 8 - length
		if leftOver > 0 {
			h += strings.Repeat(" ", leftOver)
		}

		objectives[i] = fmt.Sprintf("│ %s │", h)
	}

	if len(objectives) < 4 {
		objectives = append(objectives, "│          │")
	}

	return top + name + midTop + strings.Join(objectives, "\n") + midBot + bottom
}

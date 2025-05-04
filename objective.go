package main

import "github.com/charmbracelet/lipgloss"

type objectiveType string

const (
	ObjectiveTypeDamage   objectiveType = "damage"
	ObjectiveTypeShield   objectiveType = "shield"
	ObjectiveTypeThruster objectiveType = "thruster"
	ObjectiveTypeReactor  objectiveType = "reactor"
	ObjectiveTypeWild     objectiveType = "wild"
)

type objective struct {
	Type         objectiveType
	Amount       int
	Hazard       bool
	WhoCompleted *player
}

func (ot objectiveType) Abbr() string {
	switch ot {
	case ObjectiveTypeDamage:
		return lipgloss.NewStyle().Foreground(lipgloss.Color(colors["orange"])).Render("X ")
	case ObjectiveTypeShield:
		return lipgloss.NewStyle().Foreground(lipgloss.Color(colors["green"])).Render("# ")
	case ObjectiveTypeThruster:
		return lipgloss.NewStyle().Foreground(lipgloss.Color(colors["yellow"])).Render("> ")
	case ObjectiveTypeReactor:
		return lipgloss.NewStyle().Foreground(lipgloss.Color(colors["blue"])).Render("@ ")
	case ObjectiveTypeWild:
		return lipgloss.NewStyle().Foreground(lipgloss.Color(colors["white"])).Render("% ")
	default:
		return ""
	}
}

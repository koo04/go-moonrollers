package moonrollers

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
	Type   objectiveType
	Amount int
	Hazard bool
}

func (ot objectiveType) Abbr() string {
	switch ot {
	case ObjectiveTypeDamage:
		return lipgloss.NewStyle().Foreground(factionColors[FactionOrange]).Render("X ")
	case ObjectiveTypeShield:
		return lipgloss.NewStyle().Foreground(factionColors[FactionGreen]).Render("# ")
	case ObjectiveTypeThruster:
		return lipgloss.NewStyle().Foreground(factionColors[FactionYellow]).Render("> ")
	case ObjectiveTypeReactor:
		return lipgloss.NewStyle().Foreground(factionColors[FactionBlue]).Render("@ ")
	case ObjectiveTypeWild:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#fff")).Render("% ")
	default:
		return ""
	}
}

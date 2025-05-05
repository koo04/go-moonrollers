package moonrollers

import "github.com/charmbracelet/lipgloss"

type dieSide int

const (
	DieSideDamage dieSide = iota
	DieSideShield
	DieSideThruster
	DieSideReactor
	DieSideWild
	DieSideAdd
)

var dieStyle = lipgloss.NewStyle().Height(1).Border(lipgloss.RoundedBorder()).Padding(0, 1).Align(lipgloss.Center)

var Die = map[string]string{
	"damage":   dieStyle.BorderForeground(factionColors[FactionOrange]).Render(ObjectiveTypeDamage.Abbr()),
	"shield":   dieStyle.BorderForeground(factionColors[FactionGreen]).Render(ObjectiveTypeShield.Abbr()),
	"wild":     dieStyle.BorderForeground(lipgloss.Color("#fff")).Render(ObjectiveTypeWild.Abbr()),
	"add":      dieStyle.BorderForeground(factionColors[FactionPurple]).Foreground(factionColors[FactionPurple]).Render("+"),
	"thruster": dieStyle.BorderForeground(factionColors[FactionYellow]).Render(ObjectiveTypeThruster.Abbr()),
	"reactor":  dieStyle.BorderForeground(factionColors[FactionBlue]).Render(ObjectiveTypeReactor.Abbr()),
}

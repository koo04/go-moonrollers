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
	"damage":   dieStyle.BorderForeground(factionColors[FactionOrange]).Foreground(factionColors[FactionOrange]).Render("X"),
	"shield":   dieStyle.BorderForeground(factionColors[FactionGreen]).Foreground(factionColors[FactionGreen]).Render("#"),
	"wild":     dieStyle.BorderForeground(lipgloss.Color("#fff")).Foreground(lipgloss.Color("#fff")).Render("%"),
	"add":      dieStyle.BorderForeground(factionColors[FactionPurple]).Foreground(factionColors[FactionPurple]).Render("+"),
	"thruster": dieStyle.BorderForeground(factionColors[FactionYellow]).Foreground(factionColors[FactionYellow]).Render("^"),
	"reactor":  dieStyle.BorderForeground(factionColors[FactionBlue]).Foreground(factionColors[FactionBlue]).Render("@"),
}

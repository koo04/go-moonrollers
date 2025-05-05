package moonrollers

import (
	"math/rand"

	"github.com/charmbracelet/lipgloss"
)

type dieSide int

const (
	DieSideDamage dieSide = iota
	DieSideShield
	DieSideThruster
	DieSideReactor
	DieSideWild
	DieSideAdd
)

type Die struct {
	LandedSide dieSide
}

var dieStyle = lipgloss.NewStyle().Height(1).Border(lipgloss.RoundedBorder()).Padding(0, 1).Align(lipgloss.Center)

var dieRenders = map[dieSide]string{
	DieSideDamage:   dieStyle.BorderForeground(factionColors[FactionOrange]).Render(ObjectiveTypeDamage.Abbr()),
	DieSideShield:   dieStyle.BorderForeground(factionColors[FactionGreen]).Render(ObjectiveTypeShield.Abbr()),
	DieSideWild:     dieStyle.BorderForeground(lipgloss.Color("#fff")).Render(ObjectiveTypeWild.Abbr()),
	DieSideAdd:      dieStyle.BorderForeground(factionColors[FactionPurple]).Foreground(factionColors[FactionPurple]).Render("+"),
	DieSideThruster: dieStyle.BorderForeground(factionColors[FactionYellow]).Render(ObjectiveTypeThruster.Abbr()),
	DieSideReactor:  dieStyle.BorderForeground(factionColors[FactionBlue]).Render(ObjectiveTypeReactor.Abbr()),
}

func (d *Die) Render() string {
	return dieRenders[d.LandedSide]
}

func (d *Die) Roll() *Die {
	d.LandedSide = dieSide(rand.Intn(len(dieRenders)))
	return d
}

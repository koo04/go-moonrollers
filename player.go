package moonrollers

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type Player struct {
	Name   string
	Points int
	Cards  []*Crew
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:   name,
		Points: 0,
		Cards:  make([]*Crew, 0),
	}
}

func (p *Player) RenderPoints() string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00ff00")).
		Render(p.Name + ": " + strconv.Itoa(p.Points))
}

func (p *Player) GetCrewColor() lipgloss.Color {
	return lipgloss.Color(factionColors[p.Cards[0].Faction])
}

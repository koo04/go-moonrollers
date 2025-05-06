package moonrollers

import (
	"math/rand/v2"

	"github.com/charmbracelet/lipgloss"
)

type Player struct {
	Name   string
	Points int
	Cards  []*Crew
	Color  lipgloss.Color
}

func NewPlayer(name string, color lipgloss.Color) *Player {
	return &Player{
		Name:   name,
		Points: rand.IntN(50) + 1,
		Cards:  make([]*Crew, 0),
		Color:  color,
	}
}

func (p *Player) RenderPoints() string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00ff00")).
		Render(p.Name + ": " + strconv.Itoa(p.Points))
}

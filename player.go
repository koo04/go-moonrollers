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
		Points: rand.IntN(180) + 1,
		Cards:  make([]*Crew, 0),
		Color:  color,
	}
}

package player

import "github.com/koo04/go-moonrollers/card"

type Player struct {
	Name   string
	Points int
	Cards  []*card.Crew
}

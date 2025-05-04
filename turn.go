package main

import (
	"github.com/koo04/go-moonrollers/card"
	"github.com/koo04/go-moonrollers/player"
)

type turn struct {
	Id           int
	Player       *player.Player
	SelectedCard *card.Crew
	Roll         []dieSide
	RollSize     int
	DiePoolSize  int
}

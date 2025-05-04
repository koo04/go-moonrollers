package main

type turn struct {
	Id                int
	Player            *player
	SelectedCard      *card
	SelectedObjective *objective
	Roll              []dieSide
	RollSize          int
	DiePoolSize       int
}

package moonrollers

type turn struct {
	Id           int
	Player       *Player
	SelectedCard *Crew
	Roll         []dieSide
	RollSize     int
	DiePoolSize  int
}

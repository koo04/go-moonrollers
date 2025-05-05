package moonrollers

type Turn struct {
	Id           int
	Player       *Player
	SelectedCard *Crew
	Roll         []dieSide
	RollSize     int
	DiePoolSize  int
}

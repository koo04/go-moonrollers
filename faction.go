package moonrollers

import "github.com/charmbracelet/lipgloss"

type faction string

var (
	FactionBlue   = faction("blue")
	FactionYellow = faction("yellow")
	FactionOrange = faction("orange")
	FactionGreen  = faction("green")
	FactionPurple = faction("purple")
)

var factionColors = map[faction]lipgloss.Color{
	FactionBlue:   lipgloss.Color("#2677fe"),
	FactionYellow: lipgloss.Color("#fefe26"),
	FactionOrange: lipgloss.Color("#CC5500"),
	FactionGreen:  lipgloss.Color("#23741e"),
	FactionPurple: lipgloss.Color("#ab19a6"),
}

func factionColorsSlice() []lipgloss.Color {
	return []lipgloss.Color{
		factionColors[FactionBlue],
		factionColors[FactionYellow],
		factionColors[FactionOrange],
		factionColors[FactionGreen],
		factionColors[FactionPurple],
	}
}

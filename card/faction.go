package card

type faction string

var (
	FactionBlue   = faction("blue")
	FactionYellow = faction("yellow")
	FactionOrange = faction("orange")
	FactionGreen  = faction("green")
	FactionPurple = faction("purple")
)

var factionColors = map[faction]string{
	FactionBlue:   "#2677fe",
	FactionYellow: "#fefe26",
	FactionOrange: "#CC5500",
	FactionGreen:  "#23741e",
	FactionPurple: "#ab19a6",
}

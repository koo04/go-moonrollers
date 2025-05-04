package main

type faction struct {
	name  string
	color string
}

var (
	FactionTypeBlue   = faction{name: "blue", color: colors["blue"]}
	FactionTypeYellow = faction{name: "yellow", color: colors["yellow"]}
	FactionTypeOrange = faction{name: "orange", color: colors["orange"]}
	FactionTypeGreen  = faction{name: "green", color: colors["green"]}
	FactionTypePurple = faction{name: "purple", color: colors["purple"]}
)

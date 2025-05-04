package main

type dieSide int

const (
	DieSideDamage dieSide = iota
	DieSideShield
	DieSideThruster
	DieSideReactor
	DieSideWild
	DieSideAdd
)

var Die = map[dieSide][]string{
	DieSideDamage: {
		"┌───┐",
		"│ D │",
		"└───┘",
	},
	DieSideShield: {
		"┌───┐",
		"│ S │",
		"└───┘",
	},
	DieSideThruster: {
		"┌───┐",
		"│ T │",
		"└───┘",
	},
	DieSideReactor: {
		"┌───┐",
		"│ R │",
		"└───┘",
	},
	DieSideWild: {
		"┌───┐",
		"│ W │",
		"└───┘",
	},
	DieSideAdd: {
		"┌───┐",
		"│ + │",
		"└───┘",
	},
}

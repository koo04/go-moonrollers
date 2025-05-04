package main

type objectiveType string

const (
	ObjectiveTypeDamage   objectiveType = "damage"
	ObjectiveTypeShield   objectiveType = "shield"
	ObjectiveTypeThruster objectiveType = "thruster"
	ObjectiveTypeReactor  objectiveType = "reactor"
	ObjectiveTypeWild     objectiveType = "wild"
)

type objective struct {
	Type         objectiveType
	Amount       int
	Hazard       bool
	WhoCompleted *player
}

func (ot objectiveType) Abbr() string {
	switch ot {
	case ObjectiveTypeDamage:
		return "D"
	case ObjectiveTypeShield:
		return "S"
	case ObjectiveTypeThruster:
		return "T"
	case ObjectiveTypeReactor:
		return "R"
	case ObjectiveTypeWild:
		return "W"
	default:
		return ""
	}
}

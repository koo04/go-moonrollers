package moonrollers

import (
	"math/rand"
	"slices"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type Game struct {
	Id            int
	Players       []*Player
	Turns         []*Turn
	CurrentPlayer int
	Deck          []*Crew
	Out           []*Crew
}

func NewGame(playerCount int) *Game {

	players := make([]*Player, 0, playerCount)
	for range playerCount {
		players = append(players, NewPlayer("Player "+strconv.Itoa(len(players)+1)))
	}
	g := &Game{
		Id:            1,
		Players:       players,
		Turns:         make([]*Turn, 0),
		CurrentPlayer: 0,
		Deck:          Deck,
		Out:           make([]*Crew, 0, 6),
	}
	g.FirstRoundDraw(playerCount)

	return g
}

func (g *Game) shuffleDeck() {
	for i := len(g.Deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
	}
}

func (g *Game) drawCard() {
	if len(g.Deck) == 0 {
		return
	}
	card := g.Deck[0]
	g.Deck = g.Deck[1:]
	g.Out = append(g.Out, card)
}

func hasFaction(s []*Crew, c2 *Crew) bool {
	for _, c := range s {
		if c.Faction == c2.Faction {
			return true
		}
	}
	return false
}

func (g *Game) FirstRoundDraw(playerCount int) {
	g.shuffleDeck()

	maxDraw := 6
	switch playerCount {
	case 2:
		maxDraw = 4
	case 3:
		maxDraw = 5
	}

	onTheSide := make([]*Crew, 0)
	for len(g.Out) < maxDraw {
		card := g.Deck[0]
		if hasFaction(g.Out, card) && len(g.Out) < 5 {
			onTheSide = append(onTheSide, card)
			g.Deck = slices.Delete(g.Deck, 0, 1)
			continue
		}
		g.Out = append(g.Out, card)
		g.Deck = slices.Delete(g.Deck, 0, 1)
	}

	g.Deck = append(g.Deck, onTheSide...)
	g.shuffleDeck()
}

func (g *Game) RenderCards() string {
	rowOne := []string{}
	rowTwo := []string{}
	for i, c := range g.Out {
		if i < 3 {
			rowOne = append(rowOne, c.Render(-1))
			continue
		}

		rowTwo = append(rowTwo, c.Render(-1))
	}

	renderedCardsRowOne := lipgloss.JoinHorizontal(lipgloss.Left, rowOne...)
	renderedCardsRowTwo := lipgloss.JoinHorizontal(lipgloss.Left, rowTwo...)

	return lipgloss.JoinVertical(lipgloss.Top, renderedCardsRowOne, renderedCardsRowTwo)
}

func (g *Game) RenderPoints() string {
	out := make([]string, 0)
	for _, p := range g.Players {
		out = append(out, p.RenderPoints())
	}

	return lipgloss.JoinVertical(lipgloss.Left, out...)
}

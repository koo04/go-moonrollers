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
	Supply        []*Die
	Pool          []*Die
}

func NewGame(playerCount int) *Game {
	g := &Game{
		Id:            1,
		Turns:         make([]*Turn, 0),
		CurrentPlayer: 0,
		Deck:          CopyDeck(),
		Out:           make([]*Crew, 0, 6),
	}

	return g.setup(playerCount)
}

func (g *Game) setup(playerCount int) *Game {
	players := make([]*Player, 0, playerCount)
	for i := range playerCount {
		color := lipgloss.Color(factionColorsSlice()[i])
		players = append(players, NewPlayer(
			"Player "+strconv.Itoa(len(players)+1),
			color,
		))
	}

	supply := make([]*Die, 0, 7)
	for range 7 {
		d := &Die{}
		d.Roll()
		supply = append(supply, d)
	}
	pool := make([]*Die, 0, 12)
	for range 12 {
		d := &Die{}
		d.Roll()
		pool = append(pool, d)
	}

	g.Players = players
	g.Supply = supply
	g.Pool = pool

	return g.FirstRoundDraw(playerCount)
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

func (g *Game) FirstRoundDraw(playerCount int) *Game {
	g.shuffleDeck()

	maxDraw := 6
	switch playerCount {
	case 2:
		maxDraw = 4
	case 3:
		maxDraw = 5
	}

	unused := make([]*Crew, 0)
	for len(g.Out) < maxDraw {
		card := g.Deck[0]
		if hasFaction(g.Out, card) && len(g.Out) < 5 {
			unused = append(unused, card)
			g.Deck = slices.Delete(g.Deck, 0, 1)
			continue
		}
		g.Out = append(g.Out, card)
		g.Deck = slices.Delete(g.Deck, 0, 1)
	}

	g.Deck = append(g.Deck, unused...)
	g.shuffleDeck()

	return g
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

func (g *Game) RenderDice() string {
	supplyRender := lipgloss.NewStyle().AlignVertical(lipgloss.Left).MarginRight(5)
	poolRender := lipgloss.NewStyle().AlignVertical(lipgloss.Center)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		supplyRender.Render(g.RenderSupply()),
		poolRender.Render(g.RenderPool()),
	)
}

func (g *Game) RenderSupply() string {
	return g.renderDieGroup(g.Supply, 3, false)
}

func (g *Game) RenderPool() string {
	return g.renderDieGroup(g.Pool, 4, true)
}

func (g *Game) renderDieGroup(dieGroup []*Die, maxWidth int, useControls bool) string {
	rowOne := make([]string, 0, maxWidth)
	rowTwo := make([]string, 0, maxWidth)
	rowThree := make([]string, 0, maxWidth)
	for i, d := range dieGroup {
		if d == nil {
			continue
		}
		out := d.Render()
		if useControls {
			var controlLabel string
			switch i {
			case 9:
				controlLabel = "0"
			case 10:
				controlLabel = "-"
			case 11:
				controlLabel = "+"
			default:
				controlLabel = strconv.Itoa(i + 1)
			}
			controls := lipgloss.NewStyle().Foreground(lipgloss.Color("#fff")).Render(controlLabel)
			out = lipgloss.JoinVertical(lipgloss.Center, d.Render(), controls)
		}

		if i < maxWidth {
			rowOne = append(rowOne, out)
			continue
		}
		if i < maxWidth*2 {
			rowTwo = append(rowTwo, out)
			continue
		}
		rowThree = append(rowThree, out)
	}

	renderedRowOne := lipgloss.JoinHorizontal(lipgloss.Left, rowOne...)
	renderedRowTwo := lipgloss.JoinHorizontal(lipgloss.Left, rowTwo...)
	renderedRowThree := lipgloss.JoinHorizontal(lipgloss.Left, rowThree...)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		renderedRowOne,
		renderedRowTwo,
		renderedRowThree,
	)
}

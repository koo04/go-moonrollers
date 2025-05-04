package moonrollers

import (
	"math/rand"

	"github.com/charmbracelet/lipgloss"
)

type Board struct {
	Deck []*Crew
	Out  []*Crew
}

func (b *Board) ShuffleDeck() {
	for i := len(b.Deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		b.Deck[i], b.Deck[j] = b.Deck[j], b.Deck[i]
	}
}

func (b *Board) DrawCard() {
	if len(b.Deck) == 0 {
		return
	}
	card := b.Deck[0]
	b.Deck = b.Deck[1:]
	b.Out = append(b.Out, card)
}

func (b *Board) DrawCards(n int) {
	for range n {
		b.DrawCard()
	}
}

func (b *Board) Render() string {
	var renderedCards []string
	for _, c := range b.Out {
		renderedCards = append(renderedCards, c.Render(-1))
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, renderedCards...)
}

package main

import (
	"math/rand"

	"github.com/charmbracelet/lipgloss"
)

type board struct {
	Deck []*card
	Out  []*card
}

func (b *board) ShuffleDeck() {
	for i := len(b.Deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		b.Deck[i], b.Deck[j] = b.Deck[j], b.Deck[i]
	}
}

func (b *board) DrawCard() {
	if len(b.Deck) == 0 {
		return
	}
	card := b.Deck[0]
	b.Deck = b.Deck[1:]
	b.Out = append(b.Out, card)
}

func (b *board) DrawCards(n int) {
	for range n {
		b.DrawCard()
	}
}

func (b *board) Render() string {
	var renderedCards []string
	for _, c := range b.Out {
		renderedCards = append(renderedCards, c.Render(-1))
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, renderedCards...)
}

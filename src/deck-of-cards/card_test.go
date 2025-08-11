package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Heart, Rank: Ace}) 
	fmt.Println(Card{Suit: Spade, Rank: Two})
	fmt.Println(Card{Suit: Diamond, Rank: Nine})
	fmt.Println(Card{Suit: Club, Rank: Jack})
	fmt.Println(Card{Suit: Joker})
	// Output: 
	// Ace of Heart
	// Two of Spade
	// Nine of Diamond
	// Jack of Club
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Errorf("expected 52 cards, got %d", len(cards))
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	if cards[0] != (Card{Suit: Spade, Rank: Ace}) {
		t.Errorf("expected %v, got %v", Card{Suit: Spade, Rank: Ace}, cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	if cards[0] != (Card{Suit: Spade, Rank: Ace}) {
		t.Errorf("expected %v, got %v", Card{Suit: Spade, Rank: Ace}, cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count :=0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Errorf("expected 3 jokers, got %d", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c:= range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("expected %d cards, got %d", 13*4*3, len(cards))
	}
}
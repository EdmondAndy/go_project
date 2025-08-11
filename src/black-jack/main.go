package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Card represents a playing card
type Card struct {
	Suit  string
	Rank  string
	Value int
}

// Deck represents a deck of cards
type Deck struct {
	Cards []Card
}

// Hand represents a player's hand
type Hand struct {
	Cards []Card
	Value int
	Aces  int
}

// Player represents a player in the game
type Player struct {
	Name   string
	Hand   Hand
	Money  int
	Bet    int
	Busted bool
}

// Game represents the blackjack game state
type Game struct {
	Deck   Deck
	Player Player
	Dealer Hand
	Reader *bufio.Reader
}

// Initialize a standard deck of 52 cards
func (d *Deck) Initialize() {
	suits := []string{"♠", "♥", "♦", "♣"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	d.Cards = make([]Card, 0, 52)

	for _, suit := range suits {
		for i, rank := range ranks {
			value := i + 1
			if value > 10 {
				value = 10 // Face cards are worth 10
			}
			d.Cards = append(d.Cards, Card{
				Suit:  suit,
				Rank:  rank,
				Value: value,
			})
		}
	}
}

// Shuffle the deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

// Deal a card from the deck
func (d *Deck) Deal() Card {
	if len(d.Cards) == 0 {
		panic("Cannot deal from empty deck")
	}
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

// Add a card to the hand
func (h *Hand) AddCard(card Card) {
	h.Cards = append(h.Cards, card)
	if card.Rank == "A" {
		h.Aces++
	}
	h.calculateValue()
}

// Calculate the best value of the hand
func (h *Hand) calculateValue() {
	h.Value = 0
	for _, card := range h.Cards {
		h.Value += card.Value
	}

	// Handle Aces (1 or 11)
	aces := h.Aces
	for aces > 0 && h.Value <= 11 {
		h.Value += 10 // Make an Ace worth 11 instead of 1
		aces--
	}
}

// Check if hand is blackjack (21 with 2 cards)
func (h *Hand) IsBlackjack() bool {
	return len(h.Cards) == 2 && h.Value == 21
}

// Check if hand is busted (over 21)
func (h *Hand) IsBusted() bool {
	return h.Value > 21
}

// String representation of a card
func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank, c.Suit)
}

// String representation of a hand
func (h Hand) String() string {
	var cards []string
	for _, card := range h.Cards {
		cards = append(cards, card.String())
	}
	return strings.Join(cards, " ")
}

// Initialize a new game
func NewGame() *Game {
	game := &Game{
		Player: Player{
			Name:  "Player",
			Money: 1000,
		},
		Reader: bufio.NewReader(os.Stdin),
	}

	game.Deck.Initialize()
	game.Deck.Shuffle()

	return game
}

// Get bet from player
func (g *Game) getBet() {
	for {
		fmt.Printf("You have $%d. Enter your bet: $", g.Player.Money)
		input, _ := g.Reader.ReadString('\n')
		input = strings.TrimSpace(input)

		bet, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		if bet <= 0 {
			fmt.Println("Bet must be greater than 0.")
			continue
		}

		if bet > g.Player.Money {
			fmt.Println("You don't have enough money for that bet.")
			continue
		}

		g.Player.Bet = bet
		g.Player.Money -= bet
		break
	}
}

// Deal initial cards
func (g *Game) dealInitialCards() {
	// Reset hands
	g.Player.Hand = Hand{}
	g.Dealer = Hand{}
	g.Player.Busted = false

	// Deal two cards to player and dealer
	g.Player.Hand.AddCard(g.Deck.Deal())
	g.Dealer.AddCard(g.Deck.Deal())
	g.Player.Hand.AddCard(g.Deck.Deal())
	g.Dealer.AddCard(g.Deck.Deal())
}

// Display game state
func (g *Game) displayGame(hideDealerCard bool) {
	fmt.Println("\n" + strings.Repeat("=", 50))

	// Show dealer's hand
	if hideDealerCard {
		fmt.Printf("Dealer: %s [Hidden] (Showing: %d)\n", g.Dealer.Cards[0].String(), g.Dealer.Cards[0].Value)
	} else {
		fmt.Printf("Dealer: %s (Total: %d)\n", g.Dealer.String(), g.Dealer.Value)
	}

	// Show player's hand
	fmt.Printf("Player: %s (Total: %d)\n", g.Player.Hand.String(), g.Player.Hand.Value)
	fmt.Printf("Bet: $%d | Money: $%d\n", g.Player.Bet, g.Player.Money)
	fmt.Println(strings.Repeat("=", 50))
}

// Player's turn
func (g *Game) playerTurn() {
	for !g.Player.Hand.IsBusted() && g.Player.Hand.Value != 21 {
		g.displayGame(true)

		fmt.Print("\nDo you want to (h)it or (s)tand? ")
		input, _ := g.Reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		switch input {
		case "h", "hit":
			card := g.Deck.Deal()
			g.Player.Hand.AddCard(card)
			fmt.Printf("You drew: %s\n", card.String())

			if g.Player.Hand.IsBusted() {
				g.Player.Busted = true
				fmt.Printf("Busted! Your total is %d\n", g.Player.Hand.Value)
			}
		case "s", "stand":
			return
		default:
			fmt.Println("Please enter 'h' for hit or 's' for stand.")
		}
	}
}

// Dealer's turn
func (g *Game) dealerTurn() {
	fmt.Println("\nDealer's turn:")
	g.displayGame(false)

	// Dealer hits on 16 and below, stands on 17 and above
	for g.Dealer.Value < 17 {
		card := g.Deck.Deal()
		g.Dealer.AddCard(card)
		fmt.Printf("Dealer draws: %s\n", card.String())
		fmt.Printf("Dealer total: %d\n", g.Dealer.Value)
		time.Sleep(1 * time.Second) // Add dramatic pause
	}
}

// Determine winner and payout
func (g *Game) determineWinner() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("FINAL RESULTS")
	g.displayGame(false)

	playerBJ := g.Player.Hand.IsBlackjack()
	dealerBJ := g.Dealer.IsBlackjack()

	if g.Player.Busted {
		fmt.Println("You busted! Dealer wins.")
		// Money already deducted from bet
	} else if g.Dealer.IsBusted() {
		fmt.Println("Dealer busted! You win!")
		if playerBJ {
			g.Player.Money += int(float64(g.Player.Bet) * 2.5) // Blackjack pays 3:2
			fmt.Printf("Blackjack! You win $%d!\n", int(float64(g.Player.Bet)*1.5))
		} else {
			g.Player.Money += g.Player.Bet * 2
			fmt.Printf("You win $%d!\n", g.Player.Bet)
		}
	} else if playerBJ && dealerBJ {
		fmt.Println("Both have blackjack! It's a push.")
		g.Player.Money += g.Player.Bet // Return bet
	} else if playerBJ {
		fmt.Println("Blackjack! You win!")
		g.Player.Money += int(float64(g.Player.Bet) * 2.5) // Blackjack pays 3:2
		fmt.Printf("You win $%d!\n", int(float64(g.Player.Bet)*1.5))
	} else if dealerBJ {
		fmt.Println("Dealer has blackjack! You lose.")
		// Money already deducted from bet
	} else if g.Player.Hand.Value > g.Dealer.Value {
		fmt.Println("You win!")
		g.Player.Money += g.Player.Bet * 2
		fmt.Printf("You win $%d!\n", g.Player.Bet)
	} else if g.Player.Hand.Value < g.Dealer.Value {
		fmt.Println("Dealer wins!")
		// Money already deducted from bet
	} else {
		fmt.Println("It's a push! (Tie)")
		g.Player.Money += g.Player.Bet // Return bet
	}

	fmt.Printf("Your total money: $%d\n", g.Player.Money)
}

// Play a single round
func (g *Game) playRound() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("                    NEW ROUND")
	fmt.Println(strings.Repeat("=", 60))

	// Get bet
	g.getBet()

	// Deal initial cards
	g.dealInitialCards()

	// Check for immediate blackjacks
	if g.Player.Hand.IsBlackjack() || g.Dealer.IsBlackjack() {
		g.displayGame(false)
		g.determineWinner()
		return
	}

	// Player's turn
	g.playerTurn()

	// Dealer's turn (only if player didn't bust)
	if !g.Player.Busted {
		g.dealerTurn()
	}

	// Determine winner
	g.determineWinner()
}

// Main game loop
func (g *Game) play() {
	fmt.Println("🎰 Welcome to Blackjack! 🎰")
	fmt.Println("Blackjack pays 3:2")
	fmt.Println("Dealer stands on 17")

	for g.Player.Money > 0 {
		g.playRound()

		if g.Player.Money <= 0 {
			fmt.Println("\n💸 You're out of money! Game over!")
			break
		}

		fmt.Print("\nDo you want to play another round? (y/n): ")
		input, _ := g.Reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input != "y" && input != "yes" {
			break
		}

		// Reshuffle deck if running low on cards
		if len(g.Deck.Cards) < 10 {
			fmt.Println("Reshuffling deck...")
			g.Deck.Initialize()
			g.Deck.Shuffle()
		}
	}

	fmt.Printf("\n🎮 Thanks for playing! Final money: $%d\n", g.Player.Money)
}

func main() {
	game := NewGame()
	game.play()
}

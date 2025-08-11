# Blackjack Game in Go

A complete command-line Blackjack game implemented in Go with full game mechanics and betting system.

## Features

- **Complete Blackjack Rules**: Standard Blackjack rules with proper card values
- **Ace Handling**: Intelligent Ace value calculation (1 or 11)
- **Betting System**: Start with $1000 and place bets each round
- **Dealer AI**: Dealer follows standard rules (hits on 16, stands on 17)
- **Blackjack Detection**: Automatic detection and 3:2 payout for Blackjacks
- **Game State Display**: Clear visual representation of cards and values
- **Money Management**: Track winnings and losses
- **Deck Management**: Automatic reshuffling when deck runs low
- **Unicode Cards**: Beautiful card representations with suits (♠♥♦♣)

## How to Play

1. **Starting**: You begin with $1000
2. **Betting**: Enter your bet amount for each round
3. **Dealing**: You and the dealer each get 2 cards (one dealer card is hidden)
4. **Your Turn**: Choose to 'hit' (draw a card) or 'stand' (keep current total)
5. **Goal**: Get as close to 21 as possible without going over
6. **Dealer's Turn**: Dealer reveals hidden card and draws until reaching 17+
7. **Winning**: Beat the dealer's total without busting to win!

## Game Rules

- **Card Values**: 
  - Number cards (2-10): Face value
  - Face cards (J, Q, K): Worth 10 points
  - Aces: Worth 1 or 11 (automatically optimized)
- **Blackjack**: 21 with first 2 cards (pays 3:2)
- **Bust**: Going over 21 (automatic loss)
- **Push**: Tie with dealer (bet returned)
- **Dealer Rules**: Must hit on 16, must stand on 17

## Installation and Running

### Prerequisites
- Go 1.21 or higher

### Running the Game
```bash
# Clone or download the files
go mod tidy
go run main.go
```

### Building an Executable
```bash
go build -o blackjack main.go
./blackjack
```

## Game Commands

During gameplay:
- `h` or `hit`: Draw another card
- `s` or `stand`: Keep current total and end your turn
- `y` or `yes`: Play another round
- `n` or `no`: Quit the game

## Example Gameplay

```
🎰 Welcome to Blackjack! 🎰
Blackjack pays 3:2
Dealer stands on 17

============================================================
                    NEW ROUND
============================================================
You have $1000. Enter your bet: $50

==================================================
Dealer: K♠ [Hidden] (Showing: 10)
Player: 7♥ Q♦ (Total: 17)
Bet: $50 | Money: $950
==================================================

Do you want to (h)it or (s)tand? s

Dealer's turn:
==================================================
Dealer: K♠ 6♣ (Total: 16)
Player: 7♥ Q♦ (Total: 17)
Bet: $50 | Money: $950
==================================================
Dealer draws: 4♥
Dealer total: 20

==================================================
FINAL RESULTS
==================================================
Dealer: K♠ 6♣ 4♥ (Total: 20)
Player: 7♥ Q♦ (Total: 17)
Bet: $50 | Money: $950
==================================================
Dealer wins!
Your total money: $950
```

## Code Structure

- **Card**: Represents individual playing cards with suit, rank, and value
- **Deck**: Manages the 52-card deck with shuffling and dealing
- **Hand**: Represents a collection of cards with automatic value calculation
- **Player**: Represents the player with money management and betting
- **Game**: Main game controller handling all game logic and flow

## Advanced Features

- **Intelligent Ace Handling**: Automatically optimizes Ace values for best hand
- **Blackjack Detection**: Proper 3:2 payout for natural blackjacks
- **Deck Reshuffling**: Automatic deck refresh when cards run low
- **Input Validation**: Robust handling of user input with error messages
- **Money Management**: Complete betting system with win/loss tracking
- **Game State Display**: Clear, formatted output showing all relevant information

Enjoy playing Blackjack! 🎰

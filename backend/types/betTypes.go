package types

type card struct {
	Suit string
	Rank string
}

// Limbo
type BetLimbo struct {
	Currency string
	Amount   float64
	Multi    float64
	Seed     string
}

// CoinToss
type BetCoinToss struct {
	Currency string
	Amount   float64
	Seed     string
}

type NextCoinToss struct {
	Round     string
	Selection string
}

// DragonTower
type BetDragonTower struct {
	Currency   string
	Amount     float64
	Difficulty string
	Seed       string
}

type NextDragonTower struct {
	Egg int
}

// Dice
type BetDice struct {
	Currency  string
	Amount    float64
	Multi     float64
	OverUnder string
	Value     float64
	Seed      string
}

// Slides
// type BetSlides struct {

// }

// Mines
type BetMines struct {
	Currency string
	Amount   float64
	Mines    int
	Seed     string
}

type NextMines struct {
	Field int
}

// Wheel
type BetWheels struct {
	Currency string
	Amount   float64
	Risk     string
	Segments int
	Seed     string
}

// HighLow
type BetHighLow struct {
	Currency  string
	Amount    float64
	StartCard card
	Seed      string
}

type NextHighLow struct {
	Guess string
}

// Pump
type BetPump struct {
	Currency   string
	Amount     float64
	Difficulty string
	Seed       string
}

// BlackJack
type BetBJ struct {
	Currency string
	Amount   float64
	Seed     string
}

type NextBJ struct {
	Action string
}

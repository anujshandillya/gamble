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
}

// Wheel
type BetWheels struct {
	Currency string
	Amount   float64
	Risk     string
	Segments uint8
}

// Dice
type BetDice struct {
	Currency  string
	Amount    float64
	Multi     float64
	OverUnder string
	Value     float64
}

// Slides
// type BetSlides struct {

// }

// CoinToss
type BetCoinToss struct {
	Currency string
	Amount   float64
}

type NextCoinToss struct {
	Round     int
	Selection int
}

// DragonTower
type BetDragonTower struct {
	Currency   string
	Amount     float64
	Difficulty string
}

type NextDragonTower struct {
	Egg   int
	Level int
}

// Mines
type BetMines struct {
	Currency string
	Amount   float64
	Mines    int
}

type NextMines struct {
	Field int
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
}

type NextBJ struct {
	Action string
}

type ActiveBetCoinToss struct {
	Game       string  `json:"game"`
	ServerSeed string  `json:"serverSeed"`
	ClientSeed string  `json:"clientSeed"`
	Nonce      uint16  `json:"nonce"`
	Status     string  `json:"status"`
	Amount     float64 `json:"amount"`
	State      [][]int `json:"state"`
}
type ActiveBetDragonTower struct {
	Game       string  `json:"game"`
	Difficulty string  `json:"difficulty"`
	ServerSeed string  `json:"serverSeed"`
	ClientSeed string  `json:"clientSeed"`
	Nonce      uint16  `json:"nonce"`
	Status     string  `json:"status"`
	Amount     float64 `json:"amount"`
	State      [][]int `json:"state"`
	LevelSet   [][]int `json:"levelSet"`
}
type ActiveBetMines struct {
	Game       string  `json:"game"`
	MinesCount uint8   `json:"minesCount"`
	ServerSeed string  `json:"serverSeed"`
	ClientSeed string  `json:"clientSeed"`
	Nonce      uint16  `json:"nonce"`
	Status     string  `json:"status"`
	Amount     float64 `json:"amount"`
	State      []uint8 `json:"state"`
	MinesSet   []uint8 `json:"minesSet"`
}
type ActiveBetBJ struct {
	Game       string  `json:"game"`
	ServerSeed string  `json:"serverSeed"`
	ClientSeed string  `json:"clientSeed"`
	Nonce      uint16  `json:"nonce"`
	Status     string  `json:"status"`
	Amount     float64 `json:"amount"`
	State      [][]int `json:"state"`
}
type ActiveBetPump struct {
	Game       string  `json:"game"`
	ServerSeed string  `json:"serverSeed"`
	ClientSeed string  `json:"clientSeed"`
	Nonce      uint16  `json:"nonce"`
	Status     string  `json:"status"`
	Amount     float64 `json:"amount"`
	State      [][]int `json:"state"`
}
type ActiveBetHighLow struct {
	Game       string  `json:"game"`
	ServerSeed string  `json:"serverSeed"`
	ClientSeed string  `json:"clientSeed"`
	Nonce      uint16  `json:"nonce"`
	Status     string  `json:"status"`
	Amount     float64 `json:"amount"`
	State      [][]int `json:"state"`
}

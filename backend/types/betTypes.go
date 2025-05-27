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
	Round     uint8
	Selection uint8
}

// DragonTower
type BetDragonTower struct {
	Currency   string
	Amount     float64
	Difficulty string
}

type NextDragonTower struct {
	Egg int
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
	Game       string    `json:"game"`
	ServerSeed string    `json:"serverSeed"`
	ClientSeed string    `json:"clientSeed"`
	Nonce      uint16    `json:"nonce"`
	Status     string    `json:"status"`
	Amount     float64   `json:"amount"`
	State      [][]uint8 `json:"state"`
}
type ActiveBetDragonTower struct {
	Game       string    `json:"game"`
	ServerSeed string    `json:"serverSeed"`
	ClientSeed string    `json:"clientSeed"`
	Nonce      uint16    `json:"nonce"`
	Status     string    `json:"status"`
	Amount     float64   `json:"amount"`
	State      [][]uint8 `json:"state"`
}
type ActiveBetMines struct {
	Game       string    `json:"game"`
	ServerSeed string    `json:"serverSeed"`
	ClientSeed string    `json:"clientSeed"`
	Nonce      uint16    `json:"nonce"`
	Status     string    `json:"status"`
	Amount     float64   `json:"amount"`
	State      [][]uint8 `json:"state"`
}
type ActiveBetBJ struct {
	Game       string    `json:"game"`
	ServerSeed string    `json:"serverSeed"`
	ClientSeed string    `json:"clientSeed"`
	Nonce      uint16    `json:"nonce"`
	Status     string    `json:"status"`
	Amount     float64   `json:"amount"`
	State      [][]uint8 `json:"state"`
}
type ActiveBetPump struct {
	Game       string    `json:"game"`
	ServerSeed string    `json:"serverSeed"`
	ClientSeed string    `json:"clientSeed"`
	Nonce      uint16    `json:"nonce"`
	Status     string    `json:"status"`
	Amount     float64   `json:"amount"`
	State      [][]uint8 `json:"state"`
}
type ActiveBetHighLow struct {
	Game       string    `json:"game"`
	ServerSeed string    `json:"serverSeed"`
	ClientSeed string    `json:"clientSeed"`
	Nonce      uint16    `json:"nonce"`
	Status     string    `json:"status"`
	Amount     float64   `json:"amount"`
	State      [][]uint8 `json:"state"`
}

// Common methods for all active bet types
func (b ActiveBetCoinToss) GetGame() string       { return b.Game }
func (b ActiveBetCoinToss) GetServerSeed() string { return b.ServerSeed }
func (b ActiveBetCoinToss) GetClientSeed() string { return b.ClientSeed }
func (b ActiveBetCoinToss) GetNonce() uint16      { return b.Nonce }
func (b ActiveBetCoinToss) GetStatus() string     { return b.Status }
func (b ActiveBetCoinToss) GetAmount() float64    { return b.Amount }
func (b ActiveBetCoinToss) GetState() [][]uint8   { return b.State }

func (b ActiveBetDragonTower) GetGame() string       { return b.Game }
func (b ActiveBetDragonTower) GetServerSeed() string { return b.ServerSeed }
func (b ActiveBetDragonTower) GetClientSeed() string { return b.ClientSeed }
func (b ActiveBetDragonTower) GetNonce() uint16      { return b.Nonce }
func (b ActiveBetDragonTower) GetStatus() string     { return b.Status }
func (b ActiveBetDragonTower) GetAmount() float64    { return b.Amount }
func (b ActiveBetDragonTower) GetState() [][]uint8   { return b.State }

func (b ActiveBetMines) GetGame() string       { return b.Game }
func (b ActiveBetMines) GetServerSeed() string { return b.ServerSeed }
func (b ActiveBetMines) GetClientSeed() string { return b.ClientSeed }
func (b ActiveBetMines) GetNonce() uint16      { return b.Nonce }
func (b ActiveBetMines) GetStatus() string     { return b.Status }
func (b ActiveBetMines) GetAmount() float64    { return b.Amount }
func (b ActiveBetMines) GetState() [][]uint8   { return b.State }

func (b ActiveBetBJ) GetGame() string       { return b.Game }
func (b ActiveBetBJ) GetServerSeed() string { return b.ServerSeed }
func (b ActiveBetBJ) GetClientSeed() string { return b.ClientSeed }
func (b ActiveBetBJ) GetNonce() uint16      { return b.Nonce }
func (b ActiveBetBJ) GetStatus() string     { return b.Status }
func (b ActiveBetBJ) GetAmount() float64    { return b.Amount }
func (b ActiveBetBJ) GetState() [][]uint8   { return b.State }

func (b ActiveBetPump) GetGame() string       { return b.Game }
func (b ActiveBetPump) GetServerSeed() string { return b.ServerSeed }
func (b ActiveBetPump) GetClientSeed() string { return b.ClientSeed }
func (b ActiveBetPump) GetNonce() uint16      { return b.Nonce }
func (b ActiveBetPump) GetStatus() string     { return b.Status }
func (b ActiveBetPump) GetAmount() float64    { return b.Amount }
func (b ActiveBetPump) GetState() [][]uint8   { return b.State }

func (b ActiveBetHighLow) GetGame() string       { return b.Game }
func (b ActiveBetHighLow) GetServerSeed() string { return b.ServerSeed }
func (b ActiveBetHighLow) GetClientSeed() string { return b.ClientSeed }
func (b ActiveBetHighLow) GetNonce() uint16      { return b.Nonce }
func (b ActiveBetHighLow) GetStatus() string     { return b.Status }
func (b ActiveBetHighLow) GetAmount() float64    { return b.Amount }
func (b ActiveBetHighLow) GetState() [][]uint8   { return b.State }

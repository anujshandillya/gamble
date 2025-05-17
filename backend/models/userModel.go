package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ct struct {
	BetId      string  `json:"bet_id"`
	Wager      float64 `json:"wager"`
	Multi      float64 `json:"multi"`
	CurrentWin float64 `json:"currentwin"`
}

type dt struct {
	BetId      string  `json:"bet_id"`
	Wager      float64 `json:"wager"`
	Multi      float64 `json:"multi"`
	CurrentWin float64 `json:"currentwin"`
	States     [][]int `json:"states"`
}

type mi struct {
	BetId      string  `json:"bet_id"`
	Wager      float64 `json:"wager"`
	Multi      float64 `json:"multi"`
	CurrentWin float64 `json:"currentwin"`
	States     [][]int `json:"states"`
}

type hl struct {
	BetId      string  `json:"bet_id"`
	Wager      float64 `json:"wager"`
	Multi      float64 `json:"multi"`
	CurrentWin float64 `json:"currentwin"`
	States     [][]int `json:"states"`
}
type pp struct {
	BetId      string  `json:"bet_id"`
	Wager      float64 `json:"wager"`
	Multi      float64 `json:"multi"`
	CurrentWin float64 `json:"currentwin"`
	States     [][]int `json:"states"`
}

type bj struct {
	BetId      string  `json:"bet_id"`
	Wager      float64 `json:"wager"`
	Multi      float64 `json:"multi"`
	CurrentWin float64 `json:"currentwin"`
	States     [][]int `json:"states"`
}

type currentBets struct {
	Cointoss    ct `json:"cointoss"`
	Dragontower dt `json:"dragontower"`
	Mines       mi `json:"mines"`
	Highlow     hl `json:"highlow"`
	Pump        pp `json:"pump"`
	Blackjack   bj `json:"bj"`
}

type User struct {
	ID          bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName   string        `json:"firstName"`
	LastName    string        `json:"lastName"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	CurrentBets currentBets   `json:"currentbets"`
}

var UserCollection *mongo.Collection

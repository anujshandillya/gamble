package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ct struct {
	betId      string
	wager      float64
	multi      float64
	currentWin float64
}

type dt struct {
	betId      string
	wager      float64
	multi      float64
	currentWin float64
	states     [][]int
}

type mi struct {
	betId      string
	wager      float64
	multi      float64
	currentWin float64
	states     [][]int
}

type hl struct {
	betId      string
	wager      float64
	multi      float64
	currentWin float64
	states     [][]int
}
type pp struct {
	betId      string
	wager      float64
	multi      float64
	currentWin float64
	states     [][]int
}

type bj struct {
	betId      string
	wager      float64
	multi      float64
	currentWin float64
	states     [][]int
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

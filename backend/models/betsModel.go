package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Bets struct {
	ID         bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID     bson.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Amount     float64       `json:"amount"`
	Game       string        `json:"game"`
	Multiplier float64       `json:"multiplier"`
	Outcome    bool          `json:"outcome"`
	Payout     float64       `json:"payout"`
}

var BetsCollection *mongo.Collection

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Bets struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Amount     float64            `json:"amount"`
	Game       string             `json:"game"`
	Multiplier float64            `json:"multiplier"`
	Outcome    bool               `json:"outcome"`
	Payout     float64            `json:"payout"`
}

var BetsCollection *mongo.Collection

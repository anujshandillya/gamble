package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Transactions struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Type   string             `json:"type"`
	To     string             `json:"to"`
	From   string             `json:"from"`
	Amount float64            `json:"amount"`
}

var TransactionsCollection *mongo.Collection

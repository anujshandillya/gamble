package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Statistics struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Wins    int64              `json:"wins"`
	Losses  int64              `json:"losses"`
	Wagered float64            `json:"wagered"`
}

var StatisticsCollection *mongo.Collection

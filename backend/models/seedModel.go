package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Seed struct {
	ID       bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Seed     string        `json:"seed"`
	SeedHash string        `json:"seedhash"`
}

var SeedCollection *mongo.Collection

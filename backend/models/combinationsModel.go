package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Combinations struct {
	ID             bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ServerSeed     string        `json:"serverseed"`
	ServerSeedHash string        `json:"serverseedhash"`
	ClientSeed     string        `json:"clientseed"`
	Nonce          uint16        `json:"nonce"`
}

var CombinationsCollection *mongo.Collection

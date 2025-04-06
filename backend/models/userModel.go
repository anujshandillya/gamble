package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type User struct {
	ID        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
}

var UserCollection *mongo.Collection

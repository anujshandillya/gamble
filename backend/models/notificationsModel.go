package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Notifications struct {
	ID      bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  bson.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Type    string        `json:"type"`
	Message string        `json:"message"`
}

var NotificationsCollection *mongo.Collection

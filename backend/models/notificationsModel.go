package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Notifications struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Type    string             `json:"type"`
	Message string             `json:"message"`
}

var NotificationsCollection *mongo.Collection

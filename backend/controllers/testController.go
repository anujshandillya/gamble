package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/anujshandillya/gambleserver/helpers"
	"github.com/anujshandillya/gambleserver/lib"
	"github.com/anujshandillya/gambleserver/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Seed struct {
	ID string `json:"id"`
}

// Seed
func GenerateServerSeed(w http.ResponseWriter, r *http.Request) {
	cursor := helpers.CreateNewSeed()
	var seed *models.Seed
	err := models.SeedCollection.FindOne(context.TODO(), bson.M{"_id": cursor.InsertedID}).Decode(&seed)
	lib.CheckErrorAndLog(err, "server seed seedController.go")

	// cookie := &http.Cookie{
	// 	Name:     "serverseed",
	// 	Value:    seed.Seed,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// 	Secure:   false,
	// 	Expires:  time.Now().Add(1 * time.Hour),
	// }

	// http.SetCookie(w, cookie)
}

func GetRandomSeed(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{"$sample", bson.D{{"size", 1}}}},
	}

	cursor, err := models.SeedCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var result bson.M
	if cursor.Next(ctx) {
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Random Document:", result)
	} else {
		log.Println("No document found")
	}

	json.NewEncoder(w).Encode(map[string]any{"message": "seed document", "document": result})
}

func DeleteSeed(userID string) any {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, _ := bson.ObjectIDFromHex(userID)
	filter := bson.M{"_id": id}
	cursor := models.SeedCollection.FindOneAndDelete(ctx, filter)
	if err := cursor.Err(); err != nil {
		log.Println("Error updating nonce:", err)
	}

	return cursor
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var x Seed
	json.NewDecoder(r.Body).Decode(&x)
	id := x.ID

	cursor := DeleteSeed(id)

	json.NewEncoder(w).Encode(map[string]any{"message": "deleted", "cursor": cursor})
}

// Nonce
func GetSeedByIDAndUpdate(userID string) any {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, err := bson.ObjectIDFromHex(userID)
	filter := bson.M{"_id": id}
	update := bson.M{"$inc": bson.M{"nonce": 1}}
	res := models.SeedCollection.FindOneAndUpdate(ctx, filter, update)
	if err != nil {
		log.Println("Error updating nonce:", err)
	}
	return res
}

func IncreaseNonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var x Seed
	json.NewDecoder(r.Body).Decode(&x)
	id := x.ID
	result := GetSeedByIDAndUpdate(id)
	json.NewEncoder(w).Encode(map[string]any{"id": id, "result": result})
}

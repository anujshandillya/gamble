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

type redisSchema struct {
	ServerSeed     string `json:"serverseed"`
	ServerSeedHash string `json:"serverseedhash"`
	ClientSeed     string `json:"clientseed"`
	Nonce          uint16 `json:"nonce"`
}

type redisScheme2 struct {
	Key string `json:"key"`
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
		{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
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

func GetRandomSeedCombination(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
	}

	cursor, err := models.CombinationsCollection.Aggregate(ctx, pipeline)
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

// redis
func SetValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var scheme redisSchema

	// Decode the incoming JSON request body into the scheme struct
	err := json.NewDecoder(r.Body).Decode(&scheme)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(scheme)
	if err != nil {
		http.Error(w, "Failed to serialize data", http.StatusInternalServerError)
		return
	}

	key := "activeSeeds:" + "aman"

	err = lib.RedisInstance.Set(lib.RedisCtx, key, jsonData, time.Hour).Err()
	if err != nil {
		http.Error(w, "Failed to set value in Redis", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Value set successfully"})
}

func GetValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var scheme redisScheme2

	err := json.NewDecoder(r.Body).Decode(&scheme)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	val, err := lib.RedisInstance.Get(lib.RedisCtx, scheme.Key).Result()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var actualjson redisSchema
	err = json.Unmarshal([]byte(val), &actualjson)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	json.NewEncoder(w).Encode(actualjson)
}

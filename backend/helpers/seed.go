package helpers

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"

	"github.com/anujshandillya/gambleserver/lib"
	"github.com/anujshandillya/gambleserver/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func generateNewSeed() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
}

func GetRandomSeed() models.Seed {
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

	var result models.Seed
	if cursor.Next(ctx) {
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Random Document:", result)
	} else {
		log.Println("No document found")
	}

	return result
}

func CreateNewSeed() *mongo.InsertOneResult {
	var seed models.Seed
	newseed := generateNewSeed()
	seedhash := sha256.Sum256([]byte(newseed))
	seed.Seed = newseed
	seed.SeedHash = hex.EncodeToString(seedhash[:])

	cursor, err := models.SeedCollection.InsertOne(context.TODO(), seed)

	lib.CheckErrorAndLog(err, "CreateNewSeed seed.go")

	return cursor
}

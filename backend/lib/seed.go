package lib

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/anujshandillya/gambleserver/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var AsciiSlice = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '[', ']', '{', '}', '|', ';', ':', '\'', ',', '.', '<', '>', '/', '?', '`', '~'}

func createClientSeed() string {
	AsciiSliceSize := len(AsciiSlice)
	clientSeed := ""

	// Generate a random string of 10 characters
	for range 10 {
		// Generate a secure random index
		randomIndexBig, err := rand.Int(rand.Reader, big.NewInt(int64(AsciiSliceSize)))
		if err != nil {
			panic("Failed to generate a random index: " + err.Error())
		}
		randomIndex := int(randomIndexBig.Int64())

		// Append the character to the seed
		clientSeed += string(AsciiSlice[randomIndex])
	}

	return clientSeed
}

func GetRandomSeed() (models.Seed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
	}

	cursor, err := models.SeedCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return models.Seed{}, err
	}
	defer cursor.Close(ctx)

	var result models.Seed
	if cursor.Next(ctx) {
		err := cursor.Decode(&result)
		if err != nil {
			return models.Seed{}, err
		}
		log.Println("Random Document:", result)
	} else {
		log.Println("No document found")
	}

	return result, nil
}

func GetRandomSeedCombination() (models.Combinations, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
	}

	cursor, err := models.CombinationsCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return models.Combinations{}, err
	}
	defer cursor.Close(ctx)

	var result models.Combinations
	if cursor.Next(ctx) {
		err := cursor.Decode(&result)
		if err != nil {
			return models.Combinations{}, err
		}
		log.Println("Random Document:", result)
	} else {
		return models.Combinations{}, fmt.Errorf("noDocumentFound")
	}

	return result, nil
}

func GetNewCombination() (models.Combinations, *mongo.InsertOneResult, error) {
	clientSeed := createClientSeed()

	randomSeed, err := GetRandomSeed()

	if err != nil {
		return models.Combinations{}, &mongo.InsertOneResult{}, err
	}

	var combination models.Combinations

	combination.ClientSeed = clientSeed
	combination.Nonce = 0
	combination.ServerSeed = randomSeed.Seed
	combination.ServerSeedHash = randomSeed.SeedHash

	cursor, err := models.CombinationsCollection.InsertOne(context.TODO(), combination)

	if err != nil {
		return models.Combinations{}, &mongo.InsertOneResult{}, err
	}

	return combination, cursor, nil
}

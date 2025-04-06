package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anujshandillya/gambleserver/lib"
	"github.com/anujshandillya/gambleserver/models"
	"github.com/anujshandillya/gambleserver/routes"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func init() {
	err := godotenv.Load()
	lib.CheckErrorAndLog(err, "main.go, init() line 19")
	mongoUri := os.Getenv("MONGO_URI")
	dbName := "gamble"
	clientOptions := options.Client().ApplyURI(mongoUri)

	client, err := mongo.Connect(clientOptions)

	lib.CheckErrorAndLog(err, "main.go, init() line 26")
	fmt.Println("Connected to MongoDB.")

	models.UserCollection = client.Database(dbName).Collection("user")
	models.BetsCollection = client.Database(dbName).Collection("bets")
	models.NotificationsCollection = client.Database(dbName).Collection("notifications")
	models.StatisticsCollection = client.Database(dbName).Collection("statistics")
	models.WalletCollection = client.Database(dbName).Collection("wallet")
	models.VaultCollection = client.Database(dbName).Collection("vault")
	models.TransactionsCollection = client.Database(dbName).Collection("transsactions")
}

func main() {
	r := routes.Router
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is running on port 4000.")
}

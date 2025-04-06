package helpers

import (
	"context"
	"fmt"

	"github.com/anujshandillya/gambleserver/lib"
	"github.com/anujshandillya/gambleserver/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func FindUserByEmail(email string) bool {
	err := models.UserCollection.FindOne(context.TODO(), bson.M{"email": email}).Err()
	return err == nil
}

func GetUserByEmail(email string) (models.User, bool) {
	if userExists := FindUserByEmail(email); !userExists {
		return models.User{}, false
	}
	var user models.User
	err := models.UserCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	fmt.Println(user)
	lib.CheckErrorAndLog(err, "auth.go GetUserByEmail()")
	return user, true
}

func CreateUser(user models.User) (*mongo.InsertOneResult, bool) {
	ifUser := FindUserByEmail(user.Email)
	if !ifUser {
		user.Password, _ = lib.HashPassword(user.Password)
		newUser, err := models.UserCollection.InsertOne(context.TODO(), user)
		lib.CheckErrorAndLog(err, "auth.go CreateUser()")
		return newUser, false
	}

	return nil, true
}

func CreateWallet(userId string) {
	wallet := models.Wallet{
		UserID:  userId,
		Balance: 0,
	}
	_, err := models.WalletCollection.InsertOne(context.TODO(), wallet)
	lib.CheckErrorAndLog(err, "auth.go CreateWallet()")
}
func CreateVault(userId string) {

	vault := models.Vault{
		UserID:  userId,
		Balance: 0,
	}
	_, err := models.VaultCollection.InsertOne(context.TODO(), vault)
	lib.CheckErrorAndLog(err, "auth.go CreateVault()")
}
func CreateStatistics(userId string) {
	stats := models.Statistics{
		UserID:  userId,
		Wins:    0,
		Losses:  0,
		Wagered: 0,
	}
	_, err := models.StatisticsCollection.InsertOne(context.TODO(), stats)
	lib.CheckErrorAndLog(err, "auth.go CreateStatistics()")
}

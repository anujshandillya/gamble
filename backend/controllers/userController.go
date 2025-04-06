package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anujshandillya/gambleserver/helpers"
	"github.com/anujshandillya/gambleserver/lib"
	"github.com/anujshandillya/gambleserver/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Register(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	lib.CheckErrorAndLog(err, "userController.go, Register() line 18")
	userCreated, userExists := helpers.CreateUser(user)
	if userExists {
		res.WriteHeader(http.StatusConflict)
		json.NewEncoder(res).Encode(map[string]string{"error": "User already exists"})
	} else if userCreated != nil {
		res.WriteHeader(http.StatusCreated)
		user_id := userCreated.InsertedID.(bson.ObjectID).Hex() // Convert bson.ObjectID to string
		fmt.Printf("User ID: %T\n", user_id)
		helpers.CreateWallet(user_id)
		helpers.CreateVault(user_id)
		helpers.CreateStatistics(user_id)
		json.NewEncoder(res).Encode(map[string]any{"message": "User created.", "user": userCreated})
	}
}

func Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	fmt.Println("User: ", user)
	lib.CheckErrorAndLog(err, "userController.go, Register() line 38")
	if user.Email == "" || user.Password == "" {
		res.WriteHeader(http.StatusForbidden)
		json.NewEncoder(res).Encode(map[string]string{"error": "Email and password are required"})
		return
	}
	userFound, userExists := helpers.GetUserByEmail(user.Email)
	if !userExists {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]string{"error": "User not found"})
		return
	}
	if !lib.CheckPasswordFromHash(user.Password, userFound.Password) {
		res.WriteHeader(http.StatusForbidden)
		json.NewEncoder(res).Encode(map[string]string{"error": "Invalid password"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]any{"message": "Login successful", "user": userFound})
}

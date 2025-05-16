package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/anujshandillya/gambleserver/helpers"
	"github.com/anujshandillya/gambleserver/lib"
	"github.com/anujshandillya/gambleserver/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	lib.CheckErrorAndLog(err, "userController.go, Register() line 18")
	userCreated, userExists := helpers.CreateUser(user)
	if userExists {
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User already exists"})
	} else if userCreated != nil {
		w.WriteHeader(http.StatusCreated)
		user_id := userCreated.InsertedID.(bson.ObjectID).Hex() // Convert bson.ObjectID to string
		fmt.Printf("User ID: %T\n", user_id)
		helpers.CreateWallet(user_id)
		helpers.CreateVault(user_id)
		helpers.CreateStatistics(user_id)
		helpers.CreateNewSeed()
		json.NewEncoder(w).Encode(map[string]any{"message": "User created.", "user": userCreated})
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("User: ", user)
	lib.CheckErrorAndLog(err, "userController.go, Register() line 38")
	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "Email and password are required"})
		return
	}
	userFound, userExists := helpers.GetUserByEmail(user.Email)
	if !userExists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}
	if !lib.CheckPasswordFromHash(user.Password, userFound.Password) {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid password"})
		return
	}
	token, _ := lib.GenerateJWT(userFound.ID.String())

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 48),
		HttpOnly: true,
		Path:     "/",
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"message": "Login successful", "user": userFound})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Path:     "/",
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}

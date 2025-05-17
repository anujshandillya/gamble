package routes

import (
	"encoding/json"
	"net/http"

	"github.com/anujshandillya/gambleserver/controllers"
	"github.com/anujshandillya/gambleserver/lib"
	"github.com/gorilla/mux"
)

func helloAPI(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("helloAPI")
}

var Router *mux.Router = mux.NewRouter()
var authRouter = Router.PathPrefix("/api/auth").Subrouter()
var testRouter = Router.PathPrefix("/api/test").Subrouter()

// var seedRouter = Router.PathPrefix("/api/seed").Subrouter()
var gameRouter = Router.PathPrefix("/api/game").Subrouter()

func init() {
	// API endpoint
	Router.HandleFunc("/api", lib.VerifyJWT(helloAPI)).Methods("GET")

	// Authentication routes
	authRouter.HandleFunc("/register", controllers.Register).Methods("POST")
	authRouter.HandleFunc("/login", controllers.Login).Methods("POST")
	authRouter.HandleFunc("/logout", controllers.Logout).Methods("POST")

	// Test routes
	testRouter.HandleFunc("/newseed", lib.VerifyJWT(controllers.GenerateServerSeed)).Methods("GET")
	testRouter.HandleFunc("/getrandomseed", lib.VerifyJWT(controllers.GetRandomSeed)).Methods("GET")
	testRouter.HandleFunc("/increaseNonce", lib.VerifyJWT(controllers.IncreaseNonce)).Methods("POST")
	testRouter.HandleFunc("/deleteSeed", lib.VerifyJWT(controllers.Delete)).Methods("DELETE")

	testRouter.HandleFunc("/redis/set", lib.VerifyJWT(controllers.SetValue)).Methods("POST")
	testRouter.HandleFunc("/redis/get", lib.VerifyJWT(controllers.GetValue)).Methods("POST")

	// Game routes
	gameRouter.HandleFunc("/limbo", lib.VerifyJWT(controllers.Limbo)).Methods("POST")
	gameRouter.HandleFunc("/cointoss", lib.VerifyJWT(controllers.CoinToss)).Methods("POST")
	gameRouter.HandleFunc("/dragontower", lib.VerifyJWT(controllers.DragonTower)).Methods("POST")
	gameRouter.HandleFunc("/dice", lib.VerifyJWT(controllers.Dice)).Methods("POST")
	gameRouter.HandleFunc("/slides", lib.VerifyJWT(controllers.Slides)).Methods("POST")
	gameRouter.HandleFunc("/mines", lib.VerifyJWT(controllers.Mines)).Methods("POST")
	gameRouter.HandleFunc("/wheel", lib.VerifyJWT(controllers.Wheel)).Methods("POST")
	gameRouter.HandleFunc("/highlow", lib.VerifyJWT(controllers.HighLow)).Methods("POST")
	gameRouter.HandleFunc("/pump", lib.VerifyJWT(controllers.Pump)).Methods("POST")
	gameRouter.HandleFunc("/bj", lib.VerifyJWT(controllers.BJ)).Methods("POST")

	// Next routes
	gameRouter.HandleFunc("/cointoss/next", lib.VerifyJWT(controllers.CoinTossNext)).Methods("POST")
	gameRouter.HandleFunc("/dragontower/next", lib.VerifyJWT(controllers.DragonTowerNext)).Methods("POST")
	gameRouter.HandleFunc("/mines/next", lib.VerifyJWT(controllers.MinesNext)).Methods("POST")
	gameRouter.HandleFunc("/highlow/next", lib.VerifyJWT(controllers.HighLowNext)).Methods("POST")
	gameRouter.HandleFunc("/pump/next", lib.VerifyJWT(controllers.PumpNext)).Methods("POST")
	gameRouter.HandleFunc("/bj/next", lib.VerifyJWT(controllers.BJNext)).Methods("POST")

}

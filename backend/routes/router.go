package routes

import (
	"encoding/json"
	"net/http"

	"github.com/anujshandillya/gambleserver/controllers"
	"github.com/anujshandillya/gambleserver/middlewares"
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
	Router.HandleFunc("/api", middlewares.VerifyJWT(helloAPI)).Methods("GET")

	// Authentication routes
	authRouter.HandleFunc("/register", controllers.Register).Methods("POST")
	authRouter.HandleFunc("/login", controllers.Login).Methods("POST")
	authRouter.HandleFunc("/logout", controllers.Logout).Methods("POST")

	// Test routes
	testRouter.HandleFunc("/newseed", middlewares.VerifyJWT(controllers.GenerateServerSeed)).Methods("GET")
	testRouter.HandleFunc("/getrandomseed", middlewares.VerifyJWT(controllers.GetRandomSeed)).Methods("GET")
	testRouter.HandleFunc("/increaseNonce", middlewares.VerifyJWT(controllers.IncreaseNonce)).Methods("POST")
	testRouter.HandleFunc("/deleteSeed", middlewares.VerifyJWT(controllers.Delete)).Methods("DELETE")

	testRouter.HandleFunc("/redis/set", middlewares.VerifyJWT(controllers.SetValue)).Methods("POST")
	testRouter.HandleFunc("/redis/get", middlewares.VerifyJWT(controllers.GetValue)).Methods("POST")

	// Game routes
	gameRouter.HandleFunc("/limbo", middlewares.VerifyJWT(controllers.Limbo)).Methods("POST")
	gameRouter.HandleFunc("/dice", middlewares.VerifyJWT(controllers.Dice)).Methods("POST")
	gameRouter.HandleFunc("/wheel", middlewares.VerifyJWT(controllers.Wheel)).Methods("POST")
	gameRouter.HandleFunc("/slides", middlewares.VerifyJWT(controllers.Slides)).Methods("POST")

	gameRouter.HandleFunc("/cointoss", middlewares.VerifyJWT(controllers.CoinToss)).Methods("POST")
	gameRouter.HandleFunc("/dragontower", middlewares.VerifyJWT(controllers.DragonTower)).Methods("POST")
	gameRouter.HandleFunc("/mines", middlewares.VerifyJWT(controllers.Mines)).Methods("POST")
	gameRouter.HandleFunc("/highlow", middlewares.VerifyJWT(controllers.HighLow)).Methods("POST")
	gameRouter.HandleFunc("/pump", middlewares.VerifyJWT(controllers.Pump)).Methods("POST")
	gameRouter.HandleFunc("/bj", middlewares.VerifyJWT(controllers.BJ)).Methods("POST")

	// Next routes
	gameRouter.HandleFunc("/cointoss/next", middlewares.VerifyJWT(controllers.CoinTossNext)).Methods("POST")
	gameRouter.HandleFunc("/dragontower/next", middlewares.VerifyJWT(controllers.DragonTowerNext)).Methods("POST")
	gameRouter.HandleFunc("/mines/next", middlewares.VerifyJWT(controllers.MinesNext)).Methods("POST")
	gameRouter.HandleFunc("/highlow/next", middlewares.VerifyJWT(controllers.HighLowNext)).Methods("POST")
	gameRouter.HandleFunc("/pump/next", middlewares.VerifyJWT(controllers.PumpNext)).Methods("POST")
	gameRouter.HandleFunc("/bj/next", middlewares.VerifyJWT(controllers.BJNext)).Methods("POST")

}

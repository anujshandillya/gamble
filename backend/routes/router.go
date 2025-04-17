package routes

import (
	"encoding/json"
	"net/http"

	"github.com/anujshandillya/gambleserver/controllers"
	"github.com/gorilla/mux"
)

func helloAPI(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("helloAPI")
}

var Router *mux.Router = mux.NewRouter()
var authRouter = Router.PathPrefix("/api/auth").Subrouter()
var gameRouter = Router.PathPrefix("/api/game").Subrouter()

func init() {
	// API endpoint
	Router.HandleFunc("/api", helloAPI).Methods("GET")

	// Authentication routes
	authRouter.HandleFunc("/register", controllers.Register).Methods("POST")
	authRouter.HandleFunc("/login", controllers.Login).Methods("POST")

	// Game routes
	gameRouter.HandleFunc("/limbo", controllers.Limbo).Methods("POST")
	gameRouter.HandleFunc("/cointoss", controllers.CoinToss).Methods("POST")
	gameRouter.HandleFunc("/dragontower", controllers.DragonTower).Methods("POST")
	gameRouter.HandleFunc("/dice", controllers.Dice).Methods("POST")
	gameRouter.HandleFunc("/slides", controllers.Slides).Methods("POST")
	gameRouter.HandleFunc("/mines", controllers.Mines).Methods("POST")
	gameRouter.HandleFunc("/wheel", controllers.Wheel).Methods("POST")
	gameRouter.HandleFunc("/highlow", controllers.HighLow).Methods("POST")
	gameRouter.HandleFunc("/pump", controllers.Pump).Methods("POST")
	gameRouter.HandleFunc("/bj", controllers.BJ).Methods("POST")
}

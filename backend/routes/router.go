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

func init() {
	Router.HandleFunc("/api", helloAPI).Methods("GET")
	authRouter.HandleFunc("/register", controllers.Register).Methods("POST")
	authRouter.HandleFunc("/login", controllers.Login).Methods("POST")
}

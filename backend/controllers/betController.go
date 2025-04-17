package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anujshandillya/gambleserver/lib"
	"github.com/anujshandillya/gambleserver/types"
)

func Limbo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var bet types.BetLimbo
	err := json.NewDecoder(req.Body).Decode(&bet)
	lib.CheckErrorAndLog(err, "betController.go, Limbo() line 16")
	fmt.Println("betLimbo:", bet)
}

func CoinToss(res http.ResponseWriter, req *http.Request) {}

func DragonTower(res http.ResponseWriter, req *http.Request) {}

func Dice(res http.ResponseWriter, req *http.Request) {}

func Slides(res http.ResponseWriter, req *http.Request) {}

func Mines(res http.ResponseWriter, req *http.Request) {}

func Wheel(res http.ResponseWriter, req *http.Request) {}

func HighLow(res http.ResponseWriter, req *http.Request) {}

func Pump(res http.ResponseWriter, req *http.Request) {}

func BJ(res http.ResponseWriter, req *http.Request) {}

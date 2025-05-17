package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/anujshandillya/gambleserver/lib"
	"github.com/anujshandillya/gambleserver/types"
)

func Limbo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bet types.BetLimbo
	var outcome string
	var payout float64
	var profit float64
	err := json.NewDecoder(r.Body).Decode(&bet)
	lib.CheckErrorAndLog(err, "betController.go, Limbo() line 20")
	email, err := r.Cookie("email")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]any{"message": "unauthorized", "status": 401})
		return
	}

	userEmail := email.Value
	combinationStr := lib.GetAndSetRedisSeed(userEmail)
	combinationJson := lib.UnMarshalRedisSeed(combinationStr)
	wager := bet.Amount
	currency := bet.Currency

	f, inputHash, hexStr := lib.RandomFloat(combinationJson.ServerSeed, combinationJson.ClientSeed, int(combinationJson.Nonce))
	fmt.Println("outcome:", 1/(1-f))
	result := math.Round((1/(1-f))*100) / 100

	if result > bet.Multi {
		outcome = "win"
	} else {
		outcome = "lose"
	}

	if outcome != "win" {
		payout = 0
	} else {
		payout = bet.Multi * wager
	}

	if outcome != "lose" {
		profit = 0
	} else {
		profit = payout - wager
	}

	lib.IncreaseNonce(userEmail)

	jsonVerificationData := types.VerificationData{
		HashInput:      inputHash,
		Hash:           hexStr,
		UsedServerSeed: &combinationJson.ServerSeed,
	}

	jsonResponse := types.BetResultResponse{
		Result:       result,
		Currency:     currency,
		Outcome:      outcome,
		Payout:       payout,
		Wager:        wager,
		Profit:       profit,
		Nonce:        int(combinationJson.Nonce),
		ClientSeed:   combinationJson.ClientSeed,
		Verification: jsonVerificationData,
	}
	json.NewEncoder(w).Encode(jsonResponse)
}

func CoinToss(w http.ResponseWriter, r *http.Request) {}

func DragonTower(w http.ResponseWriter, r *http.Request) {}

func Dice(w http.ResponseWriter, r *http.Request) {}

func Slides(w http.ResponseWriter, r *http.Request) {}

func Mines(w http.ResponseWriter, r *http.Request) {}

func Wheel(w http.ResponseWriter, r *http.Request) {}

func HighLow(w http.ResponseWriter, r *http.Request) {}

func Pump(w http.ResponseWriter, r *http.Request) {}

func BJ(w http.ResponseWriter, r *http.Request) {}

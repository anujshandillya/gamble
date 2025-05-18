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

	jsonResponse := types.BetResultResponseLimbo{
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

func Dice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bet types.BetDice
	var outcome string
	var payout float64
	var profit float64
	err := json.NewDecoder(r.Body).Decode(&bet)
	lib.CheckErrorAndLog(err, "betController.go, Dice() line 85")

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
	value := bet.Value

	f, inputHash, hexStr := lib.RandomFloat(combinationJson.ServerSeed, combinationJson.ClientSeed, int(combinationJson.Nonce))
	fmt.Println("outcome:", math.Round(f*10000)/100)
	result := math.Round(f*10000) / 100

	if bet.OverUnder == "over" {
		if result >= value {
			outcome = "win"
		} else {
			outcome = "lose"
		}
	} else if bet.OverUnder == "under" {
		if result <= value {
			outcome = "win"
		} else {
			outcome = "lose"
		}
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

	jsonResponse := types.BetResultResponseDice{
		Result:       result,
		Currency:     currency,
		Outcome:      outcome,
		OverUnder:    bet.OverUnder,
		Payout:       payout,
		Wager:        wager,
		Profit:       profit,
		Nonce:        int(combinationJson.Nonce),
		ClientSeed:   combinationJson.ClientSeed,
		Verification: jsonVerificationData,
	}
	json.NewEncoder(w).Encode(jsonResponse)
}

func Slides(w http.ResponseWriter, r *http.Request) {}

func Wheel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bet types.BetWheels
	var outcome string
	var payout float64
	var profit float64
	err := json.NewDecoder(r.Body).Decode(&bet)
	lib.CheckErrorAndLog(err, "betController.go, Dice() line 85")

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
	risk := bet.Risk
	segments := bet.Segments

	f, inputHash, hexStr := lib.RandomFloat(combinationJson.ServerSeed, combinationJson.ClientSeed, int(combinationJson.Nonce))

	index := math.Round(f * float64(segments-1))
	result := lib.WheelResult(risk, segments, index)

	if result == 0 {
		outcome = "lose"
	} else {
		outcome = "win"
	}

	if outcome != "win" {
		payout = 0
	} else {
		payout = result * wager
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

	jsonResponse := types.BetResultResponseWheel{
		Result:       result,
		Index:        int(index + 1),
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

func Mines(w http.ResponseWriter, r *http.Request) {}

func HighLow(w http.ResponseWriter, r *http.Request) {}

func Pump(w http.ResponseWriter, r *http.Request) {}

func BJ(w http.ResponseWriter, r *http.Request) {}

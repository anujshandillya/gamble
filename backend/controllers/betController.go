package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anujshandillya/gambleserver/lib"
	"github.com/anujshandillya/gambleserver/models"
	"github.com/anujshandillya/gambleserver/types"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type verificationData struct {
	HashInput      string  `json:"hashInput"`
	Hash           string  `json:"hash"`
	UsedServerSeed *string `json:"usedServerSeed"`
}

type betResultResponse struct {
	Currency       string           `json:"currency"`
	Result         float64          `json:"result"`
	Outcome        string           `json:"outcome"`
	Payout         float64          `json:"payout"`
	Wager          float64          `json:"wager"`
	Profit         float64          `json:"profit"`
	Nonce          int              `json:"nonce"`
	ClientSeed     string           `json:"clientSeed"`
	ServerSeedHash string           `json:"serverSeedHash"`
	Verification   verificationData `json:"verification"`
}

func Limbo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bet types.BetLimbo
	var outcome string
	var payout float64
	var profit float64
	// var nonce int16
	err := json.NewDecoder(r.Body).Decode(&bet)
	lib.CheckErrorAndLog(err, "betController.go, Limbo() line 34")
	var seed *models.Seed
	clientSeed := bet.Seed
	wager := bet.Amount
	currency := bet.Currency
	cookie, err := r.Cookie("serverseed")
	if err != nil {
		http.Error(w, "Cookie not found", http.StatusBadRequest)
		return
	}
	err = models.SeedCollection.FindOne(context.TODO(), bson.M{"seed": cookie.Value}).Decode(&seed)
	fmt.Println(cookie.Value)
	f, inputHash, hexStr := lib.RandomFloat(seed.Seed, bet.Seed, 1)
	fmt.Println("outcome:", 1/(1-f))
	result := 1 / (1 - f)

	if result > bet.Multi {
		outcome = "lose"
	} else {
		outcome = "win"
	}

	if outcome != "lose" {
		payout = 0
	} else {
		payout = bet.Multi * wager
	}

	if outcome != "lose" {
		profit = 0
	} else {
		profit = payout - wager
	}

	jsonVerificationData := verificationData{
		HashInput:      inputHash,
		Hash:           hexStr,
		UsedServerSeed: &seed.Seed,
	}

	jsonResponse := betResultResponse{
		Result:       result,
		Currency:     currency,
		Outcome:      outcome,
		Payout:       payout,
		Wager:        wager,
		Profit:       profit,
		Nonce:        27,
		ClientSeed:   clientSeed,
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

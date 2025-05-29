package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/anujshandillya/gambleserver/lib"
	"github.com/anujshandillya/gambleserver/types"
)

func CoinTossNext(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cointoss types.NextCoinToss
	json.NewDecoder(r.Body).Decode(&cointoss)

	selection := cointoss.Selection
	round := cointoss.Round

	email, err := r.Cookie("email")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]any{"message": "unauthorized", "status": 401})
		return
	}

	userEmail := email.Value

	key := fmt.Sprintf("activeBet:%s:%s", userEmail, "cointoss")

	bet := lib.RedisInstance.Get(lib.RedisCtx, key)
	err = bet.Err()

	if err != nil {
		http.Error(w, "No active bet found", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{"message": "no active bet found", "status": 404})
		return
	}

	var activeBet types.ActiveBetCoinToss
	err = json.Unmarshal([]byte(bet.Val()), &activeBet)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(map[string]any{"message": "error marshaling the data", "status": 417, "error": fmt.Errorf("failed to unmarshal bet: %w", err)})
		return
	}

	state := activeBet.State

	combinationStr := lib.GetAndSetRedisSeed(userEmail)
	combinationJson := lib.UnMarshalRedisSeed(combinationStr)
	f, _, _ := lib.RandomFloat(combinationJson.ServerSeed, combinationJson.ClientSeed, int(round))

	fmt.Println("outcome:", f)
	outcomeInt := math.Floor(f * 2)
	var outcome []int
	if outcomeInt == 0 {
		outcome = append(outcome, 0)
	} else {
		outcome = append(outcome, 1)
	}

	state = append(state, outcome)

	activeBet.State = state
	var result string
	if selection == outcome[0] {
		result = "win"
		key := fmt.Sprintf("activeBet:%s:%s", userEmail, "cointoss")

		jsonData, err := json.Marshal(activeBet)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusExpectationFailed)
			json.NewEncoder(w).Encode(map[string]any{"message": "error marshaling the data", "status": 417, "error": fmt.Errorf("failed to marshal data: %w", err)})
			return
		}

		err = lib.RedisInstance.Set(lib.RedisCtx, key, jsonData, 0).Err()

		if err != nil {
			http.Error(w, "Failed to store bet", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{"message": "failed to store bet", "status": 500})
			return
		}
	} else {
		result = "lose"
		err = lib.DeleteRedisBet(userEmail, "cointoss")
		if err != nil {
			http.Error(w, "Failed to delete bet", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{"message": "failed to delete bet", "status": 500})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]any{
		"result":    result,
		"multi":     1.96 * math.Pow(2, float64(round-1)),
		"state":     state,
		"selection": selection,
		"outcome":   outcome,
		"round":     len(state),
	})
}

func DragonTowerNext(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dragontower types.NextDragonTower
	json.NewDecoder(r.Body).Decode(&dragontower)

	egg := dragontower.Egg
	level := dragontower.Level

	email, err := r.Cookie("email")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]any{"message": "unauthorized", "status": 401})
		return
	}

	userEmail := email.Value

	key := fmt.Sprintf("activeBet:%s:%s", userEmail, "dragontower")

	bet := lib.RedisInstance.Get(lib.RedisCtx, key)
	err = bet.Err()

	if err != nil {
		http.Error(w, "No active bet found", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{"message": "no active bet found", "status": 404})
		return
	}

	var activeBet types.ActiveBetDragonTower
	err = json.Unmarshal([]byte(bet.Val()), &activeBet)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(map[string]any{"message": "error marshaling the data", "status": 417, "error": fmt.Errorf("failed to unmarshal bet: %w", err)})
		return
	}

	state := activeBet.State
	levelSet := activeBet.LevelSet

	var outcome string
	if levelSet[level][egg] == 1 {
		outcome = "win"
		state = append(state, levelSet[level])
		activeBet.State = state

		key := fmt.Sprintf("activeBet:%s:%s", userEmail, "dragontower")

		jsonData, err := json.Marshal(activeBet)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusExpectationFailed)
			json.NewEncoder(w).Encode(map[string]any{"message": "error marshaling the data", "status": 417, "error": fmt.Errorf("failed to marshal data: %w", err)})
			return
		}

		err = lib.RedisInstance.Set(lib.RedisCtx, key, jsonData, 0).Err()

		if err != nil {
			http.Error(w, "Failed to store bet", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{"message": "failed to store bet", "status": 500})
			return
		}
	} else {
		outcome = "lose"
		err := lib.DeleteRedisBet(userEmail, "dragontower")
		if err != nil {
			http.Error(w, "Failed to delete bet", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{"message": "failed to delete bet", "status": 500})
			return
		}
		json.NewEncoder(w).Encode(map[string]any{
			"result":     outcome,
			"multi":      0,
			"state":      activeBet.LevelSet,
			"egg":        egg,
			"level":      level,
			"amount":     activeBet.Amount,
			"difficulty": activeBet.Difficulty,
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]any{
		"result":     outcome,
		"multi":      1.96 * math.Pow(2, float64(level)),
		"state":      state,
		"egg":        egg,
		"level":      level,
		"amount":     activeBet.Amount,
		"difficulty": activeBet.Difficulty,
	})
}

func MinesNext(w http.ResponseWriter, r *http.Request) {}

func HighLowNext(w http.ResponseWriter, r *http.Request) {}

func PumpNext(w http.ResponseWriter, r *http.Request) {}

func BJNext(w http.ResponseWriter, r *http.Request) {}

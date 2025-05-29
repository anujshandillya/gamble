package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/anujshandillya/gambleserver/models"
	"github.com/redis/go-redis/v9"
)

var redisDB int

func init() {
	dbStr := os.Getenv("REDIS_DB")
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		db = 0 // default to 0 if conversion fails
	}
	redisDB = db
}

var RedisInstance = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_DB_URI"),
	Password: os.Getenv("REDIS_DB_PASSWORD"),
	DB:       redisDB,
})

var RedisCtx = context.Background()

func GetAndSetRedisSeed(userEmail string) string {
	key := "activeSeeds:" + userEmail
	_, err := RedisInstance.Get(RedisCtx, key).Result()

	if err != nil {
		combination, err := GetRandomSeedCombination()
		if err != nil {
			combination, _, _ = GetNewCombination()
		}
		fmt.Println("combination", combination)
		jsonData, _ := json.Marshal(combination)
		RedisInstance.Set(RedisCtx, key, jsonData, time.Hour*24)
	}

	value, _ := RedisInstance.Get(RedisCtx, key).Result()

	return value
}

func UnMarshalRedisSeed(seedStr string) models.Combinations {
	var combination models.Combinations

	json.Unmarshal([]byte(seedStr), &combination)

	return combination
}

func IncreaseNonce(userEmail string) error {
	key := "activeSeeds:" + userEmail
	ttl, err := RedisInstance.TTL(RedisCtx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to get TTL")
	}
	value, _ := RedisInstance.Get(RedisCtx, key).Result()
	newValue := UnMarshalRedisSeed(value)
	newValue.Nonce += 1
	jsonData, _ := json.Marshal(newValue)
	err = RedisInstance.Set(RedisCtx, key, jsonData, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to update value")
	}

	if ttl > 0 {
		err = RedisInstance.Expire(RedisCtx, key, ttl).Err()
		if err != nil {
			return fmt.Errorf("failed to reapply TTL")
		}
	}

	return nil
}

func DeleteRedisBet(userEmail, game string) error {
	key := fmt.Sprintf("activeBet:%s:%s", userEmail, game)

	err := RedisInstance.Del(RedisCtx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete existing bet: %w", err)
	}

	return nil
}

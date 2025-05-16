package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func CheckErrorAndLog(err error, function string) {
	if err != nil {
		log.Fatal(err, " in ", function)
	}
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPasswordFromHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RandomFloat(serverSeed, clientSeed string, nonce int) (float64, string, string) {
	input := fmt.Sprintf("%s:%s:%d", serverSeed, clientSeed, nonce)
	hash := sha256.Sum256([]byte(input))

	hexStr := hex.EncodeToString(hash[:4])

	intVal, err := strconv.ParseUint(hexStr, 16, 64)
	CheckErrorAndLog(err, "could not handle hashing the seed")

	return float64(intVal) / float64(0xffffffff), input, hexStr
}

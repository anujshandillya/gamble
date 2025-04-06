package lib

import (
	"fmt"
	"log"

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
	fmt.Println("Hash: ", hash, " Password: ", password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

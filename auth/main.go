package auth

import (
	"github.com/alexedwards/argon2id"
	"log"
	"restapi/env"
)

func HashPassword(password string) string {
	secretKey := env.Get[string]("SECRET_KEY")
	salted := secretKey + password
	hash, err := argon2id.CreateHash(salted, argon2id.DefaultParams)
	if err != nil {
		panic(err)
	}

	return hash
}

func VerifyPassword(password, hash string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Println(err)
	}

	return match
}

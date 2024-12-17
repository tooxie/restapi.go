package db

import (
	"log"

	"github.com/google/uuid"
)

func GetUserByEmail(email string) (*User, error) {
	var user User
	res := Conn.Where("email = ?", email).First(&user)
	if res.Error != nil {
		log.Println("User not found:", email)
		return nil, res.Error
	}

	return &user, nil
}

func GetUserById(id uuid.UUID) (*User, error) {
	var user User
	res := Conn.Where("id = ?", id).First(&user)
	if res.Error != nil {
		log.Println("User not found:", id)
		return nil, res.Error
	}

	return &user, nil
}

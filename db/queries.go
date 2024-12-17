package db

import "fmt"

func GetRole(name string) *Role {
	var role Role
	res := Conn.Where("name = ?", name).First(&role)
	if res.Error != nil {
		panic(fmt.Sprintf("Role '%s' not found", name))
	}

	return &role
}

func getUser(email string, password string) *User {
	var user User

	query := Conn.Where("email = ?", email).First(&user)
	if query.Error != nil {
		return nil
	}

	return &user
}

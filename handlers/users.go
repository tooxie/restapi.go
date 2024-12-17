package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"restapi/auth"
	"restapi/db"
	"restapi/env"
	"restapi/errors"
)

func getCurrentlyLoggedInUser(c *gin.Context) *db.User {
	if c.Keys["user"] == nil {
		return nil
	}

	user := c.Keys["user"].(*db.User)
	log.Println("user:", user)
	return user
}

func getUser(email string, password string) *db.User {
	user, err := db.GetUserByEmail(email)
	if err != nil {
		log.Println("err:", err)
		return nil
	}

	secretKey := env.Get[string]("SECRET_KEY")
	salted := secretKey + password
	validCredentials := auth.VerifyPassword(salted, user.Hash)
	if !validCredentials {
		return nil
	}

	return user
}

func Login(c *gin.Context) {
	secretKey := env.Get[string]("SECRET_KEY")

	email := c.PostForm("email")
	password := c.PostForm("password")

	user := getUser(email, password)
	if user == nil {
		errors.Raise(c, errors.InvalidCredentials)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id,
		"exp":    time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		errors.Raise(c, errors.JwtTokenGenerationFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func userExists(email string) bool {
	user, _ := db.GetUserByEmail(email)
	return user != nil
}

func Signup(c *gin.Context) {
	email := c.PostForm("email")
	if userExists(email) {
		errors.Raise(c, errors.EmailAlreadyInUse)
		return
	}

	password := c.PostForm("password")
	hash := auth.HashPassword(password)

	role := db.GetRole("user")
	roles := []db.Role{*role}
	user := db.User{Email: email, Hash: hash, Roles: roles}
	db.Conn.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUser(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"data": user})
}

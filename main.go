package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"restapi/db"
	"restapi/env"
	h "restapi/handlers"
	m "restapi/middlewares"
	// "restapi/services"
)

type Some struct {
	One string
	Two int
}

type Environment struct {
	HOSTNAME       env.IPv4
	SECRET_KEY     env.String
	ENABLE_METRICS env.Bool     `env:"optional"`
	BLACKLIST      []env.String `env:"separator=' ';optional"`
	GIN_MODE       env.String   `env:"optional;default='debug'"`
	PORT           env.Int

	// Validate JSON against certain schema?
	// SOME_JSON      env.JSON(Some)
}

func registerEnvVars() {
	var vars Environment
	// err := env.Assert(vars)
	// if err != nil {
	// 	panic(err)
	// }
	missing, invalid := env.Validate(vars)
	if missing != nil {
		panic(fmt.Sprintf("Missing env vars: %v", missing))
	}
	if invalid != nil {
		panic(fmt.Sprintf("Invalid env vars: %v", invalid))
	}

	fmt.Println("ENABLE_METRICS", env.Get[env.Bool]("ENABLE_METRICS"))

	log.Println("BLACKLIST", env.Get[[]interface{}]("BLACKLIST"))
	for _, user := range env.Get[[]interface{}]("BLACKLIST") {
		if strings.Contains(user.(string), "bob") {
			log.Println("Bob is blacklisted")
		}
	}
}

// func registerServices() {
// 	services.Register("https://itsthisforthat.com/", "itft")
// 	itft := services.Get("itft").Path("api.php?json")
// }

func main() {
	registerEnvVars()
	// registerServices()

	db.Connect()

	r := gin.Default()
	// r.Use(m.Sec)

	// Auth-free
	r.POST("/signup", h.Signup)
	r.POST("/login", h.Login)
	r.GET("/books", h.ListBooks)
	r.GET("/books/:id", h.GetBook)

	// Auth-only
	r.Use(m.Auth)

	r.GET("/me", h.GetUser)
	r.POST("/books", h.CreateBook)
	r.PATCH("/books/:id", h.UpdateBook)
	r.DELETE("/books/:id", h.DeleteBook)

	// Admin
	r.GET("/admin", m.Admin, h.Admin)

	// 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Resource not found"})
	})

	r.Run()
}

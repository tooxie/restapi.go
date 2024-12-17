package middlewares

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"restapi/errors"
)

func Admin(c *gin.Context) {
	log.Println("middlewares.Admin()")
	claims := c.MustGet("claims").(jwt.MapClaims)
	role := claims["role"].(string)

	if role != "admin" {
		errors.Abort(c, errors.Forbidden)
		return
	}

	c.Next()
}

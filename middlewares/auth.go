package middlewares

import (
	e "errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"restapi/db"
	"restapi/env"
	"restapi/errors"
)

func Auth(c *gin.Context) {
	log.Println("-- middlewares.auth.Auth()")
	secretKey := env.Get[string]("SECRET_KEY")
	authToken := c.GetHeader("Authorization")

	token, err := jwt.Parse(
		authToken,
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, http.ErrAbortHandler
			}

			return []byte(secretKey), nil
		},
	)

	if err != nil {
		if e.Is(err, jwt.ErrTokenExpired) {
			errors.Abort(c, errors.JwtTokenExpired)
		} else {
			errors.AbortWithMessage(c, errors.Unauthorized, err.Error())
		}
		return
	}

	if !token.Valid {
		errors.Abort(c, errors.JwtInvalidToken)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		errors.Abort(c, errors.Unauthorized)
		return
	}

	c.Set("claims", claims)
	if claims["userId"] == nil {
		errors.Abort(c, errors.JwtInvalidToken)
		return
	}
	userId := uuid.MustParse(claims["userId"].(string))
	user, err := db.GetUserById(userId)
	if err != nil {
		errors.Abort(c, errors.Unauthorized)
		return
	}
	c.Set("user", user)
	c.Next()
}

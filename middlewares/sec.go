package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"restapi/env"
	"restapi/errors"
)

func Sec(c *gin.Context) {
	log.Println("-- middlewares.sec.Sec()")
	hostname := env.Get[string]("HOSTNAME")
	if hostname == "" {
		log.Println("Error: $HOSTNAME environment variable not set")
		errors.Raise(c, errors.InternalServerError)
		c.Abort()
		return
	}

	if c.Request.Host != hostname {
		errors.Raise(c, errors.InvalidHostHeader)
		c.Abort()
		return
	}

	// https://gin-gonic.com/docs/examples/security-headers/
	c.Header("X-Frame-Options", "DENY")
	c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
	c.Header("X-XSS-Protection", "1; mode=block")
	c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	c.Header("Referrer-Policy", "strict-origin")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
	c.Next()
}

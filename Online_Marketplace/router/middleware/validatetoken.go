package middleware

import (
	"errors"
	"net/http"
	h "onlinemarketplace/helper"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authtoken() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := h.NewResponse()
		tokenString := strings.Replace(c.Request.Header["Authorization"][0], "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(h.SecretKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, res.ErrorResponse(err))
			c.Abort()
			return
		}
		// Check if the token is valid
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, res.ErrorResponse(errors.New("Invalid token")))
			c.Abort()
			return
		}
		c.Next()
	}
}

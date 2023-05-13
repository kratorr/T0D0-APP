package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (h *Handler) UserIdentification(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
		return
	}

	splitedHeader := strings.Split(authHeader, " ")
	if len(splitedHeader) <= 1 {
		fmt.Println("auth header too short")
		return
	}
	token := splitedHeader[1]

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.secretKey), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "jwt parse error"})
	}

	c.Set("userID", claims["sub"])
	c.Set("userLogin", claims["nickname"])
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserIdentification(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
		return
	}

	user, err := h.services.Auth.GetUserByToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "auth service error"})
	}

	c.Set("userID", user.ID)
	c.Set("userLogin", user.Login)
}

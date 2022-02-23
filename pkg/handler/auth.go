package handler

import (
	"net/http"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var r models.User
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// TODO add field Password2 and compare with Password1
	err := h.services.Auth.SignUp(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"response": "user created"})
}

func (h *Handler) SignIn(c *gin.Context) {
	var r models.User
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.services.Auth.SignIn(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Token": token})
}

func (h *Handler) Test(c *gin.Context) {
	c.JSON(200, gin.H{"test": "test"})
}

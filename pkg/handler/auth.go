package handler

import (
	"net/http"

	"todo/models"

	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @Summary      SignUp new user
// @Description  SignUp new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "Add account"
// @Router       /auth/signup/ [post]
// @Success      200 {string}  string    "ok"
func (h *Handler) SignUp(c *gin.Context) {
	var userDto models.CreateUserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.services.Auth.SignUp(userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"response": "user created"})
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

// TODO SignOut epta!!!

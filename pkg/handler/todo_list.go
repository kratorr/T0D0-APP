package handler

import (
	"net/http"

	"todo/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTodoList(c *gin.Context) {
	// var r models.User
	// if err := c.ShouldBindJSON(&r); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// todoList, err := h.services.Auth.SignUp(r)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"response": "Todo list created"})
}

func (h *Handler) getTodoList(c *gin.Context) {
	var r models.User
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.services.Auth.SignUp(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "user created"})
}

func (h *Handler) deleteTodoList(c *gin.Context) {
	var r models.User
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.services.Auth.SignUp(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "user created"})
}

func (h *Handler) updateTodoList(c *gin.Context) {
	var r models.User
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.services.Auth.SignUp(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "user created"})
}

func (h *Handler) getAllTodoLists(c *gin.Context) {
	var r models.User
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.services.Auth.SignUp(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "user created"})
}

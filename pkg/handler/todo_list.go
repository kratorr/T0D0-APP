package handler

import (
	"fmt"
	"net/http"

	"todo/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) createTodoList(c *gin.Context) {
	userID := c.Value("userID").(int)

	inputData := models.TodoList{}
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		fmt.Println(err)
	}
	zap.L().Sugar().Debug("create TODO list ", userID, inputData)
	todoListID, err := h.services.TodoList.Create(userID, inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"ID": todoListID})
}

func (h *Handler) getTodoList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": "user created"})
}

func (h *Handler) deleteTodoList(c *gin.Context) {
}

func (h *Handler) updateTodoList(c *gin.Context) {
}

func (h *Handler) getAllTodoLists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": "user created"})
}

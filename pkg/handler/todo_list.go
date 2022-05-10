package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
		return
	}
	zap.L().Sugar().Debug("create TODO list ", userID, inputData)
	todoListID, err := h.services.TodoList.Create(userID, inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ID": todoListID})
}

func (h *Handler) deleteTodoList(c *gin.Context) {
	userID := c.Value("userID").(int)
	listID := c.Param("id")

	listIDInt, err := strconv.Atoi(listID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !h.isOwner(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	h.services.TodoList.Delete(userID, listIDInt)
}

func (h *Handler) getTodoList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": "user created"})
}

func (h *Handler) updateTodoList(c *gin.Context) {
}

func (h *Handler) getAllTodoLists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": "user created"})
}

func (h *Handler) isOwner(c *gin.Context) bool {
	userID := c.Value("userID").(int)
	listID := c.Param("id")

	listIDInt, err := strconv.Atoi(listID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	ownerID, err := h.services.TodoList.GetOwnerID(listIDInt)

	if ownerID != userID {
		return false
	}
	return true
}

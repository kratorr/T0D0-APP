package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"todo/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTodoElement(c *gin.Context) {
	userID := c.Value("userID").(int)
	listID := c.Param("id")

	listIDInt, err := strconv.Atoi(listID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inputData := models.TodoElement{}
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		fmt.Println(err)
		return
	}
	inputData.TodoListID = listIDInt
	fmt.Println(inputData)
	todoElementID, err := h.services.TodoElement.Create(userID, inputData)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"ID": todoElementID})
}

func (h *Handler) getTodoElement(c *gin.Context) {
	fmt.Println("getTodoElement")
}

func (h *Handler) deleteTodoElement(c *gin.Context) {
	fmt.Println("deleteTodoElement")
}

func (h *Handler) updateTodoElement(c *gin.Context) {
	fmt.Println("updateTodoElement")
}

func (h *Handler) getTodoElementsByList(c *gin.Context) {
	fmt.Println("updateTodoElement")
}

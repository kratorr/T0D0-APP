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

	todoListDto := models.CreateTodoListDTO{UserID: userID}
	err := c.ShouldBindJSON(&todoListDto)
	if err != nil {
		zap.L().Sugar().Error(err.Error())
		return
	}
	zap.L().Sugar().Debug("create TODO list ", userID, todoListDto)
	todoListID, err := h.services.TodoList.Create(userID, todoListDto)
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

	h.services.TodoList.Delete(userID, listIDInt)
	c.JSON(http.StatusOK, gin.H{"message": "list deleted"})
}

func (h *Handler) getTodoList(c *gin.Context) {
	userID := c.Value("userID").(int)
	listID := c.Param("id")

	listIDInt, err := strconv.Atoi(listID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoList, err := h.services.TodoList.GetByID(userID, listIDInt)
	if err != nil {
		fmt.Println("todoLis")
	}
	fmt.Println(todoList)

	c.JSON(http.StatusOK, todoList)
}

func (h *Handler) updateTodoList(c *gin.Context) {
	userID := c.Value("userID").(int)
	listID := c.Param("id")

	listIDInt, err := strconv.Atoi(listID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inputData := &models.TodoList{}

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		zap.L().Sugar().Info("Error create todolist struct", userID, listID, inputData)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.services.TodoList.Update(userID, listIDInt, *inputData)
	c.JSON(http.StatusOK, gin.H{"message": "todo list updated"})
}

func (h *Handler) getAllTodoLists(c *gin.Context) {
	userID := c.Value("userID").(int)

	todoLists, err := h.services.TodoList.GetAll(userID)
	if err != nil {
		fmt.Println("todoList")
	}
	fmt.Println(todoLists)
	c.JSON(http.StatusOK, todoLists)
}

package handler

import (
	"todo/pkg/service"

	_ "todo/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes(g *gin.Engine) *gin.Engine {
	auth := g.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.SignIn)
		// auth.POST("/signout", h.SignOut)
	}
	api := g.Group("/api/v1")

	api.Use(h.UserIdentification)
	api.GET("/test", h.Test)

	todoLists := api.Group("/todo_list")

	{
		todoLists.GET("/", h.getAllTodoLists)
		todoLists.POST("/", h.createTodoList)
		todoLists.GET("/:id", h.getTodoList)
		todoLists.DELETE("/:id", h.deleteTodoList)
		todoLists.PUT("/:id", h.updateTodoList)
	}

	todoElement := api.Group("/todo")
	{
		todoElement.POST("/")
		todoElement.GET("/:id")
		todoElement.DELETE("/:id")
		todoElement.PUT("/:id")
	}

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return g
}

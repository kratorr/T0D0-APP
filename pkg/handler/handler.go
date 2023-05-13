package handler

import (
	"todo/pkg/service"

	_ "todo/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services  *service.Service
	secretKey string
}

func NewHandler(services *service.Service, secretKey string) *Handler {
	return &Handler{
		services:  services,
		secretKey: secretKey,
	}
}

func (h *Handler) InitRoutes(g *gin.Engine) *gin.Engine {
	auth := g.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.SignIn)
		auth.GET("/me", h.UserIdentification, h.Me)
		// auth.POST("/signout", h.SignOut)
	}
	api := g.Group("/api/v1")

	api.Use(h.UserIdentification)

	todoLists := api.Group("/todo_list")

	{
		todoLists.GET("/", h.getAllTodoLists)
		todoLists.POST("/", h.createTodoList)
		todoLists.GET("/:id", h.getTodoList)
		todoLists.DELETE("/:id", h.deleteTodoList)
		todoLists.PUT("/:id", h.updateTodoList)

		todoLists.POST(":id/todo", h.createTodoElement) // create todo element
		todoLists.DELETE(":id/todo/:elemid", h.deleteTodoElement)
	}

	// todoElement := api.Group("/todo")
	{
		// todoElement.GET("/", h.getTodoElementsByList)
		// todoElement.POST("/", h.createTodoElement)
		// todoElement.GET("/:id", h.getTodoElement)
		// todoElement.DELETE("/:id", h.deleteTodoElement)
		// todoElement.PUT("/:id", h.updateTodoElement)
	}

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return g
}

//{"todoll_list":5 , "titl": "go to shop"}

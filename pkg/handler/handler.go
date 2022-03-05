package handler

import (
	"todo/pkg/service"

	"github.com/gin-gonic/gin"
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

	auth.POST("/signup", h.SignUp)
	auth.POST("/signin", h.SignIn)

	api := g.Group("/api/v1")
	api.Use(h.UserIdentification)
	api.GET("/test", h.Test)

	return g
}

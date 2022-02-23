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

	//	g.POST("/auth", h.auth)
	// g.GET("/i", h.i)
	// g.POST("/upload_template", h.i, h.uploadTemplate)

	// g.POST("/send", h.i, h.sendTemplate)
	// g.POST("/result", h.i, h.updateStatus)

	// receivers := g.Group("/receivers")

	// receivers.POST("/:user_id", h.i, h.createReceivers)  // C
	// receivers.GET("/:user_id", h.i, h.getReceiver)       // R
	// receivers.PUT("/:user_id", h.i, h.updateReceiver)    // U
	// receivers.DELETE("/:user_id", h.i, h.deleteReceiver) // D
	// receivers.GET("", h.i, h.getReceivers)               // L

	return g
}

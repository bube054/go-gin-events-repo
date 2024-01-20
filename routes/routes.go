package routes

import (
	"github.com/gin-gonic/gin"
	"example.com/learning/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events/:id", GetEvent)
	server.GET("/events", GetEvents)

	authGroup := server.Group("/")
	authGroup.Use(middlewares.Authenticate)
	authGroup.POST("/events", CreateEvent)
	authGroup.PUT("/events/:id", UpdateEvent)
	authGroup.DELETE("/events/:id", DeleteEvent)
	authGroup.POST("/events/:id/register", registerForEvent)
	authGroup.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", SignUp)
	server.POST("/login", Login)
}

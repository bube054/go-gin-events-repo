package main

import (
	_ "fmt"

	"example.com/learning/db"
	"example.com/learning/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}


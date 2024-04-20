package main

import (
	_ "fmt"

	"github.com/bube054/go-gin-events-scheduler/db"
	"github.com/bube054/go-gin-events-scheduler/cron"
	"github.com/bube054/go-gin-events-scheduler/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	cron.PingDB()

	server.Run(":8080")
}

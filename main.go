package main

import (
	"os"

	"github.com/bube054/go-gin-events-scheduler/cron"
	"github.com/bube054/go-gin-events-scheduler/db"
	"github.com/bube054/go-gin-events-scheduler/routes"
	"github.com/bube054/go-gin-events-scheduler/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	err := utils.LoadENV()

	if err != nil {
		panic(err.Error())
	}

	err = db.InitDB()

	if err != nil {
		panic(err.Error())
	}

	server := gin.Default()

	routes.RegisterRoutes(server)

	cron.PingDB()

	PORT := os.Getenv("TURSO_DB_URL")

	if PORT == "" {
		PORT = "8081"
	}

	server.Run(":" + PORT)
}

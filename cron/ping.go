package cron

import (
	"fmt"

	"github.com/bube054/go-gin-events-scheduler/models"
	"github.com/robfig/cron"
)

func PingDB() {
	c := cron.New()

	c.AddFunc("@daily", func() {
		event, err := models.GetEventById(1)

		if err != nil {
			fmt.Println("ping db cron fail", err)
		}else{
			fmt.Println("ping db cron good", event)
		}
	})

	c.Start()
}

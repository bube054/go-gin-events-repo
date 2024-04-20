package routes

import (
	"fmt"
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/bube054/go-gin-events-scheduler/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events try again later."})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func GetEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func CreateEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	event.UserID = userId

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event. Try again later."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func UpdateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	userId := ctx.GetInt64("userId")

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func DeleteEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event to be deleted."})
		return
	}

	if event.UserID == userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}

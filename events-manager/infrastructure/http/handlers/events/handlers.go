package events

import (
	"events-manager/domain/events/dtos"
	"events-manager/domain/events/models"
	events "events-manager/domain/events/usecases"
	"events-manager/infrastructure/events/adapters/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createEvent(
	createUseCase events.CreateEventUseCase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json models.Event
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		event, err := createUseCase.Execute(c, json)
		if err != nil {
			c.JSON(500, gin.H{
				"error": gin.H{
					"message": "can not create an event",
				},
			})
			return
		}
		c.JSON(201, event)
	}
}

func getAllEvents(
	getAllEventsUseCase events.GetAllEventsUseCase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		event, err := getAllEventsUseCase.Execute(c)
		if err != nil {
			c.JSON(500, gin.H{
				"error": gin.H{
					"message": "can not retreive events",
				},
			})
			return
		}
		c.JSON(200, event)
	}
}

func getEventById(
	getEventByIdUseCase events.GetEventByIdUseCase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		event, err := getEventByIdUseCase.Execute(c, c.Param("id"))
		if err != nil {
			c.JSON(500, gin.H{
				"error": gin.H{
					"message": "can not retreive event",
				},
			})
			return
		}
		c.JSON(200, event)
	}
}

func deleteEventById(
	deleteEventByIdUseCase events.DeleteEventByIdUseCase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		event, err := deleteEventByIdUseCase.Execute(c, c.Param("id"))
		if err != nil {
			c.JSON(500, gin.H{
				"error": gin.H{
					"message": "can not delete event",
				},
			})
			return
		}
		c.JSON(200, event)
	}
}

func updateEvent(
	updateEventUseCase events.UpdateEventUseCase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json models.Event
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		event, err := updateEventUseCase.Execute(c, json)
		if err != nil {
			c.JSON(500, gin.H{
				"error": gin.H{
					"message": "can not update event",
				},
			})
			return
		}
		c.JSON(200, event)
	}
}

func addAttendeeEvent(
	addAttendeeEventUseCase events.AddAttendeeEventUseCase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dtos.AddAttendeDTO
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		event, err := addAttendeeEventUseCase.Execute(c, json)
		if err != nil {
			if v, ok := err.(*errors.DuplicateAttendee); ok {
				c.JSON(409, gin.H{
					"error": gin.H{
						"message": v.Error(),
					},
				})
				return
			}
			c.JSON(500, gin.H{
				"error": gin.H{
					"message": "can not add attendee to event",
				},
			})
			return
		}
		c.JSON(200, event)
	}
}

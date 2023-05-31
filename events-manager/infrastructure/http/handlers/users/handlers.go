package users

import (
	"events-manager/domain/users/models"
	users "events-manager/domain/users/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(
	createUseCase users.CreateUserUseCase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json models.User
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := createUseCase.Execute(c, json)
		if err != nil {
			c.JSON(500, gin.H{
				"error": gin.H{
					"message": "can not create a new user",
				},
			})
			return
		}
		c.JSON(201, user)
	}
}

func getUserByEmail(
	getUserByEmailUseCase users.GetUserByEmailUseCase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := getUserByEmailUseCase.Execute(c, c.Param("email"))
		if err != nil {
			c.JSON(500, gin.H{
				"error": gin.H{
					"message": "can not retreive user",
				},
			})
			return
		}
		c.JSON(200, user)
	}
}

// func getAllEvents(
// 	getAllEventsUseCase events.GetAllEventsUseCase,
// ) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		event, err := getAllEventsUseCase.Execute(c)
// 		if err != nil {
// 			c.JSON(500, gin.H{
// 				"error": gin.H{
// 					"message": "can not retreive event",
// 				},
// 			})
// 			return
// 		}
// 		c.JSON(200, event)
// 	}
// }

// func deleteEventById(
// 	deleteEventByIdUseCase events.DeleteEventByIdUseCase,
// ) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		event, err := deleteEventByIdUseCase.Execute(c, c.Param("id"))
// 		if err != nil {
// 			c.JSON(500, gin.H{
// 				"error": gin.H{
// 					"message": "can not delete event",
// 				},
// 			})
// 			return
// 		}
// 		c.JSON(200, event)
// 	}
// }

// func updateEvent(
// 	updateEventUseCase events.UpdateEventUseCase,
// ) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var json models.User
// 		if err := c.ShouldBindJSON(&json); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		event, err := updateEventUseCase.Execute(c, json)
// 		if err != nil {
// 			c.JSON(500, gin.H{
// 				"error": gin.H{
// 					"message": "can not create an event",
// 				},
// 			})
// 			return
// 		}
// 		c.JSON(201, event)
// 	}
// }

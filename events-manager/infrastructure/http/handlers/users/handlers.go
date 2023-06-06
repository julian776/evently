package users

import (
	"events-manager/domain/users/dtos"
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

func login(
	loginUserUseCase users.LoginUserUseCase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dtos.LoginDTO
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := loginUserUseCase.Execute(c, json)
		if err != nil {
			c.JSON(400, gin.H{
				"error": gin.H{
					"message": "incorrect password or user not registered",
				},
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "succesfully logged in",
			"user":    user,
		})
	}
}

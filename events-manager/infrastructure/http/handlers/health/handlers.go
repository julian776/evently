package health

import (
	domain "events-manager/domain/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, OkStatus{
			AppName:    domain.APP_NAME,
			AppVersion: domain.APP_VERSION,
		})
	}
}

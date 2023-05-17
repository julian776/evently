package health

import (
	app "events-manager/infrastructure/app/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, OkStatus{
			AppName:    app.APP_NAME,
			AppVersion: app.APP_VERSION,
		})
	}
}

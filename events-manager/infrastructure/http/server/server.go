package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Creates a new instance of the `http.Server` struct and
// initializing it with an address of `:8080` and a `gin.Engine`
// instance as the handler.
// It also sets up middleware for recovering from panics.
func NewServer() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(enableCors)
	return engine
}

func enableCors(ctx *gin.Context) {
	ctx.Writer.Header().Add("Access-Control-Allow-Methods", "*")
	ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Add("Access-Control-Allow-Headers", "*")
	ctx.Writer.Header().Add("Access-Control-Max-Age", "3600")

	if ctx.Request.Method == http.MethodOptions {
		ctx.String(http.StatusOK, "")
	}
}

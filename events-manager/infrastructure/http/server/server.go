package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Creates a new instance of the `http.Server` struct and
// initializing it with an address of `:8080` and a `gin.Engine`
// instance as the handler.
// It also sets up middleware for recovering from panics.
func NewServer() *http.Server {
	engine := gin.New()
	engine.Use(gin.Recovery())

	server := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	return server
}

package routes

import (
	"kauanpecanha/devops-challenge/controllers"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func NewHTTPHandler(name string) *gin.Engine {

	router := gin.Default()
	router.Use(otelgin.Middleware(name))
	controllers.RegisterRollDiceRoutes(router)

	// Start the server
	err := router.Run("localhost:8080")
	if err != nil {
		return nil
	}

	return router
}

package routes

import (
	"kauanpecanha/devops-challenge/controllers"

	"github.com/gin-gonic/gin"
)

func NewHTTPHandler() *gin.Engine {
	router := gin.Default()
	controllers.RegisterRollDiceRoutes(router)
	return router
}

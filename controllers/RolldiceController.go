package controllers

import (
	"kauanpecanha/devops-challenge/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRollDiceRoutes(r *gin.Engine) {
	r.GET("/", handlers.WelcomeHome)
	r.POST("/roll", handlers.Play)
	r.POST("/roll/:player", handlers.Play)
	r.GET("/roll", handlers.GetAllPlays)
	r.GET("/roll/:id", handlers.GetPlayByID)
	r.PUT("/roll/:id", handlers.UpdatePlay)
	r.DELETE("/roll/:id", handlers.DeletePlay)
}

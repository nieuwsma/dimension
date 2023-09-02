package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Public routes (No auth required)

	router.GET("/rules", GetGameRules)

	router.POST("/training", StartTrainingSession)
	router.GET("/training", RetrieveTrainingIDs)
	router.GET("/training/:trainID", RetrieveTrainingStatus)

	//todo maybe I should split up the AuthMiddleware to make it easier to guess the context on what to parse?
	// e.g. that for RemovePlayerFromGame I need the Game session, not the player session?

	// Routes that require authentication

	// Training Routes
	router.PATCH("/training/:trainID", PlayTrainingSession)
	router.POST("/training/:trainID/regenerate", RegenerateTrainingSession)

	return router
}

func Index(c *gin.Context) {
	fmt.Fprint(c.Writer, "Welcome to Dimension")
}

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/rules", GetGameRules)
	// Training Routes

	router.POST("/training", StartTrainingSession)
	router.GET("/training", RetrieveTrainingIDs)
	router.GET("/training/:trainID", RetrieveTrainingSessions)
	router.PUT("/training/:trainID/turn/:playerName", PlayTrainingSessionTurn)
	router.POST("/training/:trainID/regenerate", RegenerateTrainingSession)

	return router
}

func Index(c *gin.Context) {
	fmt.Fprint(c.Writer, "Welcome to Dimension")
}

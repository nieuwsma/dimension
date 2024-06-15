package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Manually configure CORS
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust to your client's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))
	//
	//// Add CORS middleware
	//router.Use(cors.Default())

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

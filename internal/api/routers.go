package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //if this is debug mode it prints a bunch of stupid stuff in the CLI!

	router := gin.Default()

	// Game Management Routes
	router.POST("/games", CreateGame)
	router.DELETE("/games/:gameId", DeleteGame)
	router.GET("/games/:gameId", GetGameDetails)
	router.DELETE("/games/:gameId/players/:playerId", RemovePlayerFromGame)
	router.POST("/games/:gameId/players", AddPlayerToGame)

	// Rounds Routes
	router.POST("/games/:gameId/rounds", StartNewRound)
	router.GET("/games/:gameId/rounds", GetRounds)
	router.GET("/games/:gameId/rounds/:roundId", GetSpecificRoundStatus)
	router.PATCH("/games/:gameId/rounds/:roundId", ForceRoundCompletion)

	// Players Routes
	router.POST("/games/:gameId/rounds/:roundId/turns/:playerId", PlayerTakeTurn)

	// Rules Route
	router.GET("/rules", GetGameRules)

	// Training Routes
	router.POST("/training", StartTrainingSession)
	router.PATCH("/training/:trainID", PlayTrainingSession)
	router.GET("/training/:trainID", RetrieveTrainingStatus)

	return router
}

func Index(c *gin.Context) {
	fmt.Fprint(c.Writer, "Hello World!")
}

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Public routes (No auth required)
	router.POST("/games", CreateGame)

	router.GET("/games/:gameId", GetGameDetails)
	router.POST("/games/:gameId/players", AddPlayerToGame) //anyone could add a player to a game? maybe have a join session? probably in the body
	//maybe the gameID knowledge itself is the join session, like how that one online website works...
	router.GET("/games/:gameId/rounds", GetRounds)
	router.GET("/games/:gameId/rounds/:roundId", GetSpecificRoundStatus)
	router.GET("/rules", GetGameRules)

	router.POST("/training", StartTrainingSession)
	router.GET("/training/:trainID", RetrieveTrainingStatus)

	//todo maybe I should split up the AuthMiddleware to make it easier to guess the context on what to parse?
	// e.g. that for RemovePlayerFromGame I need the Game session, not the player session?

	// Routes that require authentication
	// Game Management Routes
	router.DELETE("/games/:gameId", AuthMiddleware("gameId"), DeleteGame)
	router.DELETE("/games/:gameId/players/:playerId", AuthMiddleware("gameId"), RemovePlayerFromGame) //todo should you be able to delete yourself RemoveSelfFromGame or ResignFromGame??

	// Rounds Routes
	router.POST("/games/:gameId/rounds", AuthMiddleware("gameId"), ForceStartNewRound)
	router.PATCH("/games/:gameId/rounds/:roundId", AuthMiddleware("gameId"), ForceRoundCompletion)

	// Players Routes
	router.POST("/games/:gameId/rounds/:roundId/turns/:playerId", AuthMiddleware("playerId"), PlayerTakeTurn)

	// Training Routes
	router.PATCH("/training/:trainID", AuthMiddleware("trainID"), PlayTrainingSession)
	router.POST("/training/:trainID/regnerate", AuthMiddleware("trainID"), RegenerateTrainingSession)

	return router
}

func Index(c *gin.Context) {
	fmt.Fprint(c.Writer, "Hello World!")
}

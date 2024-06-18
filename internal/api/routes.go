package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nieuwsma/dimension/internal/logger"
	"github.com/nieuwsma/dimension/internal/middleware"
	"github.com/nieuwsma/dimension/pkg/logic"
	"github.com/nieuwsma/dimension/pkg/presentation"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Rules Route

func GetGameRules(c *gin.Context) {
	// Logic to get game rules

	rules, colors, geometries := middleware.GetGameRules()

	var gameRules presentation.GetRulesResponse

	for _, v := range rules.Set {
		gameRules.Tasks = append(gameRules.Tasks, presentation.Task{
			Name:        v.Name,
			Quantity:    v.Quantity,
			Description: v.Description,
		})
	}
	gameRules.Colors = colors
	gameRules.Geometries = geometries

	pb := presentation.BuildSuccessPassback(http.StatusOK, gameRules)
	WriteJsonWithHeaders(c, pb)
}

// Training Routes

func RetrieveTrainingIDs(c *gin.Context) {
	var pb presentation.APIPayload

	trainingSessions, err := middleware.RetrieveTrainingSessions()
	if err != nil {
		pb = presentation.BuildErrorPassback(http.StatusInternalServerError, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("internal server error")
		WriteJsonWithHeaders(c, pb)
		return
	}

	var response presentation.GetTrainingSessionIDResponse
	for k, _ := range trainingSessions {
		response.TrainingSessionID = append(response.TrainingSessionID, k)
	}

	pb = presentation.BuildSuccessPassback(http.StatusOK, response)
	WriteJsonWithHeaders(c, pb)
	return
}

func StartTrainingSession(c *gin.Context) {
	// Logic to start a training session
	var pb presentation.APIPayload

	var parameters presentation.PostTrainingSessionRequest

	if err := c.ShouldBindJSON(&parameters); err != nil {
		pb = presentation.BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteJsonWithHeaders(c, pb)
		return
	}

	var tasksRequest logic.Tasks
	for _, v := range parameters.Types {
		tasksRequest = append(tasksRequest, logic.Task(v))
	}

	trainID, tasks, err := middleware.StartTrainingSession(tasksRequest)

	if err != nil {
		pb = presentation.BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteJsonWithHeaders(c, pb)
		return
	}

	response := presentation.PostTrainingSessionResponse{
		TrainID: trainID,
		Tasks:   tasks,
	}

	pb = presentation.BuildSuccessPassback(http.StatusCreated, response)
	WriteJsonWithHeaders(c, pb)
	return
}

func PlayTrainingSessionTurn(c *gin.Context) {

	var pb presentation.APIPayload

	// Logic to play the training session
	trainID := c.Param("trainID")
	if trainID == "" {
		err := errors.New("invalid request, missing trainID parameter")
		pb := presentation.BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteJsonWithHeaders(c, pb)
		return
	}

	playerName := c.Param("playerName")
	if trainID == "" {
		err := errors.New("invalid request, missing playerName parameter")
		pb := presentation.BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteJsonWithHeaders(c, pb)
		return
	}

	var parameters presentation.PostTrainingSessionTurnRequest

	if err := c.ShouldBindJSON(&parameters); err != nil {
		pb = presentation.BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteJsonWithHeaders(c, pb)
		return
	}
	dimension, err := parameters.ToLogicDimension()
	if err != nil {
		err := errors.New("invalid dimension:" + err.Error())
		pb := presentation.BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteJsonWithHeaders(c, pb)
		return
	}
	trainingSession, err := middleware.PlayTrainingSession(trainID, playerName, *dimension)

	if err != nil {
		err := errors.New("invalid request trainID not found")
		pb := presentation.BuildErrorPassback(http.StatusNotFound, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteJsonWithHeaders(c, pb)
		return
	}

	response := &presentation.PutTrainingSessionTurnResponse{

		Tasks: trainingSession.Tasks,
		TrainingSessionTurn: presentation.TrainingSessionTurn{
			PlayerName:     playerName,
			Score:          trainingSession.Turns[logic.PlayerName(playerName)].Score,
			BonusPoints:    trainingSession.Turns[logic.PlayerName(playerName)].Bonus,
			Dimension:      presentation.NewDimensionResponse(trainingSession.Turns[logic.PlayerName(playerName)].Dimension),
			TaskViolations: trainingSession.Turns[logic.PlayerName(playerName)].TaskViolations,
		},
		ExpirationTime: presentation.CustomTime{trainingSession.ExpirationTime},
	}

	pb = presentation.BuildSuccessPassback(http.StatusOK, response)
	WriteJsonWithHeaders(c, pb)
	return
}

func RetrieveTrainingSessions(c *gin.Context) {
	//{
	//	"tasks": [
	//"string"
	//],
	//"turns": [
	//{
	//"playerName": "string",
	//"score": 0,
	//"bonusPoints": true,
	//"submittedDimension": {
	//"a": "G",
	//"b": "G",
	//"c": "G",
	//"d": "G",
	//"e": "G",
	//"f": "G",
	//"g": "G",
	//"h": "G",
	//"i": "G",
	//"j": "G",
	//"k": "G",
	//"l": "G",
	//"m": "G",
	//"n": "G"
	//},
	//"taskViolations": [
	//"string"
	//]
	//}
	//],
	//"expirationTime": "2023-09-02T00:04:55.964Z"
	//}
	trainID := c.Param("trainID")
	if trainID == "" {
		err := errors.New("invalid request, missing trainID parameter")
		pb := presentation.BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteJsonWithHeaders(c, pb)
		return
	}

	trainingSession, err := middleware.RetrieveTrainingStatus(trainID)

	if err != nil {
		err := errors.New("invalid request trainID not found")
		pb := presentation.BuildErrorPassback(http.StatusNotFound, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteJsonWithHeaders(c, pb)
		return
	}

	response := &presentation.GetTrainingSessionsResponse{
		Tasks:          trainingSession.Tasks,
		ExpirationTime: presentation.CustomTime{trainingSession.ExpirationTime},
	}
	for _, v := range trainingSession.Turns {
		response.TrainingSessionTurn = append(response.TrainingSessionTurn, presentation.TrainingSessionTurn{
			PlayerName:     string(v.PlayerName),
			Score:          v.Score,
			BonusPoints:    v.Bonus,
			Dimension:      presentation.NewDimensionResponse(v.Dimension),
			TaskViolations: v.TaskViolations,
		})
	}

	pb := presentation.BuildSuccessPassback(http.StatusOK, response)
	WriteJsonWithHeaders(c, pb)
	return
}

func RegenerateTrainingSession(c *gin.Context) {
	// Logic to retrieve training status
	var pb presentation.APIPayload

	// Logic to play the training session
	trainID := c.Param("trainID")
	if trainID == "" {
		err := errors.New("invalid request, missing trainID parameter")
		pb := presentation.BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteJsonWithHeaders(c, pb)
		return
	}

	tasks, err := middleware.RegenerateTrainingSession(trainID)

	if err != nil {
		pb := presentation.BuildErrorPassback(http.StatusNotFound, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteJsonWithHeaders(c, pb)
		return
	}

	response := &presentation.PostRegenerateTrainingSessionResponse{
		Tasks: tasks,
	}

	pb = presentation.BuildSuccessPassback(http.StatusOK, response)
	WriteJsonWithHeaders(c, pb)
	return
}

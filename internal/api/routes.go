package api

import (
	"dimension/internal/logger"
	"dimension/internal/middleware"
	"dimension/pkg/logic"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Rules Route

func GetGameRules(c *gin.Context) {
	// Logic to get game rules

	rules, colors, geometries := middleware.GetGameRules()

	var gameRules RulesResponse

	for _, v := range rules.Set {
		gameRules.Tasks = append(gameRules.Tasks, Task{
			Name:        v.Name,
			Quantity:    v.Quantity,
			Description: v.Description,
		})
	}

	for _, v := range colors {
		gameRules.Colors = append(gameRules.Colors, Color{
			Name: v.Name,
			Code: v.Code,
		})
	}

	for _, v := range geometries {
		gameRules.Geometries = append(gameRules.Geometries, GeometryItem{
			PolarAngle:       v.PolarAngle,
			InclinationAngle: v.InclinationAngle,
			RadialDistance:   v.RadialDistance,
			ID:               v.ID,
			Neighbors:        v.Neighbors,
		})
	}

	pb := BuildSuccessPassback(http.StatusOK, gameRules)
	WriteHeaders(c, pb)
}

// Training Routes

func RetrieveTrainingIDs(c *gin.Context) {
	var pb APIPayload

	trainingSessions, err := middleware.RetrieveTrainingSessions()
	if err != nil {
		pb = BuildErrorPassback(http.StatusInternalServerError, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("internal server error")
		WriteHeaders(c, pb)
		return
	}

	var response GetTrainingSessionID
	for k, _ := range trainingSessions {
		response.TrainingSessionID = append(response.TrainingSessionID, k)
	}

	pb = BuildSuccessPassback(http.StatusOK, response)
	WriteHeaders(c, pb)
	return
}

func StartTrainingSession(c *gin.Context) {
	// Logic to start a training session
	var pb APIPayload

	var parameters PostTrainingSessionRequest

	if err := c.ShouldBindJSON(&parameters); err != nil {
		pb = BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}

	var tasksRequest logic.Tasks
	for _, v := range parameters.types {
		tasksRequest = append(tasksRequest, logic.Task(v))
	}

	trainID, tasks, err := middleware.StartTrainingSession(tasksRequest)

	if err != nil {
		pb = BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}

	response := PostTrainingSessionResponse{
		TrainID: trainID,
		Tasks:   tasks,
	}

	pb = BuildSuccessPassback(http.StatusCreated, response)
	WriteHeaders(c, pb)
	return
}

func PlayTrainingSessionTurn(c *gin.Context) {

	var pb APIPayload

	// Logic to play the training session
	trainID := c.Param("trainID")
	if trainID == "" {
		err := errors.New("invalid request, missing trainID parameter")
		pb := BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}

	playerName := c.Param("playerName")
	if trainID == "" {
		err := errors.New("invalid request, missing playerName parameter")
		pb := BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}

	var parameters Dimension

	if err := c.ShouldBindJSON(&parameters); err != nil {
		pb = BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}
	dimension, _ := parameters.ToLogicDimension()

	trainingSession, err := middleware.PlayTrainingSession(trainID, playerName, *dimension)

	if err != nil {
		err := errors.New("invalid request trainID not found")
		pb := BuildErrorPassback(http.StatusNotFound, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}

	response := &PutTrainingSessionTurnResponse{

		Tasks: trainingSession.Tasks,
		TrainingSessionTurn: TrainingSessionTurn{
			PlayerName:     playerName,
			Score:          trainingSession.Turns[logic.PlayerName(playerName)].Score,
			BonusPoints:    trainingSession.Turns[logic.PlayerName(playerName)].Bonus,
			Dimension:      NewDimensionResponse(trainingSession.Turns[logic.PlayerName(playerName)].Dimension).Dimension,
			TaskViolations: Unwrap(trainingSession.Turns[logic.PlayerName(playerName)].TaskViolations),
		},
		ExpirationTime: CustomTime{trainingSession.ExpirationTime},
	}

	pb = BuildSuccessPassback(http.StatusOK, response)
	WriteHeaders(c, pb)
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
		pb := BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}

	trainingSession, err := middleware.RetrieveTrainingStatus(trainID)

	if err != nil {
		err := errors.New("invalid request trainID not found")
		pb := BuildErrorPassback(http.StatusNotFound, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}

	response := &GetTrainingSessionsResponse{
		Tasks:          trainingSession.Tasks,
		ExpirationTime: CustomTime{trainingSession.ExpirationTime},
	}
	for _, v := range trainingSession.Turns {
		response.TrainingSessionTurn = append(response.TrainingSessionTurn, TrainingSessionTurn{
			PlayerName:     string(v.PlayerName),
			Score:          v.Score,
			BonusPoints:    v.Bonus,
			Dimension:      NewDimensionResponse(v.Dimension).Dimension,
			TaskViolations: Unwrap(v.TaskViolations),
		})
	}

	pb := BuildSuccessPassback(http.StatusOK, response)
	WriteHeaders(c, pb)
	return
}

func RegenerateTrainingSession(c *gin.Context) {
	// Logic to retrieve training status
	var pb APIPayload

	// Logic to play the training session
	trainID := c.Param("trainID")
	if trainID == "" {
		err := errors.New("invalid request, missing trainID parameter")
		pb := BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}

	tasks, err := middleware.RegenerateTrainingSession(trainID)

	if err != nil {
		pb := BuildErrorPassback(http.StatusNotFound, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}

	response := &PostRegenerateTrainingSessionResponse{
		Tasks: tasks,
	}

	pb = BuildSuccessPassback(http.StatusOK, response)
	WriteHeaders(c, pb)
	return
}

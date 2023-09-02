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
	c.JSON(http.StatusOK, gin.H{
		"message": "Game rules",
	})

	//r := RulesResponse{
	//	Tasks:      nil,
	//	Geometries: nil,
	//	Colors:     logic.GetColors(),
	//}
}

// Training Routes

func RetrieveTrainingIDs(c *gin.Context) {

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

	ts := &GetTrainingSessionResponse{
		Score:              trainingSession.Turn.Score,
		BonusPoints:        trainingSession.Turn.Bonus,
		SubmittedDimension: NewDimensionResponse(trainingSession.Turn.Dimension),
		Tasks:              trainingSession.Tasks,
		TaskViolations:     Unwrap(trainingSession.Turn.TaskViolations),
		ExpirationTime:     CustomTime{trainingSession.ExpirationTime},
	}

	pb := BuildSuccessPassback(http.StatusOK, ts)
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

func PlayTrainingSession(c *gin.Context) {

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

	var parameters Dimension

	if err := c.ShouldBindJSON(&parameters); err != nil {
		pb = BuildErrorPassback(http.StatusBadRequest, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}
	dimension, _ := parameters.ToLogicDimension()

	trainingSession, err := middleware.PlayTrainingSession(trainID, *dimension)

	if err != nil {
		err := errors.New("invalid request trainID not found")
		pb := BuildErrorPassback(http.StatusNotFound, err)
		logger.Log.WithFields(logrus.Fields{"ERROR": err, "HttpStatusCode": pb.StatusCode}).Error("bad request")
		WriteHeaders(c, pb)
		return
	}

	response := &GetTrainingSessionResponse{
		Score:              trainingSession.Turn.Score,
		BonusPoints:        trainingSession.Turn.Bonus,
		SubmittedDimension: NewDimensionResponse(trainingSession.Turn.Dimension),
		Tasks:              trainingSession.Tasks,
		TaskViolations:     Unwrap(trainingSession.Turn.TaskViolations),
		ExpirationTime:     CustomTime{trainingSession.ExpirationTime},
	}

	pb = BuildSuccessPassback(http.StatusOK, response)
	WriteHeaders(c, pb)
	return
}

func RetrieveTrainingStatus(c *gin.Context) {

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

	ts := &GetTrainingSessionResponse{
		Score:              trainingSession.Turn.Score,
		BonusPoints:        trainingSession.Turn.Bonus,
		SubmittedDimension: NewDimensionResponse(trainingSession.Turn.Dimension),
		Tasks:              trainingSession.Tasks,
		TaskViolations:     Unwrap(trainingSession.Turn.TaskViolations),
		ExpirationTime:     CustomTime{trainingSession.ExpirationTime},
	}

	pb := BuildSuccessPassback(http.StatusOK, ts)
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

package presentation

import (
	"github.com/nieuwsma/dimension/pkg/geometry"
	"github.com/nieuwsma/dimension/pkg/logic"
)

// PostTrainingSessionTurnRequest represents dimension properties
type PostTrainingSessionTurnRequest struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
	E string `json:"e"`
	F string `json:"f"`
	G string `json:"g"`
	H string `json:"h"`
	I string `json:"i"`
	J string `json:"j"`
	K string `json:"k"`
	L string `json:"l"`
	M string `json:"m"`
	N string `json:"n"`
}

// GetRulesResponse represents a response for game rules
type GetRulesResponse struct {
	Tasks      []Task              `json:"Tasks"`
	Geometries geometry.Geometries `json:"Geometries"`
	Colors     logic.Colors        `json:"Colors"`
}

type GetTrainingSessionIDResponse struct {
	TrainingSessionID []string `json:"trainingSessions"`
}

type PostTrainingSessionRequest struct {
	Types []string `json:"taskTypes"`
}

type PostTrainingSessionResponse struct {
	TrainID string      `json:"trainID"`
	Tasks   logic.Tasks `json:"tasks"`
}

type PostRegenerateTrainingSessionResponse struct {
	Tasks logic.Tasks `json:"tasks"`
}

type GetTrainingSessionsResponse struct {
	TrainingSessionTurn []TrainingSessionTurn `json:"turn"`
	Tasks               logic.Tasks           `json:"tasks"`
	ExpirationTime      CustomTime            `json:"expirationTime"`
}

type PutTrainingSessionTurnResponse struct {
	TrainingSessionTurn TrainingSessionTurn `json:"turn"`
	Tasks               logic.Tasks         `json:"tasks"`
	ExpirationTime      CustomTime          `json:"expirationTime"`
}

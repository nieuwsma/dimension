package api

import (
	"errors"
	"github.com/nieuwsma/dimension/pkg/logic"
	"time"
)

// Dimension represents dimension properties
type Dimension struct {
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

type DimensionResponse struct {
	Dimension map[string]string
}

// Task represents a task
type Task struct {
	Name        string `json:"Name"`
	Quantity    int    `json:"Quantity"`
	Description string `json:"Description"`
}

// RulesResponse represents a response for game rules
type RulesResponse struct {
	Tasks      []Task         `json:"Tasks"`
	Geometries []GeometryItem `json:"Geometries"`
	Colors     []Color        `json:"Colors"`
}

// GeometryItem represents a geometry item
type GeometryItem struct {
	PolarAngle       int      `json:"polarAngle"`
	InclinationAngle int      `json:"inclinationAngle"`
	RadialDistance   int      `json:"radialDistance"`
	ID               string   `json:"id"`
	Neighbors        []string `json:"neighbors"`
}

// Color represents a color
type Color struct {
	Name string `json:"Name"`
	Code string `json:"Code"`
}

type PostTrainingSessionRequest struct {
	types []string `json:"taskTypes"`
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

type TrainingSessionTurn struct {
	PlayerName     string            `json:"playerName""`
	Score          int               `json:"score"`
	BonusPoints    bool              `json:"bonusPoints"`
	Dimension      map[string]string `json:"dimension"`
	TaskViolations []string          `json:"taskViolations"`
}

type GetTrainingSessionID struct {
	TrainingSessionID []string `json:"trainingSessions"`
}

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05.999Z"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	s = s[1 : len(s)-1] // Strip quotes
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Time.Format(ctLayout) + `"`), nil
}

func Unwrap(err error) (errs []string) {
	for err != nil {
		errs = append(errs, err.Error())
		err = errors.Unwrap(err)
	}
	return

}

func NewDimensionResponse(dimension logic.Dimension) (ndr DimensionResponse) {
	ndr.Dimension = make(map[string]string)

	for k, v := range dimension.Dimension {
		ndr.Dimension[k] = v.Color.String()
	}
	return
}

func (reqDim *Dimension) ToLogicDimension() (dimension *logic.Dimension, err error) {
	var pairs []logic.SpherePair

	// Helper function to generate SpherePair from SphereID and Color string
	addSpherePair := func(id logic.SphereID, colorStr string) {
		// If colorStr is empty, skip creating SpherePair
		if colorStr == "" {
			return
		}

		color := logic.NewColorShort(colorStr)
		pairs = append(pairs, *logic.NewSpherePair(id, color))
	}

	// Convert each field from REQUESTDimension
	addSpherePair("a", reqDim.A)
	addSpherePair("b", reqDim.B)
	addSpherePair("c", reqDim.C)
	addSpherePair("d", reqDim.D)
	addSpherePair("e", reqDim.E)
	addSpherePair("f", reqDim.F)
	addSpherePair("g", reqDim.G)
	addSpherePair("h", reqDim.H)
	addSpherePair("i", reqDim.I)
	addSpherePair("j", reqDim.J)
	addSpherePair("k", reqDim.K)
	addSpherePair("l", reqDim.L)
	addSpherePair("m", reqDim.M)
	addSpherePair("n", reqDim.N)

	// Create the dimension from SpherePairs
	dimension, err = logic.NewDimension(pairs...)
	if err != nil {
		return nil, errors.New("failed to create dimension: " + err.Error())
	}

	return dimension, nil
}

package middleware

import "dimension/pkg/logic"

//TODO consider adding a game history and tracking players over time! #goldplatting
//type GameHistory []Game

var GamesMap = make(map[string]logic.Game)
var TrainingMap = make(map[string]logic.TrainingSession)

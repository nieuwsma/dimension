package rules

import (
	"dimension/pkg/domain"
	_ "embed"
	"encoding/json"
)

// TODO move to API!!!!
type RulesArrayDataFormat []RulesDataFormat

type RulesDataFormat struct {
	Name        string
	Quantity    int
	Description string
}

// this is the master deck! //todo this needs to change because when I run main the relative path import is missing!
//
//go:embed rules.json
var RulesData []byte

func GetAllRules() (r RulesArrayDataFormat) {
	_ = json.Unmarshal(RulesData, &r)
	return
}

func (f *RulesArrayDataFormat) ToTasks() (tasks domain.Tasks) {
	for _, rule := range *f {
		tasks = append(tasks, domain.Task(rule.Name))
	}
	return
}

func init() {
	r := GetAllRules()
	masterTasks := r.ToTasks()
	DefaultTasks = masterTasks
}

// todo do I want different tasks? how does this differ from what i need to display in API?
var DefaultTasks domain.Tasks

//TODO END OF MOVE TO API

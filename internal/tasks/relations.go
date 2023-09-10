package tasks

import (
	"fmt"
)

//
//// Sample master thesaurus
//var thesaurus = map[string][]string{
//	"hot":      {"warm", "swelter"},
//	"warm":     {"hot", "swelter"},
//	"swelter":  {"hot", "warm"},
//	"cold":     {"freezing"},
//	"freezing": {"cold"},
//	"stinky":   {"putrid"},
//	"putrid":   {"stinky"},
//}

// Function to check if two words are related
func areRelated(task1, task2 string, relations map[string][]string) bool {
	if tasks, exists := relations[task1]; exists {
		for _, task := range tasks {
			if task == task2 {
				return true
			}
		}
	}
	return false
}

func (t *TasksCollection) MapRelations() [][]string {

	relations := t.KnownRelations()
	var tasks []string
	for _, task := range t.tasks {
		tasks = append(tasks, string(task))
	}
	//tasks := []string{"hot", "warm", "swelter", "cold", "freezing", "stinky", "putrid", "tree"}
	visited := make(map[string]bool)
	groups := [][]string{}

	for _, task := range tasks {
		if visited[task] {
			continue
		}
		relatedTasks := []string{task}
		for _, innerTask := range tasks {
			if task == innerTask || visited[innerTask] {
				continue
			}
			if areRelated(task, innerTask, relations) {
				relatedTasks = append(relatedTasks, innerTask)
				visited[innerTask] = true
			}
		}
		groups = append(groups, relatedTasks)
		visited[task] = true
	}

	for _, group := range groups {
		fmt.Println(group)
	}
	return groups
}

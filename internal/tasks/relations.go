package tasks

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

func (t *TasksCollection) mapRelations() [][]string {

	relations := t.knownRelations()
	var tasks []string
	for _, task := range t.Tasks {
		tasks = append(tasks, string(task))
	}
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

	return groups
}

func (t *TasksCollection) knownRelations() (relations map[string][]string) {
	relations = make(map[string][]string)
	colorTaskMap := t.fillColorTasksDependency()

	for _, task := range colorTaskMap {
		for _, subtask := range task {
			relation := relations[subtask]
			relation = append(relation, task...)
			relations[subtask] = relation
		}
	}

	for task, tasks := range relations {
		relations[task] = removeTask(deduplicateTasks(tasks), task)

	}

	return
}

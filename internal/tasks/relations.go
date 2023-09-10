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

func (t *TasksCollection) mapRelationsRecursively() [][]string {

	relations := t.knownRelations()
	var tasks []string
	for _, task := range t.Tasks {
		tasks = append(tasks, string(task))
	}
	visited := make(map[string]bool)
	groups := [][]string{}

	// Helper function to get related tasks recursively
	var getRelatedTasks func(string, map[string]bool) []string
	getRelatedTasks = func(task string, localVisited map[string]bool) []string {
		related := []string{}

		for _, innerTask := range tasks {
			if task == innerTask || localVisited[innerTask] {
				continue
			}
			if areRelated(task, innerTask, relations) {
				related = append(related, innerTask)
				localVisited[innerTask] = true
				// recursively get related tasks of this innerTask
				for _, deeperTask := range getRelatedTasks(innerTask, localVisited) {
					if !contains(related, deeperTask) { // Check if the task isn't already in the related list
						related = append(related, deeperTask)
					}
				}
			}
		}
		return related
	}

	for _, task := range tasks {
		if visited[task] {
			continue
		}
		relatedTasks := []string{task}
		localVisited := make(map[string]bool) // to track visited tasks in this loop iteration
		localVisited[task] = true
		for _, deeperRelatedTask := range getRelatedTasks(task, localVisited) {
			if !contains(relatedTasks, deeperRelatedTask) {
				relatedTasks = append(relatedTasks, deeperRelatedTask)
			}
		}
		groups = append(groups, relatedTasks)
		for _, rt := range relatedTasks {
			visited[rt] = true
		}
	}

	return groups
}

// Helper function to check if a string slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

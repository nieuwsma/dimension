package tasks

// Function to check if two words are related
func areRelated(task1, task2 TaskTuple, relations map[TaskTuple][]TaskTuple) bool {
	if tasks, exists := relations[task1]; exists {
		for _, task := range tasks {
			if task == task2 {
				return true
			}
		}
	}
	return false
}

func (t *TasksCollection) knownRelations() (relations map[TaskTuple][]TaskTuple) {
	relations = make(map[TaskTuple][]TaskTuple)
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

func (t *TasksCollection) mapRelationsRecursively() [][]TaskTuple {

	relations := t.knownRelations()
	var tasks []TaskTuple
	for _, task := range t.Tasks {
		tasks = append(tasks, NewTaskTuple(string(task)))
	}
	visited := make(map[TaskTuple]bool)
	groups := [][]TaskTuple{}

	// Helper function to get related tasks recursively
	var getRelatedTasks func(TaskTuple, map[TaskTuple]bool) []TaskTuple
	getRelatedTasks = func(task TaskTuple, localVisited map[TaskTuple]bool) []TaskTuple {
		related := []TaskTuple{}

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
		relatedTasks := []TaskTuple{task}
		localVisited := make(map[TaskTuple]bool) // to track visited tasks in this loop iteration
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
func contains(slice []TaskTuple, item TaskTuple) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

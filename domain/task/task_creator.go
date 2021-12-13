// 5.6
package task

import "dddfaq/domain/activityhistory"

type TaskCreator struct {
	taskRepository            TaskRepository
	activityHistoryRepository activityhistory.ActivityHistroryRepository
}

func (tc *TaskCreator) create(taskName string) {
	task := Task{Name: taskName}
	tc.taskRepository.Insert(task)

	activityhistory := activityhistory.CreateActivityHistoryFromTask(task)
	tc.activityHistoryRepository.Insert(activityhistory)
}

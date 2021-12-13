// 5.3
package task

import (
	"dddfaq/domain/activityhistory"
	"dddfaq/domain/task"
)

type CreateTaskUseCase struct {
	taskRepository          task.TaskRepository
	activityHistyRepository activityhistory.ActivityHistroryRepository
}

func (ctu *CreateTaskUseCase) execute(taskName string) {
	task := task.Task{Name: taskName}
	ctu.taskRepository.Insert(task)

	activityHistory := activityhistory.CreateActivityHistoryFromTask(task)
	ctu.activityHistyRepository.Insert(activityHistory)
}

// 3.3
package task

import (
	"errors"
	"time"
)

const maxPostPoneCount = 3

type Task struct {
	Name          string
	postponeCount int
	status        TaskStatus
	dueDate       time.Time
}

func (t *Task) postpone() error {
	if t.postponeCount >= maxPostPoneCount {
		return errors.New("最大延期回数を越えています")
	}
	t.dueDate = t.dueDate.AddDate(0, 0, 1)
	t.postponeCount += 1
	return nil
}

func CreateTask(
	name string,
) Task {
	return Task{
		Name:          name,
		postponeCount: 0,
		status:        TaskStatusToDo,
	}
}

func ReconstractTask(
	name string,
	postponeCount int,
	status TaskStatus,
) Task {
	return Task{
		Name:          name,
		postponeCount: postponeCount,
		status:        status,
	}
}

type TaskStatus int

const (
	TaskStatusInvalid TaskStatus = iota
	TaskStatusToDo
	TaskStatusDone
)

func ValueOfTaskStatus(taskStatus int) TaskStatus {
	if taskStatus == int(TaskStatusToDo) {
		return TaskStatusToDo
	} else if taskStatus == int(TaskStatusDone) {
		return TaskStatusDone
	} else {
		return TaskStatusInvalid
	}
}

// 3.3
package task

type Task struct {
	Name          string
	postponeCount int
	status        TaskStatus
}

func (t *Task) postpone() {
	t.postponeCount++
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

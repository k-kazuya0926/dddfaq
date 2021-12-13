// 5.2
package activityhistory

import (
	"dddfaq/domain/task"
	"fmt"
)

type ActivityHistory struct {
	detail string
}

func CreateActivityHistoryFromTask(task task.Task) ActivityHistory {
	return ActivityHistory{detail: fmt.Sprintf("%sが作成されました", task.Name)}
}

type ActivityHistroryRepository interface {
	Insert(activityHistory ActivityHistory)
}

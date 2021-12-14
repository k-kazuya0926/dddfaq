package task

import (
	"testing"
	"time"
)

func TestTask_postpone(t1 *testing.T) {
	type fields struct {
		Name          string
		postponeCount int
		status        TaskStatus
		dueDate       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "正常系",
			fields: fields{
				postponeCount: 0,
				status:        TaskStatusToDo,
				dueDate:       time.Date(2021, 12, 1, 0, 0, 0, 0, time.Local),
			},
			wantErr: false,
		},
		{
			name: "最大回数オーバー",
			fields: fields{
				postponeCount: maxPostPoneCount,
				status:        TaskStatusToDo,
				dueDate:       time.Date(2021, 12, 1, 0, 0, 0, 0, time.Local),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Task{
				Name:          tt.fields.Name,
				postponeCount: tt.fields.postponeCount,
				status:        tt.fields.status,
				dueDate:       tt.fields.dueDate,
			}
			if err := t.postpone(); (err != nil) != tt.wantErr {
				t1.Errorf("postpone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

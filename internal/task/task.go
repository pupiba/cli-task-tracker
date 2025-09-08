package task

import (
	"cli-task-tracker/pkg/utils"
	"fmt"
	"time"
)

type Task struct {
	id          int
	description string
	status      string
	createdAt   time.Time
	updatedAt   time.Time
}

func NewTask(description string) *Task {
	now := time.Now()
	var id int

	if !utils.FileExists("../../tasks.json") {
		fmt.Println("НЕТУ ФАЙЛА")
		id = 1
	}

	return &Task{
		id:          id,
		description: description,
		status:      "todo",
		createdAt:   now,
		updatedAt:   now,
	}
}

func (t Task) GetTimeNow() string {
	return t.createdAt.Format("15:04 02.01.2006")
}

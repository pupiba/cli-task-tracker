package main

import (
	"cli-task-tracker/internal/task"
	"fmt"
)

func main() {
	t := task.NewTask("Задача 1")
	fmt.Println(t.GetTimeNow())
}

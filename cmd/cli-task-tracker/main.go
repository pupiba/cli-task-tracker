package main

import (
	"cli-task-tracker/internal/command"
	"cli-task-tracker/internal/task"
	"cli-task-tracker/pkg/utils"
	"fmt"
)

func main() {
	// Проверка целостности файлов
	utils.InitStorageFile("tasks.json")

	// Инициализация памяти
	storage, err := task.NewStorageJSON()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Запуск отслеживание команд
	command.Management(storage)
}

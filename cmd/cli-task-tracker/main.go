package main

import (
	"cli-task-tracker/internal/task"
	"cli-task-tracker/pkg/utils"
	"fmt"
)

func main() {
	// TODO: Проверить целостность файла
	utils.InitStorageFile("tasks.json")

	// TODO: Инициализировать память
	storage, err := task.NewStorageJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Тестовый просмотр структуры
	fmt.Println(storage)

	storage.AddToJSON(task.NewTask(3, "Новая задача!"))
	storage.WriteData()

	fmt.Println(storage)

	// TODO: Запустить отслеживание команд
}

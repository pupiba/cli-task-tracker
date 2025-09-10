package command

import (
	"cli-task-tracker/internal/task"
	"cli-task-tracker/pkg/utils"
	"fmt"
	"time"
)

func Add(storage *task.StorageJSON, str string) {
	new_id := storage.GetLastID() + 1
	storage.AddToJSON(task.NewTask(new_id, str))
	fmt.Printf("Task added successfully (ID: %v)\n", new_id)
}

func UpdateDescription(storage *task.StorageJSON, id int, des string) {
	slice_id, err := storage.GetTaskByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	storage.Data[slice_id].Description = des
	storage.Data[slice_id].UpdatedAt = utils.FormattedTime(time.Now())
	fmt.Printf("Task description has been updated (ID: %v)\n", storage.Data[slice_id].ID)
}

func Delete(storage *task.StorageJSON, id int) {
	slice_id, err := storage.GetTaskByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if slice_id == storage.GetLen()-1 {
		storage.Data = storage.Data[:slice_id]
	} else if slice_id == 0 {
		storage.Data = storage.Data[1:]
	} else {
		storage.Data = append(
			storage.Data[:slice_id],
			storage.Data[slice_id+1:]...,
		)
	}
	fmt.Printf("The task has been deleted (ID: %v)\n", id)
}

// Mark a task as in progress or done

func UpdateStatus(storage *task.StorageJSON, id int, status string) {
	slice_id, err := storage.GetTaskByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	storage.Data[slice_id].Status = status
}

// List all tasks

func GetList(storage *task.StorageJSON) {
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	if storage.GetLen() == 0 {
		fmt.Println("- - - - - - - - - - - - | N O   T A S K S | - - - - - - - - - - - - -")
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
		return
	}
	for _, v := range storage.Data {
		fmt.Printf("ID: %v\n\"%v\"\nStatus: %v\nCreated at: %v\nUpdated at: %v\n",
			v.ID, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	}
}

// List all tasks that are done

// List all tasks that are not done

// List all tasks that are in progres

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

func GetList(storage *task.StorageJSON, flag bool, data string) {
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	for _, v := range storage.Data {
		if flag {
			if v.Status == data {
				fmt.Printf("ID: %v\n\"%v\"\nStatus: %v\nCreated at: %v\nUpdated at: %v\n",
					v.ID, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
				fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
			}
			continue
		} else {
			fmt.Printf("ID: %v\n\"%v\"\nStatus: %v\nCreated at: %v\nUpdated at: %v\n",
				v.ID, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
			fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
		}
	}
}

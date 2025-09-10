package command

import (
	"cli-task-tracker/internal/task"
	"fmt"
	"os"
	"strconv"
)

func Management(storage *task.StorageJSON) {
	len_args := len(os.Args)
	if len_args < 2 {
		fmt.Println("there are no arguments")
		return
	}

	switch os.Args[1] {
	case "add":
		if len_args != 3 {
			fmt.Println(`Usage: "./task-cli add "Example task"`)
			return
		}
		Add(storage, os.Args[2])
	case "update":
		if len_args != 4 {
			fmt.Println(`Usage: "./task-cli update <ID> "Example updated task"`)
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		UpdateDescription(storage, id, os.Args[3])
	case "delete":
		if len_args != 3 {
			fmt.Println(`Usage: "./task-cli delete <ID>"`)
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		Delete(storage, id)
	case "mark-in-progress":
		if len_args != 3 {
			fmt.Println(`Usage: "./task-cli mark-in-progress <ID>"`)
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		UpdateStatus(storage, id, "mark-in-progress")
	case "mark-done":
		if len_args != 3 {
			fmt.Println(`Usage: "./task-cli mark-done <ID>"`)
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		UpdateStatus(storage, id, "mark-done")
	case "list":
		switch len_args {
		case 3:
			//
		case 2:
			//
		default:
			fmt.Println(`Usage: "./task-cli list" or "./task-cli <STATUS>"`)
			return
		}
		GetList(storage)
	default:

	}
	storage.WriteData()
}

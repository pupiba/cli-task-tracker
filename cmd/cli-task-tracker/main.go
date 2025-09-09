package main

import (
	"cli-task-tracker/pkg/utils"
	"fmt"
)

func main() {

	utils.InitStorage(fmt.Sprintf("/../../%v", utils.FILENAME))

	// js := task.NewStorageJSON()

	// if err := js.ReadData(); err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(js.GetLen())

	// command.Management()
}

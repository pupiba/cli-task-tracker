package command

import (
	"fmt"
	"os"
)

func Management() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	fmt.Println(os.Args[0])
	switch os.Args[1] {

	case "add":
		//
	case "update":
		//
	case "delete":
		//
	case "mark-in-progress":
		//
	case "mark-done":
		//
	case "list":
		//
	}
}

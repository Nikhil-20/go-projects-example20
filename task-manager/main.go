package main

import (
	"fmt"
	"os"
	"strconv"
	"task-manager/task"
)

func main() {

	fs, err := task.NewFilestore("Newtask.json")

	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Invalid entries from input")
		return
	}

	cmd := os.Args[1]

	switch cmd {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("provide title details")
			return
		}

		ts, err := fs.Add(os.Args[2])
		if err == nil {
			fmt.Println("Added task with ID:", ts.ID)
		} else {
			fmt.Println("Failed adding task")
			return
		}

	case "list":

		tasklist, _ := fs.List()
		for _, ts := range tasklist {
			fmt.Printf("[%d] %s (completed: %v) \n", ts.ID, ts.Title, ts.Completed)
		}

	case "complete":
		id, _ := strconv.Atoi(os.Args[2])
		fs.Complete(id)

	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		fs.Delete(id)

	default:
		fmt.Println("Unknow command")
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

type ToDoList struct {
	tasks []string
}

func (list *ToDoList) addTask(task string) {
	list.tasks = append(list.tasks, task)
	fmt.Printf("Task \"%s\" added to the to-do list.\n", task)
}

func (list *ToDoList) viewTasks() {
	if len(list.tasks) > 0 {
		fmt.Println("Your To-Do List:")
		for index, task := range list.tasks {
			fmt.Printf("%d. %s\n", index+1, task)
		}
	} else {
		fmt.Println("Your to-do list is empty.")
	}
}

func main() {
	todoList := ToDoList{}

	for {
		fmt.Println("\n1. Add Task\n2. View Tasks\n3. Quit")
		fmt.Print("Enter your choice (1/2/3): ")

		var choice string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			choice = scanner.Text()
		}

		switch choice {
		case "1":
			fmt.Print("Enter the task: ")
			var task string
			if scanner.Scan() {
				task = scanner.Text()
			}
			todoList.addTask(task)
		case "2":
			todoList.viewTasks()
		case "3":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		}
	}
}

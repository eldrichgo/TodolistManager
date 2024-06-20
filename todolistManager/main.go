package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func printOptions() {
	fmt.Println()
	fmt.Println("Please choose an option:")
	fmt.Println("1. Add a new task")
	fmt.Println("2. View all tasks")
	fmt.Println("3. Update task status")
	fmt.Println("4. Delete a task")
	fmt.Println("5. Exit")
	fmt.Println()
}

func getOption() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your choice: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again.")
			fmt.Println()
			continue
		}

		input = strings.TrimSpace(input)
		option, err := strconv.Atoi(input)
		if err != nil || option < 1 || option > 5 {
			fmt.Println("Invalid option. Please try again.")
			fmt.Println()
			continue
		}

		return option
	}
}

type Task struct {
	ID     int `gorm:"primaryKey"`
	Title  string
	Status string
}

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=todolist port=5432 sslmode=prefer TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database")

	// Migrate the schema
	db.AutoMigrate(&Task{})

	// Slize of tasks
	tasks := make([]Task, 0)

	fmt.Println("Welcome to the To-Do List Manager!")
	reader := bufio.NewReader(os.Stdin)
	for {
		printOptions()
		choice := getOption()

		switch choice {
		case 1: //Add Task
			fmt.Print("Enter task title: ")
			title, _ := reader.ReadString('\n')

			task := Task{ID: len(tasks) + 1, Title: title, Status: "Pending"}
			tasks = append(tasks, task)

			fmt.Println("Task added successfully!")

		case 2: //View Tasks
			for i := range tasks {
				fmt.Println("ID:", tasks[i].ID, "| Title:", tasks[i].Title, "| Status:", tasks[i].Status)
			}

		case 3: //Update task status
			fmt.Print("Enter task ID: ")
			var ID int
			fmt.Scanln(&ID)

			fmt.Print("Enter new status (Pending/Completed): ")
			newStatus, _ := reader.ReadString('\n')

			// Variable to check if task ID was found
			taskFound := false
			for i := range tasks {
				if tasks[i].ID == ID {
					tasks[i].Status = newStatus
					fmt.Println("Task status updated successfully!")
					taskFound = true
					break
				}
			}

			if !taskFound {
				fmt.Println("Task with specified ID not found.")
			}

		case 4: //Delete Task
			fmt.Print("Enter task ID: ")
			var ID int
			fmt.Scanln(&ID)

			// Variable to check if task was found
			taskFound := false
			for i := range tasks {
				if tasks[i].ID == ID {
					tasks = append(tasks[:i], tasks[i+1:]...)

					/* // Shift elements to the left to overwrite the element at index i
					copy(tasks[i:], tasks[i+1:])
					// Resize the slice to remove the last element
					tasks = tasks[:len(tasks)-1]*/

					fmt.Println("Task deleted successfully!")
					taskFound = true
					break
				}
			}

			if !taskFound {
				fmt.Println("Task with specified ID not found.")
			}

		case 5: //Exit
			fmt.Println("Thank you for using the To-Do List Manager! Goodbye.")
			return
		}
	}

}

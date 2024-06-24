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
	ID        int `gorm:"primaryKey"`
	Title     string
	Status    string
	isDeleted gorm.DeletedAt
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

	fmt.Println("Welcome to the To-Do List Manager!")
	reader := bufio.NewReader(os.Stdin)
	for {
		printOptions()
		choice := getOption()

		switch choice {
		case 1: //Add Task
			fmt.Print("Enter task title: ")
			title, _ := reader.ReadString('\n')

			task := Task{Title: title, Status: "Pending"}
			result := db.Create(&task)

			if result.Error != nil {
				fmt.Println("An error occurred while adding the task. Please try again.")
				continue
			}

			fmt.Println("Task added successfully!")

		case 2: //View Tasks
			var tasks []Task
			result := db.Find(&tasks)

			if result.Error != nil {
				fmt.Println("An error occurred while fetching tasks. Please try again.")
				continue
			}

			for _, task := range tasks {
				fmt.Println("ID:", task.ID, "| Title:", task.Title, "| Status:", task.Status)
			}

		case 3: //Update task status
			var task Task
			fmt.Print("Enter task ID: ")
			var ID int
			fmt.Scanln(&ID)

			result := db.First(&task, ID)

			if result.Error != nil {
				fmt.Println("Task with specified ID not found.")
				continue
			}

			fmt.Print("Enter new status (Pending/Completed): ")
			newStatus, _ := reader.ReadString('\n')
			newStatus = strings.TrimSpace(newStatus)

			task.Status = newStatus
			db.Save(&task)

		case 4: //Delete Task
			var task Task
			fmt.Print("Enter task ID: ")
			var ID int
			fmt.Scanln(&ID)

			result := db.Delete(&task, ID)

			if result.RowsAffected == 0 {
				fmt.Println("Task with specified ID not found.")
				continue
			}

			fmt.Println("Task deleted successfully!")

		case 5: //Exit
			fmt.Println("Thank you for using the To-Do List Manager! Goodbye.")
			return
		}
	}

}

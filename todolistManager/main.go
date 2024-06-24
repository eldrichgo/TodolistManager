package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"type:varchar(100)"`
	Status    string `gorm:"type:varchar(20)"`
	DeletedAt gorm.DeletedAt
	//UpdatedAt gorm.UpdatedAt
}

func initLogger() logger.Interface {
	logLevel := logger.Info
	f, _ := os.Create("gorm.log")
	newLogger := logger.New(
		log.New(
			io.MultiWriter(f, os.Stdout), "\r\n", log.LstdFlags), logger.Config{
			Colorful:                  true,
			LogLevel:                  logLevel,
			SlowThreshold:             time.Second,
			IgnoreRecordNotFoundError: true,
		})

	return newLogger
}

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=todolist port=5432 sslmode=prefer TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: initLogger()})
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

			fmt.Print("Enter new status (Pending/Completed): ")
			newStatus, _ := reader.ReadString('\n')
			newStatus = strings.TrimSpace(newStatus)

			task.ID = ID
			task.Status = newStatus
			result := db.Updates(&task)

			if result.Error != nil {
				fmt.Println("An error occurred while updating the task. Please try again.")
				continue
			}

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

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID     int
	Title  string
	Status string
}

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

func main() {
	fmt.Println("Welcome to the To-Do List Manager!")
	reader := bufio.NewReader(os.Stdin)
	for {
		printOptions()
		choice := getOption()

		switch choice {
		case 1: //Add Task
			fmt.Print("Enter task title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			task := Task{Title: title, Status: "Pending"}

			//convert task to jason
			taskJson, _ := json.Marshal(task)

			resp, err := http.Post("http://localhost:8080/addtask", "application/json", bytes.NewBuffer(taskJson))

			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Println("Error: received non-OK response code:", resp.StatusCode)
				return
			}

			fmt.Println("Task added successfully!")

		case 2: //View Tasks
			var tasks []Task
			resp, err := http.Get("http://localhost:8080/viewtasks")

			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Println("Error: received non-OK response code:", resp.StatusCode)
				return
			}

			err = json.NewDecoder(resp.Body).Decode(&tasks)
			if err != nil {
				fmt.Println("An error occurred while decoding the response:", err)
				return
			}

			for _, task := range tasks {
				fmt.Println("ID:", task.ID, "| Title:", task.Title, "| Status:", task.Status)
			}

		// case 3: //Update task status
		// 	var task Task
		// 	fmt.Print("Enter task ID: ")
		// 	var ID int
		// 	fmt.Scanln(&ID)

		// 	fmt.Print("Enter new status (Pending/Completed): ")
		// 	newStatus, _ := reader.ReadString('\n')
		// 	newStatus = strings.TrimSpace(newStatus)

		// 	task.ID = ID
		// 	task.Status = newStatus
		// 	result := db.Updates(&task)

		// 	if result.Error != nil {
		// 		fmt.Println("An error occurred while updating the task. Please try again.")
		// 		continue
		// 	}

		case 4: //Delete Task
			fmt.Print("Enter task ID: ")
			var ID int
			fmt.Scanln(&ID)

			// Create the request URL
			url := fmt.Sprintf("http://localhost:8080/deletetask/%d", ID)

			// Create the HTTP DELETE request
			req, err := http.NewRequest(http.MethodDelete, url, nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			// Send the HTTP DELETE request
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()

			// Check the response status
			if resp.StatusCode != http.StatusOK {
				var errMsg struct {
					Error string `json:"error"`
				}
				if err := json.NewDecoder(resp.Body).Decode(&errMsg); err == nil {
					fmt.Println("Error:", errMsg.Error)
				} else {
					fmt.Println("Error: received non-OK response code:", resp.StatusCode)
				}
				return
			}

			fmt.Println("Task deleted successfully!")

		case 5: //Exit
			fmt.Println("Thank you for using the To-Do List Manager! Goodbye.")
			return
		}
	}
}

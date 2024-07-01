package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"server/dal/repository"
	"server/dal/service"
	"server/models"
)

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

	// Migrate the schema
	//db.AutoMigrate(&models.Task{})

	// Initialize repository and service
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)

	router := gin.Default()

	router.POST("/addtask", func(c *gin.Context) {
		var task models.Task
		if err := c.ShouldBind(&task); err != nil {
			log.Println("An error occurred while binding the JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request. Please provide a valid task.",
			})
			return
		}

		task.Status = "Pending"
		if err := taskService.CreateTask(&task); err != nil {
			log.Println("An error occurred while adding the task:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "An error occurred while adding the task. Please try again.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"task": task,
		})
	})

	router.GET("/viewtasks", func(c *gin.Context) {
		tasks, err := taskService.GetAllTasks()
		if err != nil {
			log.Println("An error occurred while fetching tasks:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "An error occurred while fetching tasks. Please try again.",
			})
			return
		}

		c.JSON(http.StatusOK, tasks)
	})

	router.PATCH("/updatetask/:id", func(c *gin.Context) {
		id := c.Param("id")
		taskID, _ := strconv.Atoi(id)
		var task models.Task

		if err := c.ShouldBind(&task); err != nil {
			log.Println("An error occurred while binding the JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request. Please provide a valid task.",
			})
			return
		}

		if err := taskService.UpdateTaskStatus(taskID, task.Status); err != nil {
			log.Println("An error occurred while updating the task:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "An error occurred while updating the task. Please try again.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"task": task,
		})
	})

	router.DELETE("/deletetask/:id", func(c *gin.Context) {
		id := c.Param("id")
		taskID, _ := strconv.Atoi(id)

		if err := taskService.DeleteTask(taskID); err != nil {
			log.Println("An error occurred while deleting the task:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "An error occurred while deleting the task. Please try again.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Task deleted successfully!",
		})
	})

	router.Run()
}

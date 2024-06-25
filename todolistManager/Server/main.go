package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

type Task struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"type:varchar(100)"`
	Status    string `gorm:"type:varchar(20)"`
	DeletedAt gorm.DeletedAt
}

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=todolist port=5432 sslmode=prefer TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: initLogger()})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database")

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/addtask", func(c *gin.Context) {
		var task Task
		if err := c.ShouldBind(&task); err != nil {
			log.Println("An error occurred while binding the JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request. Please provide a valid task.",
			})
			return
		}

		result := db.Create(&task)
		if result.Error != nil {
			log.Println("An error occurred while adding the task:", result.Error)
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
		var tasks []Task
		result := db.Find(&tasks)

		if result.Error != nil {
			log.Println("An error occurred while fetching tasks:", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "An error occurred while fetching tasks. Please try again.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"tasks": tasks,
		})
	})

	router.Run()
}

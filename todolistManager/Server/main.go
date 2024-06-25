package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Task struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"type:varchar(100)"`
	Status    string `gorm:"type:varchar(20)"`
	DeletedAt gorm.DeletedAt
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/viewtasks", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "View all tasks",
		})
	})

	router.Run()
}

//CRUD todolist

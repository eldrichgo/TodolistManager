package dataloader

import (
	"context"
	"io"
	"log"
	"os"
	"server/dal/todo"
	"server/graph/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type loaderString string

const loadersKey = loaderString("dataloaders")

type Loaders struct {
	TasksbyUserID TasksbyUserIDLoader
	Tasks         TasksLoader
	UsersbyTaskID UsersbyTaskIDLoader
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

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c, loadersKey, &Loaders{
			TasksbyUserID: TasksbyUserIDLoader{
				maxBatch: 2,
				wait:     500 * time.Millisecond,
				fetch: func(userIDs []int) ([][]*model.Task, []error) {
					dsn := "host=localhost user=postgres password=1234 dbname=todolist port=5432 sslmode=prefer TimeZone=Asia/Shanghai"
					db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: initLogger()})
					if err != nil {
						log.Fatalf("Failed to connect to the database: %v", err)
					}

					svc := todo.NewTodoService(todo.NewTodoRepository(db))
					resp, err := svc.GetTasksbyUserIDs(userIDs) //contains all tasks of all users. duplicate tasks are present
					if err != nil {
						return nil, []error{err}
					}

					tasksbyID := map[int][]*model.Task{}
					for _, usertask := range resp {
						tasksbyID[usertask.UserID] = append(tasksbyID[usertask.UserID], &model.Task{
							ID:     usertask.TaskID,
							Title:  usertask.Title,
							Status: usertask.Status,
						})
					}

					items := make([][]*model.Task, len(userIDs))
					for i, userID := range userIDs {
						items[i] = tasksbyID[userID]
					}

					return items, nil
				},
			},
			UsersbyTaskID: UsersbyTaskIDLoader{
				maxBatch: 2,
				wait:     500 * time.Millisecond,
				fetch: func(taskIDs []int) ([][]*model.User, []error) {
					dsn := "host=localhost user=postgres password=1234 dbname=todolist port=5432 sslmode=prefer TimeZone=Asia/Shanghai"
					db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: initLogger()})
					if err != nil {
						log.Fatalf("Failed to connect to the database: %v", err)
					}

					svc := todo.NewTodoService(todo.NewTodoRepository(db))
					resp, err := svc.GetUsersbyTaskIDs(taskIDs)
					if err != nil {
						return nil, []error{err}
					}

					usersbyID := map[int][]*model.User{}
					for _, usertask := range resp {
						usersbyID[usertask.TaskID] = append(usersbyID[usertask.TaskID], &model.User{
							ID:   usertask.UserID,
							Name: usertask.Name,
						})
					}

					items := make([][]*model.User, len(taskIDs))
					for i, taskID := range taskIDs {
						items[i] = usersbyID[taskID]
					}

					return items, nil
				},
			},
		})

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

package todo

import (
	"server/graph/model"

	"gorm.io/gorm"
)

// TaskRepository defines the methods that any implementation of a task repository must have
type TaskRepository interface {
	CreateTask(task *model.Task) (*model.Task, error)
	FindAllTasks() ([]model.Task, error)
	FindTask(taskID int) (*model.Task, error)
	UpdateTaskStatus(taskID int, status string) (*model.Task, error)
	DeleteTask(taskID int) error
}

// Task is the implementation of TaskRepositoryInterface
type Task struct {
	db *gorm.DB
}

// NewTaskRepository creates a new instance of TaskRepository
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &Task{db: db}
}

func (r *Task) CreateTask(task *model.Task) (*model.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *Task) FindAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Where("deleted_at is null").Find(&tasks).Error
	return tasks, err
}

func (r *Task) FindTask(taskID int) (*model.Task, error) {
	var task *model.Task
	if err := r.db.First(&task, taskID).Error; err != nil {
		return nil, err
	}

	return task, nil
}
func (r *Task) UpdateTaskStatus(taskID int, status string) (*model.Task, error) {
	var task *model.Task

	if err := r.db.Model(&model.Task{}).Where("id = ?", taskID).Update("status", status).Error; err != nil {
		return nil, err
	}

	if err := r.db.First(&task, taskID).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *Task) DeleteTask(taskID int) error {
	return r.db.Delete(&model.Task{}, taskID).Error
	// time.Time.UTC()
}

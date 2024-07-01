package repository

import (
	"server/models"

	"gorm.io/gorm"
)

// TaskRepository defines the methods that any implementation of a task repository must have
type TaskRepository interface {
	Create(task *models.Task) error
	FindAll() ([]models.Task, error)
	UpdateStatus(taskID int, status string) error
	Delete(taskID int) error
}

// Task is the implementation of TaskRepositoryInterface
type Task struct {
	db *gorm.DB
}

// NewTaskRepository creates a new instance of TaskRepository
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &Task{db: db}
}

func (r *Task) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *Task) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *Task) UpdateStatus(taskID int, status string) error {
	return r.db.Model(&models.Task{}).Where("id = ?", taskID).Update("status", status).Error
}

func (r *Task) Delete(taskID int) error {
	return r.db.Delete(&models.Task{}, taskID).Error
}

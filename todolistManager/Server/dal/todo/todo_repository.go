package todo

import (
	"server/graph/model"

	"gorm.io/gorm"
)

// TaskRepository defines the methods that any implementation of a task repository must have
type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	FindAll() ([]model.Task, error)
	UpdateStatus(taskID int, status string) (*model.Task, error)
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

func (r *Task) Create(task *model.Task) (*model.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *Task) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *Task) UpdateStatus(taskID int, status string) (*model.Task, error) {
	var task *model.Task

	if err := r.db.Model(&model.Task{}).Where("id = ?", taskID).Update("status", status).Error; err != nil {
		return nil, err
	}

	if err := r.db.First(&task, taskID).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *Task) Delete(taskID int) error {
	return r.db.Delete(&model.Task{}, taskID).Error
}

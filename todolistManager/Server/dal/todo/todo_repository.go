package todo

import (
	"server/graph/model"

	"gorm.io/gorm"
)

// TodoRepository defines the methods that any implementation of a task repository must have
type TodoRepository interface {
	CreateTask(task *model.Task) (*model.Task, error)
	FindAllTasks() ([]model.Task, error)
	FindTask(taskID int) (*model.Task, error)
	UpdateTaskStatus(taskID int, status string) (*model.Task, error)
	DeleteTask(taskID int) error

	CreateUser(user *model.User) (*model.User, error)
	FindAllUsers() ([]model.User, error)
	FindUser(userID int) (*model.User, error)
	UpdateUserName(userID int, name string) (*model.User, error)
	DeleteUser(userID int) error
}

// Todo is the implementation of TodoRepositoryInterface
type Todo struct {
	db *gorm.DB
}

// NewTodoRepository creates a new instance of TaskRepository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &Todo{db: db}
}

func (r *Todo) CreateTask(task *model.Task) (*model.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *Todo) FindAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Where("deleted_at is null").Find(&tasks).Error
	return tasks, err
}

func (r *Todo) FindTask(taskID int) (*model.Task, error) {
	var task *model.Task
	if err := r.db.First(&task, taskID).Error; err != nil {
		return nil, err
	}

	return task, nil
}
func (r *Todo) UpdateTaskStatus(taskID int, status string) (*model.Task, error) {
	var task *model.Task

	if err := r.db.Model(&model.Task{}).Where("id = ?", taskID).Update("status", status).Error; err != nil {
		return nil, err
	}

	if err := r.db.First(&task, taskID).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *Todo) DeleteTask(taskID int) error {
	return r.db.Delete(&model.Task{}, taskID).Error
	// time.Time.UTC()
}

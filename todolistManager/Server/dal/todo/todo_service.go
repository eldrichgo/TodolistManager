package todo

import "server/graph/model"

// TodoService defines the methods that any implementation of a todo service must have
type TodoService struct {
	repo TodoRepository
}

// NewTodoService creates a new instance of TodoService
func NewTodoService(repo TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTask(taskinput model.InputTask) (*model.Task, error) {
	if taskinput.Status == "" {
		taskinput.Status = "Pending"
	}

	task := &model.Task{
		Title:  taskinput.Title,
		Status: taskinput.Status,
	}

	return s.repo.CreateTask(task)
}

func (s *TodoService) GetAllTasks() ([]model.Task, error) {
	return s.repo.FindAllTasks()
}

func (s *TodoService) GetTask(taskID int) (*model.Task, error) {
	return s.repo.FindTask(taskID)
}

func (s *TodoService) UpdateTaskStatus(taskID int, status string) (*model.Task, error) {
	return s.repo.UpdateTaskStatus(taskID, status)
}

func (s *TodoService) DeleteTask(taskID int) error {
	return s.repo.DeleteTask(taskID)
}

package todo

import "server/graph/model"

// TaskService is the implementation of TaskServiceInterface
type TaskService struct {
	repo TaskRepository
}

// NewTaskService creates a new instance of TaskService
func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(taskinput model.InputTask) (*model.Task, error) {
	if taskinput.Status == "" {
		taskinput.Status = "Pending"
	}

	task := &model.Task{
		Title:  taskinput.Title,
		Status: taskinput.Status,
	}

	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.repo.FindAllTasks()
}

func (s *TaskService) GetTask(taskID int) (*model.Task, error) {
	return s.repo.FindTask(taskID)
}

func (s *TaskService) UpdateTaskStatus(taskID int, status string) (*model.Task, error) {
	return s.repo.UpdateTaskStatus(taskID, status)
}

func (s *TaskService) DeleteTask(taskID int) error {
	return s.repo.DeleteTask(taskID)
}

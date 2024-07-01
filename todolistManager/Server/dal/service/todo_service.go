package service

import (
	"server/dal/repository"
	"server/models"
)

// TaskService is the implementation of TaskServiceInterface
type TaskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a new instance of TaskService
func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.repo.Create(task)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.FindAll()
}

func (s *TaskService) UpdateTaskStatus(taskID int, status string) error {
	return s.repo.UpdateStatus(taskID, status)
}

func (s *TaskService) DeleteTask(taskID int) error {
	return s.repo.Delete(taskID)
}

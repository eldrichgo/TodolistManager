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
	if status == "" {
		status = "Pending"
	}

	return s.repo.UpdateTaskStatus(taskID, status)
}

func (s *TodoService) DeleteTask(taskID int) error {
	return s.repo.DeleteTask(taskID)
}

func (s *TodoService) GetAllUsersOfTask(taskID int) ([]model.User, error) {
	return s.repo.FindUsersofTask(taskID)
}

func (s *TodoService) CreateUser(name string) (*model.User, error) {
	if name == "" {
		return nil, nil
	}

	user := &model.User{
		Name: name,
	}

	return s.repo.CreateUser(user)
}

func (s *TodoService) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAllUsers()
}

func (s *TodoService) GetUser(userID int) (*model.User, error) {
	return s.repo.FindUser(userID)
}

func (s *TodoService) UpdateUserName(userID int, name string) (*model.User, error) {
	if name == "" {
		return nil, nil
	}

	return s.repo.UpdateUserName(userID, name)
}

func (s *TodoService) DeleteUser(userID int) error {
	return s.repo.DeleteUser(userID)
}

func (s *TodoService) GetAllTasksofUser(userID int) ([]model.Task, error) {
	return s.repo.FindTasksofUser(userID)
}

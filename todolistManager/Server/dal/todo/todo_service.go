package todo

import (
	"errors"
	"regexp"
	"server/graph/model"
)

// TodoService defines the methods that any implementation of a todo service must have
type TodoService struct {
	repo TodoRepository
}

// NewTodoService creates a new instance of TodoService
func NewTodoService(repo TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

// Helper function to check if a string contains numbers
func containsNumbers(s string) bool {
	re := regexp.MustCompile(`[0-9]`)
	return re.MatchString(s)
}

func (s *TodoService) CreateTask(taskinput model.InputTask) (*model.Task, error) {
	status := "Pending"
	if taskinput.Status == nil || *taskinput.Status == "" {
		taskinput.Status = &status
	}
	if *taskinput.Status != "Pending" && *taskinput.Status != "Completed" && *taskinput.Status != "In Progress" {
		return nil, errors.New("invalid status")
	}

	task := &model.Task{
		Title:  taskinput.Title,
		Status: *taskinput.Status,
	}

	return s.repo.CreateTask(task)
}

func (s *TodoService) GetAllTasks() ([]model.Task, error) {
	return s.repo.FindAllTasks()
}

func (s *TodoService) GetTask(taskID int) (*model.Task, error) {
	if taskID <= 0 {
		return nil, errors.New("invalid task id")
	}

	return s.repo.FindTask(taskID)
}

func (s *TodoService) GetTasksbyIDs(taskIDs []int) ([]model.Task, error) {
	return s.repo.FindTasksbyID(taskIDs)
}

func (s *TodoService) UpdateTaskStatus(taskID int, status string) (*model.Task, error) {
	if taskID <= 0 {
		return nil, errors.New("invalid task id")
	}

	if status != "Pending" && status != "Completed" && status != "In Progress" {
		return nil, errors.New("invalid status")
	}

	return s.repo.UpdateTaskStatus(taskID, status)
}

func (s *TodoService) DeleteTask(taskID int) error {
	if taskID <= 0 {
		return errors.New("invalid task id")
	}

	return s.repo.DeleteTask(taskID)
}

func (s *TodoService) GetAllUsersOfTask(taskID int) ([]model.User, error) {
	if taskID <= 0 {
		return nil, errors.New("invalid task id")
	}
	return s.repo.FindUsersofTask(taskID)
}

func (s *TodoService) GetTasksbyUserIDs(userIDs []int) ([]*model.UserTask, error) {
	return s.repo.FindTasksbyUserIDs(userIDs)
}

func (s *TodoService) CreateUser(name string) (*model.User, error) {
	if name == "" {
		return nil, errors.New("invalid name")
	}

	if containsNumbers(name) {
		return nil, errors.New("invalid name")
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
	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	return s.repo.FindUser(userID)
}

func (s *TodoService) UpdateUserName(userID int, name string) (*model.User, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	if name == "" {
		return nil, errors.New("invalid name")
	}

	if containsNumbers(name) {
		return nil, errors.New("invalid name")
	}

	return s.repo.UpdateUserName(userID, name)
}

func (s *TodoService) DeleteUser(userID int) error {
	if userID <= 0 {
		return errors.New("invalid user id")
	}

	return s.repo.DeleteUser(userID)
}

func (s *TodoService) GetAllTasksofUser(userID int) ([]model.Task, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	return s.repo.FindTasksofUser(userID)
}

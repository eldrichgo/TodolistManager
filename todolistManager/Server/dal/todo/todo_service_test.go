package todo

import (
	"errors"
	"server/graph/model"
	"testing"

	"github.com/stretchr/testify/mock"
)

type taskRepoMock struct {
	mock.Mock
}

// Mock functions for Task operations in the repository
func (m *taskRepoMock) CreateTask(task *model.Task) (*model.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (m *taskRepoMock) FindAllTasks() ([]model.Task, error) {
	args := m.Called()
	return args.Get(0).([]model.Task), args.Error(1)
}

func (m *taskRepoMock) FindTask(taskID int) (*model.Task, error) {
	args := m.Called(taskID)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (m *taskRepoMock) UpdateTaskStatus(taskID int, status string) (*model.Task, error) {
	args := m.Called(taskID, status)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (m *taskRepoMock) DeleteTask(taskID int) error {
	args := m.Called(taskID)
	return args.Error(0)
}

func (m *taskRepoMock) FindUsersofTask(taskID int) ([]model.User, error) {
	args := m.Called(taskID)
	return args.Get(0).([]model.User), args.Error(1)
}

// Mock functions for User operations in the repository
func (m *taskRepoMock) CreateUser(user *model.User) (*model.User, error) {
	args := m.Called(user)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *taskRepoMock) FindAllUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *taskRepoMock) FindUser(userID int) (*model.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *taskRepoMock) UpdateUserName(userID int, name string) (*model.User, error) {
	args := m.Called(userID, name)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *taskRepoMock) DeleteUser(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *taskRepoMock) FindTasksofUser(userID int) ([]model.Task, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.Task), args.Error(1)
}

func TestCreateTask(t *testing.T) {
	var tests = []struct {
		name     string
		input    model.InputTask
		repoMock func() *taskRepoMock
		wantErr  bool
	}{
		{
			name: "Success",
			input: model.InputTask{
				Title:  "Test Task",
				Status: "Pending",
			},
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("CreateTask", mock.Anything).Return(nil)
				return m
			},
			wantErr: false,
		},
		{
			name: "Error",
			input: model.InputTask{
				Title:  "Test Task",
				Status: "Pending",
			},
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("CreateTask", mock.Anything).Return(errors.New("Error"))
				return m
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())

			_, err := svc.CreateTask(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

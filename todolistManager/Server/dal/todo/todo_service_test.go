package todo

import (
	"errors"
	"server/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
	status := "Pending"
	var tests = []struct {
		name         string
		repoMock     func() *taskRepoMock
		input        model.InputTask
		expectedTask *model.Task
		expectedErr  bool
	}{
		{
			name: "Success",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("CreateTask", mock.Anything).Return(&model.Task{
					ID:     1,
					Title:  "Test",
					Status: "Pending",
				}, nil)
				return m
			},
			input: model.InputTask{
				Title:  "Test",
				Status: &status,
			},
			expectedTask: &model.Task{
				ID:     1,
				Title:  "Test",
				Status: "Pending",
			},
			expectedErr: false,
		},
		{
			name: "Error",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("CreateTask", mock.Anything).Return(&model.Task{}, errors.New("error"))
				return m
			},
			input: model.InputTask{
				Title:  "Test",
				Status: &status,
			},
			expectedTask: nil,
			expectedErr:  true,
		},
		{
			name: "Invalid Status",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("CreateTask", mock.Anything).Return(&model.Task{}, errors.New("invalid status"))
				return m
			},
			input: model.InputTask{
				Title:  "Test",
				Status: &[]string{"Blocked"}[0],
			},
			expectedTask: nil,
			expectedErr:  true,
		},
		{
			name: "Valid Status",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("CreateTask", mock.Anything).Return(&model.Task{
					ID:     1,
					Title:  "Test",
					Status: "Pending"}, nil)
				return m
			},
			input: model.InputTask{
				Title:  "Test",
				Status: &status,
			},
			expectedTask: &model.Task{
				ID:     1,
				Title:  "Test",
				Status: "Pending",
			},
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			task, err := svc.CreateTask(tt.input)

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.NotEmpty(t, tt.input.Status)
				require.Nil(t, err)
				assert.Equal(t, tt.expectedTask, task)
			}

			tt.repoMock().AssertExpectations(t)
		})
	}
}

func TestGetAllTasks(t *testing.T) {
	var tests = []struct {
		name          string
		repoMock      func() *taskRepoMock
		expectedTasks []model.Task
		expectedErr   bool
	}{
		{
			name: "Get All Tasks",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindAllTasks").Return([]model.Task{
					{
						ID:     1,
						Title:  "Test",
						Status: "Pending",
					},
				}, nil)
				return m
			},
			expectedTasks: []model.Task{
				{
					ID:     1,
					Title:  "Test",
					Status: "Pending",
				},
			},
			expectedErr: false,
		},
		{
			name: "Error",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindAllTasks").Return([]model.Task{}, errors.New("error"))
				return m
			},
			expectedTasks: nil,
			expectedErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			tasks, err := svc.GetAllTasks()

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				assert.Equal(t, tt.expectedTasks, tasks)
			}
		})
	}
}

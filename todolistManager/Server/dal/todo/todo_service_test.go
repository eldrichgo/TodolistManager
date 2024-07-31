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
				m.On("CreateTask", mock.Anything).Return(&model.Task{}, nil)
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
			name: "Blank Status",
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
				Status: &[]string{""}[0],
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
				//require.NotEmpty(t, tt.input.Status)
				require.Nil(t, err)
				assert.Equal(t, tt.expectedTask, task)
			}

			//tt.repoMock().AssertExpectations(t)
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
			name: "Success",
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
		{
			name: "Empty Return",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindAllTasks").Return([]model.Task{}, nil)
				return m
			},
			expectedTasks: []model.Task{},
			expectedErr:   false,
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

func TestGetTask(t *testing.T) {
	var tests = []struct {
		name         string
		repoMock     func() *taskRepoMock
		input_taskID int
		expectedTask *model.Task
		expectedErr  bool
	}{
		{
			name: "Success",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindTask", 1).Return(&model.Task{
					ID:     1,
					Title:  "Test",
					Status: "Pending",
				}, nil)
				return m
			},
			input_taskID: 1,
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
				m.On("FindTask", 1).Return(&model.Task{}, errors.New("error"))
				return m
			},
			input_taskID: 1,
			expectedTask: nil,
			expectedErr:  true,
		},
		{
			name: "Negative Range",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindTask", -1).Return(&model.Task{}, nil)
				return m
			},
			input_taskID: -1,
			expectedTask: nil,
			expectedErr:  true,
		},
		// {
		// 	name: "None Integer",
		// 	repoMock: func() *taskRepoMock {
		// 		m := &taskRepoMock{}
		// 		m.On("FindTask", 'a').Return(&model.Task{}, errors.New("record not found"))
		// 		return m
		// 	},
		// 	input_taskID: 'a',
		// 	expectedTask: nil,
		// 	expectedErr:  true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			task, err := svc.GetTask(tt.input_taskID)

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				assert.Equal(t, tt.expectedTask, task)
			}
		})
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	var tests = []struct {
		name         string
		repoMock     func() *taskRepoMock
		input_taskID int
		input_status string
		expectedTask *model.Task
		expectedErr  bool
	}{
		{
			name: "Success",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("UpdateTaskStatus", 1, "Completed").Return(&model.Task{
					ID:     1,
					Title:  "Test",
					Status: "Completed",
				}, nil)
				return m
			},
			input_taskID: 1,
			input_status: "Completed",
			expectedTask: &model.Task{
				ID:     1,
				Title:  "Test",
				Status: "Completed",
			},
			expectedErr: false,
		},
		{
			name: "Error",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("UpdateTaskStatus", 1, "Completed").Return(&model.Task{}, errors.New("error"))
				return m
			},
			input_taskID: 1,
			input_status: "Completed",
			expectedTask: nil,
			expectedErr:  true,
		},
		{
			name: "Invalid Status",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("UpdateTaskStatus", 1, "Ayaw ko").Return(&model.Task{}, nil)
				return m
			},
			input_taskID: 1,
			input_status: "Ayaw ko",
			expectedTask: nil,
			expectedErr:  true,
		},
		{
			name: "Negative Range",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("UpdateTaskStatus", -1, "Completed").Return(&model.Task{}, nil)
				return m
			},
			input_taskID: -1,
			input_status: "Completed",
			expectedTask: nil,
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			task, err := svc.UpdateTaskStatus(tt.input_taskID, tt.input_status)

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				assert.Equal(t, tt.expectedTask, task)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	var tests = []struct {
		name         string
		repoMock     func() *taskRepoMock
		input_taskID int
		expectedErr  bool
	}{
		{
			name: "Success",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("DeleteTask", 1).Return(nil)
				return m
			},
			input_taskID: 1,
			expectedErr:  false,
		},
		{
			name: "Error",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("DeleteTask", 1).Return(errors.New("error"))
				return m
			},
			input_taskID: 1,
			expectedErr:  true,
		},
		{
			name: "Negative Range",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("DeleteTask", -1).Return(nil)
				return m
			},
			input_taskID: -1,
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			err := svc.DeleteTask(tt.input_taskID)

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestGetAllUsersofTask(t *testing.T) {
	var tests = []struct {
		name          string
		repoMock      func() *taskRepoMock
		input_taskID  int
		expectedUsers []model.User
		expectedErr   bool
	}{
		{
			name: "Success",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindUsersofTask", 1).Return([]model.User{
					{
						ID:   1,
						Name: "TestUser",
						// tasks: []*model.Task{}?
					},
				}, nil)
				return m
			},
			input_taskID: 1,
			expectedUsers: []model.User{
				{
					ID:   1,
					Name: "TestUser",
				},
			},
			expectedErr: false,
		},
		{
			name: "Error",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindUsersofTask", 1).Return([]model.User{}, errors.New("error"))
				return m
			},
			input_taskID:  1,
			expectedUsers: nil,
			expectedErr:   true,
		},
		{
			name: "Negative Range",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindUsersofTask", -1).Return([]model.User{}, nil)
				return m
			},
			input_taskID:  -1,
			expectedUsers: nil,
			expectedErr:   true,
		},
		{
			name: "Empty Return",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindUsersofTask", 1).Return([]model.User{}, nil)
				return m
			},
			input_taskID:  1,
			expectedUsers: []model.User{},
			expectedErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			users, err := svc.GetAllUsersOfTask(tt.input_taskID)

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				assert.Equal(t, tt.expectedUsers, users)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	var tests = []struct {
		name         string
		repoMock     func() *taskRepoMock
		input_name   string
		expectedUser *model.User
		expectedErr  bool
	}{
		{
			name: "Success",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("CreateUser", mock.Anything).Return(&model.User{
					ID:   1,
					Name: "TestUser",
				}, nil)
				return m
			},
			input_name: "TestUser",
			expectedUser: &model.User{
				ID:   1,
				Name: "TestUser",
			},
			expectedErr: false,
		},
		{
			name: "Error",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("CreateUser", mock.Anything).Return(&model.User{}, errors.New("error"))
				return m
			},
			input_name:   "TestUser",
			expectedUser: nil,
			expectedErr:  true,
		},
		{
			name: "Blank Name",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("CreateUser", mock.Anything).Return(&model.User{}, nil)
				return m
			},
			input_name:   "",
			expectedUser: nil,
			expectedErr:  true,
		},
		{
			name: "Name with Numbers",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("CreateUser", mock.Anything).Return(&model.User{}, nil)
				return m
			},
			input_name:   "TestUser1",
			expectedUser: nil,
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			user, err := svc.CreateUser(tt.input_name)

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				assert.Equal(t, tt.expectedUser, user)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	var tests = []struct {
		name         string
		repoMock     func() *taskRepoMock
		expectedUser []model.User
		expectedErr  bool
	}{
		{
			name: "Success",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindAllUsers").Return([]model.User{
					{
						ID:   1,
						Name: "TestUser",
					},
				}, nil)
				return m
			},
			expectedUser: []model.User{
				{
					ID:   1,
					Name: "TestUser",
				},
			},
			expectedErr: false,
		},
		{
			name: "Error",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindAllUsers").Return([]model.User{}, errors.New("error"))
				return m
			},
			expectedUser: nil,
			expectedErr:  true,
		},
		{
			name: "Empty Return",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindAllUsers").Return([]model.User{}, nil)
				return m
			},
			expectedUser: []model.User{},
			expectedErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			users, err := svc.GetAllUsers()

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				assert.Equal(t, tt.expectedUser, users)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	var tests = []struct {
		name         string
		repoMock     func() *taskRepoMock
		input_userID int
		expectedUser *model.User
		expectedErr  bool
	}{
		{
			name: "Success",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindUser", 1).Return(&model.User{
					ID:   1,
					Name: "TestUser",
				}, nil)
				return m
			},
			input_userID: 1,
			expectedUser: &model.User{
				ID:   1,
				Name: "TestUser",
			},
			expectedErr: false,
		},
		{
			name: "Error",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindUser", 1).Return(&model.User{}, errors.New("error"))
				return m
			},
			input_userID: 1,
			expectedUser: nil,
			expectedErr:  true,
		},
		{
			name: "Negative Range",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindUser", -1).Return(&model.User{}, nil)
				return m
			},
			input_userID: -1,
			expectedUser: nil,
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			user, err := svc.GetUser(tt.input_userID)

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				assert.Equal(t, tt.expectedUser, user)
			}
		})
	}
}

func TestUpdateUserName(t *testing.T) {
	var tests = []struct {
		name         string
		repoMock     func() *taskRepoMock
		input_userID int
		input_name   string
		expectedUser *model.User
		expectedErr  bool
	}{
		{
			name: "Success",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("UpdateUserName", 1, "NewName").Return(&model.User{
					ID:   1,
					Name: "NewName",
				}, nil)
				return m
			},
			input_userID: 1,
			input_name:   "NewName",
			expectedUser: &model.User{
				ID:   1,
				Name: "NewName",
			},
			expectedErr: false,
		},
		{
			name: "Error",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("UpdateUserName", 1, "NewName").Return(&model.User{}, errors.New("error"))
				return m
			},
			input_userID: 1,
			input_name:   "NewName",
			expectedUser: nil,
			expectedErr:  true,
		},
		{
			name: "Blank Name",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("UpdateUserName", 1, "").Return(&model.User{}, nil)
				return m
			},
			input_userID: 1,
			input_name:   "",
			expectedUser: nil,
			expectedErr:  true,
		},
		{
			name: "Name with Numbers",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("UpdateUserName", 1, "NewName1").Return(&model.User{}, nil)
				return m
			},
			input_userID: 1,
			input_name:   "NewName1",
			expectedUser: nil,
			expectedErr:  true,
		},
		{
			name: "Negative Range",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("UpdateUserName", -1, "NewName").Return(&model.User{}, nil)
				return m
			},
			input_userID: -1,
			input_name:   "NewName",
			expectedUser: nil,
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			user, err := svc.UpdateUserName(tt.input_userID, tt.input_name)

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				assert.Equal(t, tt.expectedUser, user)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	var tests = []struct {
		name         string
		repoMock     func() *taskRepoMock
		input_userID int
		expectedErr  bool
	}{
		{
			name: "Success",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("DeleteUser", 1).Return(nil)
				return m
			},
			input_userID: 1,
			expectedErr:  false,
		},
		{
			name: "Error",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("DeleteUser", 1).Return(errors.New("error"))
				return m
			},
			input_userID: 1,
			expectedErr:  true,
		},
		{
			name: "Negative Range",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("DeleteUser", -1).Return(nil)
				return m
			},
			input_userID: -1,
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			err := svc.DeleteUser(tt.input_userID)

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestGetAllTasksofUser(t *testing.T) {
	var tests = []struct {
		name          string
		repoMock      func() *taskRepoMock
		input_userID  int
		expectedTasks []model.Task
		expectedErr   bool
	}{
		{
			name: "Success",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindTasksofUser", 1).Return([]model.Task{
					{
						ID:     1,
						Title:  "Test",
						Status: "Pending",
					},
				}, nil)
				return m
			},
			input_userID: 1,
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
				m.On("FindTasksofUser", 1).Return([]model.Task{}, errors.New("error"))
				return m
			},
			input_userID:  1,
			expectedTasks: nil,
			expectedErr:   true,
		},
		{
			name: "Negative Range",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindTasksofUser", -1).Return([]model.Task{}, nil)
				return m
			},
			input_userID:  -1,
			expectedTasks: nil,
			expectedErr:   true,
		},
		{
			name: "Empty Return",
			repoMock: func() *taskRepoMock {
				m := &taskRepoMock{}
				m.On("FindTasksofUser", 1).Return([]model.Task{}, nil)
				return m
			},
			input_userID:  1,
			expectedTasks: []model.Task{},
			expectedErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTodoService(tt.repoMock())
			tasks, err := svc.GetAllTasksofUser(tt.input_userID)

			if tt.expectedErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				assert.Equal(t, tt.expectedTasks, tasks)
			}
		})
	}
}

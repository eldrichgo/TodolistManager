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
	FindTasksbyID(taskIDs []int) ([]model.Task, error)
	UpdateTaskStatus(taskID int, status string) (*model.Task, error)
	DeleteTask(taskID int) error
	FindUsersofTask(taskID int) ([]model.User, error)
	FindTasksbyUserIDs(userIDs []int) ([]*model.UserTask, error)

	CreateUser(user *model.User) (*model.User, error)
	FindAllUsers() ([]model.User, error)
	FindUser(userID int) (*model.User, error)
	UpdateUserName(userID int, name string) (*model.User, error)
	DeleteUser(userID int) error
	FindTasksofUser(userID int) ([]model.Task, error)
	FindUsersbyTaskIDs(taskIDs []int) ([]*model.UserTask, error)
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
	if err := r.db.Where("deleted_at IS NULL").Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Todo) FindTask(taskID int) (*model.Task, error) {
	var task *model.Task
	if err := r.db.Where("deleted_at IS NULL").First(&task, taskID).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *Todo) FindTasksbyID(taskIDs []int) ([]model.Task, error) {
	var tasks []model.Task
	if err := r.db.Where("deleted_at IS NULL and id IN ?", taskIDs).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Todo) UpdateTaskStatus(taskID int, status string) (*model.Task, error) {
	var task *model.Task

	if err := r.db.Model(&model.Task{}).Where("deleted_at IS NULL and id = ?", taskID).Update("status", status).Error; err != nil {
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

func (r *Todo) FindUsersofTask(taskID int) ([]model.User, error) {
	var users []model.User
	if err := r.db.Model(&model.User{}).Where("users_tasks.tasks_id", taskID).
		Joins("JOIN users_tasks ON users_tasks.users_id = users.id").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Todo) FindTasksbyUserIDs(userIDs []int) ([]*model.UserTask, error) {
	var tasks []*model.UserTask
	err := r.db.Model(&model.Task{}).Select("users_tasks.user_id, tasks.*").
		Where("users_tasks.user_id IN ?", userIDs).
		Joins("JOIN users_tasks ON users_tasks.task_id = tasks.id").
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, err
}

func (r *Todo) CreateUser(user *model.User) (*model.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Todo) FindAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.db.Where("deleted_at IS NULL").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Todo) FindUser(userID int) (*model.User, error) {
	var user *model.User
	if err := r.db.Where("deleted_at IS NULL").First(&user, userID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Todo) UpdateUserName(userID int, name string) (*model.User, error) {
	var user *model.User

	if err := r.db.Model(&model.User{}).Where("deleted_at IS NULL and id = ?", userID).Update("name", name).Error; err != nil {
		return nil, err
	}

	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Todo) DeleteUser(userID int) error {
	return r.db.Delete(&model.User{}, userID).Error
}

func (r *Todo) FindTasksofUser(userID int) ([]model.Task, error) {
	var tasks []model.Task
	if err := r.db.Model(&model.Task{}).Where("users_tasks.users_id", userID).
		Joins("JOIN users_tasks ON users_tasks.tasks_id = tasks.id").
		Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Todo) FindUsersbyTaskIDs(taskIDs []int) ([]*model.UserTask, error) {
	var users []*model.UserTask
	err := r.db.Model(&model.User{}).Select("users_tasks.task_id, users.*").
		Where("users_tasks.task_id IN ?", taskIDs).
		Joins("JOIN users_tasks ON users_tasks.user_id = users.id").
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, err
}

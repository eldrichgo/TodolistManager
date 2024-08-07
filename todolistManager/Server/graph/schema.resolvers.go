package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"server/dal/todo"
	"server/dataloader"
	"server/graph/model"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.InputTask) (*model.Task, error) {
	svc := todo.NewTodoService(todo.NewTodoRepository(r.Db))
	task, err := svc.CreateTask(input)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// UpdateTaskStatus is the resolver for the updateTaskStatus field.
func (r *mutationResolver) UpdateTaskStatus(ctx context.Context, id int, status string) (*model.Task, error) {
	svc := todo.NewTodoService(todo.NewTodoRepository(r.Db))
	task, err := svc.UpdateTaskStatus(id, status)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id int) (*bool, error) {
	svc := todo.NewTodoService(todo.NewTodoRepository(r.Db))
	err := svc.DeleteTask(id)
	if err != nil {
		return nil, err
	}

	success := true
	return &success, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*model.User, error) {
	svc := todo.NewTodoService(todo.NewTodoRepository(r.Db))
	user, err := svc.CreateUser(name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserName is the resolver for the updateUserName field.
func (r *mutationResolver) UpdateUserName(ctx context.Context, id int, name string) (*model.User, error) {
	svc := todo.NewTodoService(todo.NewTodoRepository(r.Db))
	user, err := svc.UpdateUserName(id, name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*bool, error) {
	svc := todo.NewTodoService(todo.NewTodoRepository(r.Db))
	err := svc.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	success := true
	return &success, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	svc := todo.NewTodoService(todo.NewTodoRepository(r.Db))
	tasks, err := svc.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var result []*model.Task
	for _, task := range tasks {
		result = append(result, &model.Task{
			ID:     task.ID,
			Title:  task.Title,
			Status: task.Status,
		})
	}

	return result, nil
	//return dataloader.For(ctx).Tasks.LoadAll()
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id int) (*model.Task, error) {
	svc := todo.NewTodoService(todo.NewTodoRepository(r.Db))
	task, err := svc.GetTask(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	svc := todo.NewTodoService(todo.NewTodoRepository(r.Db))
	users, err := svc.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var result []*model.User
	for _, user := range users {
		result = append(result, &model.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return result, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	svc := todo.NewTodoService(todo.NewTodoRepository(r.Db))
	user, err := svc.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Users is the resolver for the users field.
func (r *taskResolver) Users(ctx context.Context, obj *model.Task) ([]*model.User, error) {
	/*svc := todo.NewTodoService(todo.NewTodoRepository(r.Db))
	users, err := svc.GetAllUsersOfTask(obj.ID)
	if err != nil {
		return nil, err
	}

	var result []*model.User
	for _, user := range users {
		result = append(result, &model.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return result, nil*/
	return dataloader.For(ctx).UsersbyTaskID.Load(obj.ID)
}

// Tasks is the resolver for the tasks field.
func (r *userResolver) Tasks(ctx context.Context, obj *model.User) ([]*model.Task, error) {
	return dataloader.For(ctx).TasksbyUserID.Load(obj.ID)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Task returns TaskResolver implementation.
func (r *Resolver) Task() TaskResolver { return &taskResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

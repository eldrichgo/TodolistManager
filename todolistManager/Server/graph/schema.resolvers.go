package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"server/dal/todo"
	"server/graph/model"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.InputTask) (*model.Task, error) {
	svc := todo.NewTaskService(todo.NewTaskRepository(r.Db))
	task, err := svc.CreateTask(input)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, id int, status string) (*model.Task, error) {
	svc := todo.NewTaskService(todo.NewTaskRepository(r.Db))
	task, err := svc.UpdateTaskStatus(id, status)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id int) (*bool, error) {
	svc := todo.NewTaskService(todo.NewTaskRepository(r.Db))
	err := svc.DeleteTask(id)
	if err != nil {
		return nil, err
	}
	success := true
	return &success, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	svc := todo.NewTaskService(todo.NewTaskRepository(r.Db))
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
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id int) (*model.Task, error) {
	svc := todo.NewTaskService(todo.NewTaskRepository(r.Db))
	task, err := svc.GetTask(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

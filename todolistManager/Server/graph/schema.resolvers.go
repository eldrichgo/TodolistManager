package graph

import (
	"context"
	"fmt"
	"server/graph/model"
	"server/models"
	"strconv"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, title string, status string) (*model.Task, error) {
	task := &models.Task{
		Title:  title,
		Status: status,
	}
	err := r.TaskService.CreateTask(task)
	if err != nil {
		return nil, err
	}
	return &model.Task{
		ID:     strconv.Itoa(task.ID),
		Title:  task.Title,
		Status: task.Status,
	}, nil
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, id string, title *string, status *string) (*model.Task, error) {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID format")
	}
	// Get all tasks and find the task to be updated
	tasks, err := r.TaskService.GetAllTasks()
	if err != nil {
		return nil, err
	}
	var task *models.Task
	for _, t := range tasks {
		if t.ID == taskID {
			task = &t
			break
		}
	}
	if task == nil {
		return nil, fmt.Errorf("task not found")
	}
	if title != nil {
		task.Title = *title
	}
	if status != nil {
		task.Status = *status
	}
	err = r.TaskService.UpdateTaskStatus(task.ID, task.Status)
	if err != nil {
		return nil, err
	}
	return &model.Task{
		ID:     strconv.Itoa(task.ID),
		Title:  task.Title,
		Status: task.Status,
	}, nil
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (*bool, error) {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID format")
	}
	err = r.TaskService.DeleteTask(taskID)
	if err != nil {
		return nil, err
	}
	success := true
	return &success, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	tasks, err := r.TaskService.GetAllTasks()
	if err != nil {
		return nil, err
	}
	var result []*model.Task
	for _, task := range tasks {
		result = append(result, &model.Task{
			ID:     strconv.Itoa(task.ID),
			Title:  task.Title,
			Status: task.Status,
		})
	}
	return result, nil
}

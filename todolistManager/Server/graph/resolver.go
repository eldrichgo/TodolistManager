package graph

import (
	"server/dal/service" // Import your service package
)

type Resolver struct {
	TaskService *service.TaskService
}

func NewResolver(taskService *service.TaskService) *Resolver {
	return &Resolver{
		TaskService: taskService,
	}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

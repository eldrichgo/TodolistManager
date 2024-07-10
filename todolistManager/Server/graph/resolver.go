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

//Implemented by gqlgen automatically when you run the command "go run github.com/99designs/gqlgen" in the Server/graph directory in schema.resolvers.go???
// func (r *Resolver) Query() QueryResolver {
// 	return &queryResolver{r}
// }

// func (r *Resolver) Mutation() MutationResolver {
// 	return &mutationResolver{r}
// }

// type queryResolver struct{ *Resolver }
// type mutationResolver struct{ *Resolver }

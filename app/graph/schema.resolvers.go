package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"
	"gin/graph/model"
	"gin/service"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	newTodo := Todo{
		id:   fmt.Sprintf("T%d", len(todoList)+1),
		text: input.Text,
		done: false,
	}

	todoList = append(todoList, newTodo)
	return &model.Todo{
		ID:   newTodo.id,
		Text: newTodo.text,
		Done: newTodo.done,
	}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string, email *string) (*model.User, error) {
	user, err := service.CreateUser(name, email)
	return user, err
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := service.GetUsers()
	return users, err
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// convert todoList to model.Todo
	todo := []*model.Todo{}
	for _, v := range todoList {
		todo = append(todo, &model.Todo{
			ID:   v.id,
			Text: v.text,
			Done: v.done,
		})
	}

	return todo, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
type Todo struct {
	id   string
	text string
	done bool
}

var todoList = []Todo{}
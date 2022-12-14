package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mujak27/GO/auth"
	"github.com/mujak27/GO/graph/generated"
	"github.com/mujak27/GO/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	fmt.Println("---------------")
	fmt.Println(auth.JwtGetValue(ctx))
	fmt.Println(auth.JwtGetValue(ctx).Userid)
	newTodo := &model.Todo{
		ID:     uuid.NewString(),
		Text:   input.Text,
		Done:   false,
		UserID: auth.JwtGetValue(ctx).Userid,
	}
	r.todos = append(r.todos, newTodo)

	err := r.DB.Create(newTodo).Error
	return newTodo, err
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.InputTodo) (*model.Todo, error) {
	var todo *model.Todo
	if err := r.DB.First(todo, "id = ?", input.ID).Error; err != nil {
		return nil, err
	}
	todo.Text = input.Text
	return todo, r.DB.Save(todo).Error
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	return todos, r.DB.Find(&todos).Error
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) NewMutation(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

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

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, input model.InputLogin) (interface{}, error) {
	raw, _ := ctx.Value(string("tokenValue")).(*auth.JwtCustomClaim)
	fmt.Println(raw.Userid)
	return map[string]interface{}{
		"token": raw,
	}, nil
}

// Register is the resolver for the register field.
func (r *queryResolver) Register(ctx context.Context, input *model.InputRegister) (interface{}, error) {
	newUser := &model.User{
		ID:           uuid.NewString(),
		UserName:     input.UserName,
		UserEmail:    input.UserEmail,
		UserPassword: input.UserPassword,
	}
	r.users = append(r.users, newUser)

	token, err := auth.JwtGenerate(ctx, newUser.ID)
	if err != nil {
		return nil, err
	}
	r.DB.Create(newUser)

	return map[string]interface{}{
		"token": token,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

package graph

import (
	"github.com/mujak27/GO/graph/model"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//go:generate go run github.com/99designs/gqlgen generate

// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	users []*model.User
	DB    *gorm.DB
}

package graph
//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"la-cipollina-budgeter-api/graph/model"
)

type Resolver struct{
	todos []*model.Todo

}

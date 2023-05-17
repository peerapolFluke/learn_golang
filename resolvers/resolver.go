package resolvers

import (
	"github.com/99designs/gqlgen/graphql"
	"seamoor"
	"seamoor/cors/directives"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
}

func NewSchema() graphql.ExecutableSchema {
	c := seamoor.Config{Resolvers: &Resolver{}}
	c.Directives.Binding = directives.Binding
	c.Directives.AuthGuard = directives.AuthGuard

	return seamoor.NewExecutableSchema(c)
}

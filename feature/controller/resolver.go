package controller

// This file will not be regenerated automatically.
import (
	"github.com/pumy2517/ginent/ent"

	"github.com/99designs/gqlgen/graphql"
	"github.com/pumy2517/ginent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return ginent.NewExecutableSchema(ginent.Config{
		Resolvers: &Resolver{
			client,
		},
	})
}

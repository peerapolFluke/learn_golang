package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"seamoor/cors/auth"
	"seamoor/cors/utils"
)

func AuthGuard(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, utils.CustomGqlError(ctx, "Access Denied", 403)
	}

	return next(ctx)
}

package utils

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func CustomGqlError(ctx context.Context, errorMessage string, statusCode int) error {
	return &gqlerror.Error{
		Message: errorMessage,
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"statusCode": statusCode,
		},
	}
}

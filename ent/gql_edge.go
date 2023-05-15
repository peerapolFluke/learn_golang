// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (t *Test01) Parent(ctx context.Context) (*Test01, error) {
	result, err := t.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Test01) Children(ctx context.Context) (result []*Test01, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.ChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryChildren().All(ctx)
	}
	return result, err
}

func (t *Todo) Parent(ctx context.Context) (*Todo, error) {
	result, err := t.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Todo) Children(ctx context.Context) (result []*Todo, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.ChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryChildren().All(ctx)
	}
	return result, err
}

package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Test01 struct {
	ent.Schema
}

// Fields of the Todo.
func (Test01) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").
			NotEmpty().Annotations(
			entgql.OrderField("TEXT"),
		),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Time("update_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("UPDATE_AT"),
			),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS").
			Annotations(
				entgql.OrderField("STATUS"),
			),
		field.Int("priority").
			Default(0).
			Annotations(
				entgql.OrderField("PRIORITY"),
			),
	}
}

// Edges of the Todo.
func (Test01) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Test01.Type).
			From("parent").
			Annotations(
				entgql.OrderField("PARENT_PRIORITY"),
			).
			Unique(),
	}
}

func (Test01) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
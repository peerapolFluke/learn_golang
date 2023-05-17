package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("surname").NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("position").Optional(),
		field.String("phone_number").Optional(),
		field.String("address").Optional(),
		field.String("profile_img_url").Optional(),
		field.String("rubber_stamp_url").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("last_login_at").Optional(),
		field.Enum("role").Values("super_admin", "admin", "member"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

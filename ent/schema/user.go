package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Default("unknown"),
		field.Time("created_at").Default(time.Now),
		field.String("password").Default("isnotasecret"),
		field.String("private_key").Default("unknown"),
		field.String("public_key").Default("unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
		edge.From("chats", Chat.Type).Ref("members"),
	}
}

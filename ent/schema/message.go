package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("body").NotEmpty(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "from" of type `User`
		// and reference it to the "messages" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("from", User.Type).Ref("messages").
			// setting the edge to unique, ensure
			// that a message can have only one from.
			Unique(),
		edge.From("where", Chat.Type).Ref("messages"),
	}
}

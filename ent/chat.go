// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/amaru0601/fluent/ent/chat"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Chat is the model entity for the Chat schema.
type Chat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type chat.Type `json:"type,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChatQuery when eager-loading is set.
	Edges ChatEdges `json:"edges"`
}

// ChatEdges holds the relations/edges for other nodes in the graph.
type ChatEdges struct {
	// Members holds the value of the members edge.
	Members []*User `json:"members,omitempty"`
	// Has holds the value of the has edge.
	Has []*Message `json:"has,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e ChatEdges) MembersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// HasOrErr returns the Has value or an error if the edge
// was not loaded in eager-loading.
func (e ChatEdges) HasOrErr() ([]*Message, error) {
	if e.loadedTypes[1] {
		return e.Has, nil
	}
	return nil, &NotLoadedError{edge: "has"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Chat) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case chat.FieldDeleted:
			values[i] = new(sql.NullBool)
		case chat.FieldID:
			values[i] = new(sql.NullInt64)
		case chat.FieldName, chat.FieldType:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Chat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Chat fields.
func (c *Chat) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case chat.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case chat.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = chat.Type(value.String)
			}
		case chat.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				c.Deleted = value.Bool
			}
		}
	}
	return nil
}

// QueryMembers queries the "members" edge of the Chat entity.
func (c *Chat) QueryMembers() *UserQuery {
	return (&ChatClient{config: c.config}).QueryMembers(c)
}

// QueryHas queries the "has" edge of the Chat entity.
func (c *Chat) QueryHas() *MessageQuery {
	return (&ChatClient{config: c.config}).QueryHas(c)
}

// Update returns a builder for updating this Chat.
// Note that you need to call Chat.Unwrap() before calling this method if this Chat
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Chat) Update() *ChatUpdateOne {
	return (&ChatClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Chat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Chat) Unwrap() *Chat {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Chat is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Chat) String() string {
	var builder strings.Builder
	builder.WriteString("Chat(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteString(", deleted=")
	builder.WriteString(fmt.Sprintf("%v", c.Deleted))
	builder.WriteByte(')')
	return builder.String()
}

// Chats is a parsable slice of Chat.
type Chats []*Chat

func (c Chats) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}

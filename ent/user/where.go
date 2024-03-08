// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/jaox1/chat-server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PrivateKey applies equality check predicate on the "private_key" field. It's identical to PrivateKeyEQ.
func PrivateKey(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPrivateKey, v))
}

// PublicKey applies equality check predicate on the "public_key" field. It's identical to PublicKeyEQ.
func PublicKey(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPublicKey, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// PrivateKeyEQ applies the EQ predicate on the "private_key" field.
func PrivateKeyEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPrivateKey, v))
}

// PrivateKeyNEQ applies the NEQ predicate on the "private_key" field.
func PrivateKeyNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPrivateKey, v))
}

// PrivateKeyIn applies the In predicate on the "private_key" field.
func PrivateKeyIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPrivateKey, vs...))
}

// PrivateKeyNotIn applies the NotIn predicate on the "private_key" field.
func PrivateKeyNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPrivateKey, vs...))
}

// PrivateKeyGT applies the GT predicate on the "private_key" field.
func PrivateKeyGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPrivateKey, v))
}

// PrivateKeyGTE applies the GTE predicate on the "private_key" field.
func PrivateKeyGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPrivateKey, v))
}

// PrivateKeyLT applies the LT predicate on the "private_key" field.
func PrivateKeyLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPrivateKey, v))
}

// PrivateKeyLTE applies the LTE predicate on the "private_key" field.
func PrivateKeyLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPrivateKey, v))
}

// PrivateKeyContains applies the Contains predicate on the "private_key" field.
func PrivateKeyContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPrivateKey, v))
}

// PrivateKeyHasPrefix applies the HasPrefix predicate on the "private_key" field.
func PrivateKeyHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPrivateKey, v))
}

// PrivateKeyHasSuffix applies the HasSuffix predicate on the "private_key" field.
func PrivateKeyHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPrivateKey, v))
}

// PrivateKeyEqualFold applies the EqualFold predicate on the "private_key" field.
func PrivateKeyEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPrivateKey, v))
}

// PrivateKeyContainsFold applies the ContainsFold predicate on the "private_key" field.
func PrivateKeyContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPrivateKey, v))
}

// PublicKeyEQ applies the EQ predicate on the "public_key" field.
func PublicKeyEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPublicKey, v))
}

// PublicKeyNEQ applies the NEQ predicate on the "public_key" field.
func PublicKeyNEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPublicKey, v))
}

// PublicKeyIn applies the In predicate on the "public_key" field.
func PublicKeyIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldIn(FieldPublicKey, vs...))
}

// PublicKeyNotIn applies the NotIn predicate on the "public_key" field.
func PublicKeyNotIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPublicKey, vs...))
}

// PublicKeyGT applies the GT predicate on the "public_key" field.
func PublicKeyGT(v []byte) predicate.User {
	return predicate.User(sql.FieldGT(FieldPublicKey, v))
}

// PublicKeyGTE applies the GTE predicate on the "public_key" field.
func PublicKeyGTE(v []byte) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPublicKey, v))
}

// PublicKeyLT applies the LT predicate on the "public_key" field.
func PublicKeyLT(v []byte) predicate.User {
	return predicate.User(sql.FieldLT(FieldPublicKey, v))
}

// PublicKeyLTE applies the LTE predicate on the "public_key" field.
func PublicKeyLTE(v []byte) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPublicKey, v))
}

// HasMessages applies the HasEdge predicate on the "messages" edge.
func HasMessages() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessagesWith applies the HasEdge predicate on the "messages" edge with a given conditions (other predicates).
func HasMessagesWith(preds ...predicate.Message) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newMessagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChats applies the HasEdge predicate on the "chats" edge.
func HasChats() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ChatsTable, ChatsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChatsWith applies the HasEdge predicate on the "chats" edge with a given conditions (other predicates).
func HasChatsWith(preds ...predicate.Chat) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newChatsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}

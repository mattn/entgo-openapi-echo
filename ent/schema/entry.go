package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Entry holds the schema definition for the Entry entity.
type Entry struct {
	ent.Schema
}

// Fields of the Entry.
func (Entry) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
		field.Time("created_at").
			Default(time.Now),
	}
}

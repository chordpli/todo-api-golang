package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("title").NotEmpty(),
		field.Text("description").Optional(),
		field.Enum("status").
			Values("PENDING", "COMPLETED", "PROGRESS").
			Default("PENDING"),
	}
}

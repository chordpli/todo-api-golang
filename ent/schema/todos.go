package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Text("description").Optional(),
		field.Enum("status").
			Values("PENDING", "COMPLETED", "PROGRESS").
			Default("PENDING"),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

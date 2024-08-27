package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
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
		CustomTimeMixin{}, // 커스텀 믹스인 사용
	}
}

type CustomTimeMixin struct {
	mixin.Schema
}

// Fields implements the ent.Mixin interface.
func (CustomTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("The time when the entity was created."),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("The time when the entity was last updated."),
	}
}

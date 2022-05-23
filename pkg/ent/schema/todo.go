package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Tоdo holds the schema definition for the Tоdo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive(),
		field.String("slug").NotEmpty(),
		field.String("name"),
		field.Enum("status").
			Values("Todo", "InProgress", "Done", "Review").
			Default("Todo"),

		field.Time("created_at").Optional().Nillable().Default(time.Now()),
		field.Time("updated_at").Optional().Nillable(),
		field.Time("finish_at").Optional().Nillable(),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("status", Status.Type),
	}
}

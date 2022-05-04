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
		field.Int("id").
			Positive(),
		field.String("slug").NotEmpty(),
		field.String("name").NotEmpty(),
		field.Enum("status").
			Values("Todo", "InProgress", "Done", "Review").
			Default("Todo"),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("status", Status.Type),
	}
}

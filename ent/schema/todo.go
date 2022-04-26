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
		field.String("Name").
			Default("Null"),
		field.String("Slug").
			Default("Null"),
		field.Enum("Status").
			Values("Draft", "Start", "InProcess", "Review", "Done").
			Default("Start"),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}

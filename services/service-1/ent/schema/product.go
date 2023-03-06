package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Float("price").Positive(),
		field.String("name"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Customer schema.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		 field.String("NAME"),
		 field.Int("AGE"),
		 field.String("EMAIL"),
		 field.String("USERNAME"),
		 field.String("PASSWORD"),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
    return []ent.Edge{
		edge.To("checkout1", Checkout.Type),
    }
}

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Room schema.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		 field.String("ROOMNUMBER"),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
    return []ent.Edge{
         edge.To("checkout3", Checkout.Type),
    }
}

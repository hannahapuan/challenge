package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// IPAddress holds the schema definition for the IPAddress entity.
type IPAddress struct {
	ent.Schema
}

// TODO: add validation for ipv4 to ipaddress field value
// TODO: add validation for response codes
// Fields of the IPAddress.
func (IPAddress) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").
			Immutable().
			Unique(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now),
		field.String("response_code"),
		field.String("ip_address").
			Unique().
			Immutable(),
	}
}

// Edges of the IPAddress.
func (IPAddress) Edges() []ent.Edge {
	return nil
}

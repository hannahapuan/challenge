package schema

import (
	"time"

	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// IPAddress holds the schema definition for the IPAddress entity.
type IPAddress struct {
	ent.Schema
}

// TODO: add validation for ipv4 to ipaddress field value
// Fields of the IPAddress.
func (IPAddress) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.New()).
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now),
		field.String("response_code"),
		field.String("ip_address"),
	}
}

// Edges of the IPAddress.
func (IPAddress) Edges() []ent.Edge {
	return nil
}

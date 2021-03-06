// Code generated by entc, DO NOT EDIT.

package ent

import (
	"challenge/ent/ipaddress"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// IPAddress is the model entity for the IPAddress schema.
type IPAddress struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ResponseCode holds the value of the "response_code" field.
	ResponseCode string `json:"response_code,omitempty"`
	// IPAddress holds the value of the "ip_address" field.
	IPAddress string `json:"ip_address,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IPAddress) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ipaddress.FieldID:
			values[i] = new(sql.NullInt64)
		case ipaddress.FieldUUID, ipaddress.FieldResponseCode, ipaddress.FieldIPAddress:
			values[i] = new(sql.NullString)
		case ipaddress.FieldCreatedAt, ipaddress.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type IPAddress", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IPAddress fields.
func (ia *IPAddress) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ipaddress.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ia.ID = int(value.Int64)
		case ipaddress.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				ia.UUID = value.String
			}
		case ipaddress.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ia.CreatedAt = value.Time
			}
		case ipaddress.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ia.UpdatedAt = value.Time
			}
		case ipaddress.FieldResponseCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field response_code", values[i])
			} else if value.Valid {
				ia.ResponseCode = value.String
			}
		case ipaddress.FieldIPAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_address", values[i])
			} else if value.Valid {
				ia.IPAddress = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this IPAddress.
// Note that you need to call IPAddress.Unwrap() before calling this method if this IPAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (ia *IPAddress) Update() *IPAddressUpdateOne {
	return (&IPAddressClient{config: ia.config}).UpdateOne(ia)
}

// Unwrap unwraps the IPAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ia *IPAddress) Unwrap() *IPAddress {
	tx, ok := ia.config.driver.(*txDriver)
	if !ok {
		panic("ent: IPAddress is not a transactional entity")
	}
	ia.config.driver = tx.drv
	return ia
}

// String implements the fmt.Stringer.
func (ia *IPAddress) String() string {
	var builder strings.Builder
	builder.WriteString("IPAddress(")
	builder.WriteString(fmt.Sprintf("id=%v", ia.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(ia.UUID)
	builder.WriteString(", created_at=")
	builder.WriteString(ia.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ia.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", response_code=")
	builder.WriteString(ia.ResponseCode)
	builder.WriteString(", ip_address=")
	builder.WriteString(ia.IPAddress)
	builder.WriteByte(')')
	return builder.String()
}

// IPAddresses is a parsable slice of IPAddress.
type IPAddresses []*IPAddress

func (ia IPAddresses) config(cfg config) {
	for _i := range ia {
		ia[_i].config = cfg
	}
}

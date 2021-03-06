// Code generated by entc, DO NOT EDIT.

package ent

import (
	"challenge/ent/ipaddress"
	"challenge/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	ipaddressFields := schema.IPAddress{}.Fields()
	_ = ipaddressFields
	// ipaddressDescCreatedAt is the schema descriptor for created_at field.
	ipaddressDescCreatedAt := ipaddressFields[1].Descriptor()
	// ipaddress.DefaultCreatedAt holds the default value on creation for the created_at field.
	ipaddress.DefaultCreatedAt = ipaddressDescCreatedAt.Default.(func() time.Time)
	// ipaddressDescUpdatedAt is the schema descriptor for updated_at field.
	ipaddressDescUpdatedAt := ipaddressFields[2].Descriptor()
	// ipaddress.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ipaddress.DefaultUpdatedAt = ipaddressDescUpdatedAt.Default.(func() time.Time)
}

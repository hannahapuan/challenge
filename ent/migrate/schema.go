// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// IPAddressesColumns holds the columns for the "ip_addresses" table.
	IPAddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "response_code", Type: field.TypeString},
		{Name: "ip_address", Type: field.TypeString},
	}
	// IPAddressesTable holds the schema information for the "ip_addresses" table.
	IPAddressesTable = &schema.Table{
		Name:       "ip_addresses",
		Columns:    IPAddressesColumns,
		PrimaryKey: []*schema.Column{IPAddressesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		IPAddressesTable,
	}
)

func init() {
}

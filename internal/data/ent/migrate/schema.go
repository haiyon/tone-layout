// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SampGreeterColumns holds the columns for the "samp_greeter" table.
	SampGreeterColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SampGreeterTable holds the schema information for the "samp_greeter" table.
	SampGreeterTable = &schema.Table{
		Name:       "samp_greeter",
		Columns:    SampGreeterColumns,
		PrimaryKey: []*schema.Column{SampGreeterColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SampGreeterTable,
	}
)

func init() {
	SampGreeterTable.Annotation = &entsql.Annotation{
		Table: "samp_greeter",
	}
}

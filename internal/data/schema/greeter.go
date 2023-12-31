package schema

import (
	"sample/helper/utils"
	"strings"
	"time"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Greeter holds the schema definition for the Greeter entity.
type Greeter struct {
	ent.Schema
}

// Annotations of the Greeter.
func (Greeter) Annotations() []schema.Annotation {
	table := strings.Join([]string{"samp", "greeter"}, "_")
	return []schema.Annotation{
		entsql.Annotation{Table: table},
	}
}

// Fields of the Greeter.
func (Greeter) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(utils.PrimaryKey()), // primary key
		field.String("name").Unique(),
		field.Time("created_at").Default(time.Now),                         // created at
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now), // updated at
		field.Time("deleted_at").Nillable().Optional(),                     // deleted at
	}
}

// Edges of the Greeter.
func (Greeter) Edges() []ent.Edge {
	return []ent.Edge{}
}

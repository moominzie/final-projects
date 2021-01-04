package schema

import "github.com/facebookincubator/ent"

// Table holds the schema definition for the Table entity.
type Table struct {
	ent.Schema
}

// Fields of the Table.
func (Table) Fields() []ent.Field {
	return nil
}

// Edges of the Table.
func (Table) Edges() []ent.Edge {
	return nil
}

package schema

import "entgo.io/ent"

// Cluster holds the schema definition for the Cluster entity.
type Cluster struct {
	ent.Schema
}

// Fields of the Cluster.
func (Cluster) Fields() []ent.Field {
	return nil
}

// Edges of the Cluster.
func (Cluster) Edges() []ent.Edge {
	return nil
}

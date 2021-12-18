package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GitHubUser holds the schema definition for the GitHubUser entity.
type GitHubUser struct {
	ent.Schema
}

// Fields of the GitHubUser.
func (GitHubUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("login").Unique(),
		field.String("email").Unique(),
		field.String("name"),
	}
}

// Edges of the GitHubUser.
func (GitHubUser) Edges() []ent.Edge {
	return nil
}

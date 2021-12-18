package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GitHubUser holds the schema definition for the GitHubUser entity.
type GitHubUser struct {
	ent.Schema
}

// Fields of the GitHubUser.
func (GitHubUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique(),

		field.Int("github_id").Positive().Unique(),
		field.String("login").Unique(),
		field.String("email").Unique(),
		field.String("name"),
	}
}

// Edges of the GitHubUser.
func (GitHubUser) Edges() []ent.Edge {
	return []ent.Edge{
		// out edges
		// issues created by this user
		edge.To("created_issues", GitHubIssue.Type),

		// issues closed by this user
		edge.To("closed_issues", GitHubIssue.Type),

		// in (inverse) edges

		// person that owns this GitHub user account
		edge.From("person", Person.Type).
			Ref("github_account").
			Unique(),

		// issues this GitHub user is assigned to
		edge.From("assigned_issues", GitHubIssue.Type).
			Ref("assignee"),
	}
}

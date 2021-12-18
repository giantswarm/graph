package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GitHubIssue holds the schema definition for the GitHubIssue entity.
type GitHubIssue struct {
	ent.Schema
}

// Fields of the GitHubIssue.
func (GitHubIssue) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.Int("number").Positive(),
		field.String("title").NotEmpty(),
		field.String("body"),
		field.String("html_url").Unique(),
		field.Enum("state").Values("open", "closed"),
		field.Bool("locked").Default(false),
		field.String("active_lock_reason"),
		field.Int("comments_count").NonNegative(),
		field.String("created_at").NotEmpty().Validate(func(createdAt string) error {
			// validate datetime
			return nil
		}),
		field.String("updated_at").Validate(func(createdAt string) error {
			// validate datetime
			return nil
		}),
		field.String("closed_at").Validate(func(createdAt string) error {
			// validate datetime
			return nil
		}),
		field.String("author_association"),
	}
}

// Edges of the GitHubIssue.
func (GitHubIssue) Edges() []ent.Edge {
	return nil
}

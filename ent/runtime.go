// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/giantswarm/graph/ent/githubissue"
	"github.com/giantswarm/graph/ent/githubuser"
	"github.com/giantswarm/graph/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	githubissueFields := schema.GitHubIssue{}.Fields()
	_ = githubissueFields
	// githubissueDescGithubID is the schema descriptor for github_id field.
	githubissueDescGithubID := githubissueFields[1].Descriptor()
	// githubissue.GithubIDValidator is a validator for the "github_id" field. It is called by the builders before save.
	githubissue.GithubIDValidator = githubissueDescGithubID.Validators[0].(func(int) error)
	// githubissueDescNumber is the schema descriptor for number field.
	githubissueDescNumber := githubissueFields[2].Descriptor()
	// githubissue.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	githubissue.NumberValidator = githubissueDescNumber.Validators[0].(func(int) error)
	// githubissueDescTitle is the schema descriptor for title field.
	githubissueDescTitle := githubissueFields[3].Descriptor()
	// githubissue.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	githubissue.TitleValidator = githubissueDescTitle.Validators[0].(func(string) error)
	// githubissueDescLocked is the schema descriptor for locked field.
	githubissueDescLocked := githubissueFields[7].Descriptor()
	// githubissue.DefaultLocked holds the default value on creation for the locked field.
	githubissue.DefaultLocked = githubissueDescLocked.Default.(bool)
	// githubissueDescCommentsCount is the schema descriptor for comments_count field.
	githubissueDescCommentsCount := githubissueFields[9].Descriptor()
	// githubissue.CommentsCountValidator is a validator for the "comments_count" field. It is called by the builders before save.
	githubissue.CommentsCountValidator = githubissueDescCommentsCount.Validators[0].(func(int) error)
	// githubissueDescCreatedAt is the schema descriptor for created_at field.
	githubissueDescCreatedAt := githubissueFields[10].Descriptor()
	// githubissue.CreatedAtValidator is a validator for the "created_at" field. It is called by the builders before save.
	githubissue.CreatedAtValidator = func() func(string) error {
		validators := githubissueDescCreatedAt.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(created_at string) error {
			for _, fn := range fns {
				if err := fn(created_at); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// githubissueDescUpdatedAt is the schema descriptor for updated_at field.
	githubissueDescUpdatedAt := githubissueFields[11].Descriptor()
	// githubissue.UpdatedAtValidator is a validator for the "updated_at" field. It is called by the builders before save.
	githubissue.UpdatedAtValidator = githubissueDescUpdatedAt.Validators[0].(func(string) error)
	// githubissueDescClosedAt is the schema descriptor for closed_at field.
	githubissueDescClosedAt := githubissueFields[12].Descriptor()
	// githubissue.ClosedAtValidator is a validator for the "closed_at" field. It is called by the builders before save.
	githubissue.ClosedAtValidator = githubissueDescClosedAt.Validators[0].(func(string) error)
	// githubissueDescID is the schema descriptor for id field.
	githubissueDescID := githubissueFields[0].Descriptor()
	// githubissue.IDValidator is a validator for the "id" field. It is called by the builders before save.
	githubissue.IDValidator = githubissueDescID.Validators[0].(func(string) error)
	githubuserFields := schema.GitHubUser{}.Fields()
	_ = githubuserFields
	// githubuserDescGithubID is the schema descriptor for github_id field.
	githubuserDescGithubID := githubuserFields[1].Descriptor()
	// githubuser.GithubIDValidator is a validator for the "github_id" field. It is called by the builders before save.
	githubuser.GithubIDValidator = githubuserDescGithubID.Validators[0].(func(int) error)
	// githubuserDescID is the schema descriptor for id field.
	githubuserDescID := githubuserFields[0].Descriptor()
	// githubuser.IDValidator is a validator for the "id" field. It is called by the builders before save.
	githubuser.IDValidator = githubuserDescID.Validators[0].(func(string) error)
}

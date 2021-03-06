// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"github.com/giantswarm/graph/ent/githubissue"
	"github.com/giantswarm/graph/ent/githubuser"
)

// GitHubIssueCreate is the builder for creating a GitHubIssue entity.
type GitHubIssueCreate struct {
	config
	mutation *GitHubIssueMutation
	hooks    []Hook
}

// SetGithubID sets the "github_id" field.
func (ghic *GitHubIssueCreate) SetGithubID(i int) *GitHubIssueCreate {
	ghic.mutation.SetGithubID(i)
	return ghic
}

// SetNumber sets the "number" field.
func (ghic *GitHubIssueCreate) SetNumber(i int) *GitHubIssueCreate {
	ghic.mutation.SetNumber(i)
	return ghic
}

// SetTitle sets the "title" field.
func (ghic *GitHubIssueCreate) SetTitle(s string) *GitHubIssueCreate {
	ghic.mutation.SetTitle(s)
	return ghic
}

// SetBody sets the "body" field.
func (ghic *GitHubIssueCreate) SetBody(s string) *GitHubIssueCreate {
	ghic.mutation.SetBody(s)
	return ghic
}

// SetHTMLURL sets the "html_url" field.
func (ghic *GitHubIssueCreate) SetHTMLURL(s string) *GitHubIssueCreate {
	ghic.mutation.SetHTMLURL(s)
	return ghic
}

// SetState sets the "state" field.
func (ghic *GitHubIssueCreate) SetState(gi githubissue.State) *GitHubIssueCreate {
	ghic.mutation.SetState(gi)
	return ghic
}

// SetLocked sets the "locked" field.
func (ghic *GitHubIssueCreate) SetLocked(b bool) *GitHubIssueCreate {
	ghic.mutation.SetLocked(b)
	return ghic
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (ghic *GitHubIssueCreate) SetNillableLocked(b *bool) *GitHubIssueCreate {
	if b != nil {
		ghic.SetLocked(*b)
	}
	return ghic
}

// SetActiveLockReason sets the "active_lock_reason" field.
func (ghic *GitHubIssueCreate) SetActiveLockReason(s string) *GitHubIssueCreate {
	ghic.mutation.SetActiveLockReason(s)
	return ghic
}

// SetCommentsCount sets the "comments_count" field.
func (ghic *GitHubIssueCreate) SetCommentsCount(i int) *GitHubIssueCreate {
	ghic.mutation.SetCommentsCount(i)
	return ghic
}

// SetCreatedAt sets the "created_at" field.
func (ghic *GitHubIssueCreate) SetCreatedAt(s string) *GitHubIssueCreate {
	ghic.mutation.SetCreatedAt(s)
	return ghic
}

// SetUpdatedAt sets the "updated_at" field.
func (ghic *GitHubIssueCreate) SetUpdatedAt(s string) *GitHubIssueCreate {
	ghic.mutation.SetUpdatedAt(s)
	return ghic
}

// SetClosedAt sets the "closed_at" field.
func (ghic *GitHubIssueCreate) SetClosedAt(s string) *GitHubIssueCreate {
	ghic.mutation.SetClosedAt(s)
	return ghic
}

// SetAuthorAssociation sets the "author_association" field.
func (ghic *GitHubIssueCreate) SetAuthorAssociation(s string) *GitHubIssueCreate {
	ghic.mutation.SetAuthorAssociation(s)
	return ghic
}

// SetID sets the "id" field.
func (ghic *GitHubIssueCreate) SetID(s string) *GitHubIssueCreate {
	ghic.mutation.SetID(s)
	return ghic
}

// AddAssigneeIDs adds the "assignees" edge to the GitHubUser entity by IDs.
func (ghic *GitHubIssueCreate) AddAssigneeIDs(ids ...string) *GitHubIssueCreate {
	ghic.mutation.AddAssigneeIDs(ids...)
	return ghic
}

// AddAssignees adds the "assignees" edges to the GitHubUser entity.
func (ghic *GitHubIssueCreate) AddAssignees(g ...*GitHubUser) *GitHubIssueCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghic.AddAssigneeIDs(ids...)
}

// SetAuthorID sets the "author" edge to the GitHubUser entity by ID.
func (ghic *GitHubIssueCreate) SetAuthorID(id string) *GitHubIssueCreate {
	ghic.mutation.SetAuthorID(id)
	return ghic
}

// SetAuthor sets the "author" edge to the GitHubUser entity.
func (ghic *GitHubIssueCreate) SetAuthor(g *GitHubUser) *GitHubIssueCreate {
	return ghic.SetAuthorID(g.ID)
}

// SetClosedByID sets the "closed_by" edge to the GitHubUser entity by ID.
func (ghic *GitHubIssueCreate) SetClosedByID(id string) *GitHubIssueCreate {
	ghic.mutation.SetClosedByID(id)
	return ghic
}

// SetNillableClosedByID sets the "closed_by" edge to the GitHubUser entity by ID if the given value is not nil.
func (ghic *GitHubIssueCreate) SetNillableClosedByID(id *string) *GitHubIssueCreate {
	if id != nil {
		ghic = ghic.SetClosedByID(*id)
	}
	return ghic
}

// SetClosedBy sets the "closed_by" edge to the GitHubUser entity.
func (ghic *GitHubIssueCreate) SetClosedBy(g *GitHubUser) *GitHubIssueCreate {
	return ghic.SetClosedByID(g.ID)
}

// Mutation returns the GitHubIssueMutation object of the builder.
func (ghic *GitHubIssueCreate) Mutation() *GitHubIssueMutation {
	return ghic.mutation
}

// Save creates the GitHubIssue in the database.
func (ghic *GitHubIssueCreate) Save(ctx context.Context) (*GitHubIssue, error) {
	var (
		err  error
		node *GitHubIssue
	)
	ghic.defaults()
	if len(ghic.hooks) == 0 {
		if err = ghic.check(); err != nil {
			return nil, err
		}
		node, err = ghic.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GitHubIssueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ghic.check(); err != nil {
				return nil, err
			}
			ghic.mutation = mutation
			if node, err = ghic.gremlinSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ghic.hooks) - 1; i >= 0; i-- {
			if ghic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ghic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ghic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ghic *GitHubIssueCreate) SaveX(ctx context.Context) *GitHubIssue {
	v, err := ghic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ghic *GitHubIssueCreate) Exec(ctx context.Context) error {
	_, err := ghic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ghic *GitHubIssueCreate) ExecX(ctx context.Context) {
	if err := ghic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ghic *GitHubIssueCreate) defaults() {
	if _, ok := ghic.mutation.Locked(); !ok {
		v := githubissue.DefaultLocked
		ghic.mutation.SetLocked(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ghic *GitHubIssueCreate) check() error {
	if _, ok := ghic.mutation.GithubID(); !ok {
		return &ValidationError{Name: "github_id", err: errors.New(`ent: missing required field "github_id"`)}
	}
	if v, ok := ghic.mutation.GithubID(); ok {
		if err := githubissue.GithubIDValidator(v); err != nil {
			return &ValidationError{Name: "github_id", err: fmt.Errorf(`ent: validator failed for field "github_id": %w`, err)}
		}
	}
	if _, ok := ghic.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "number"`)}
	}
	if v, ok := ghic.mutation.Number(); ok {
		if err := githubissue.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "number": %w`, err)}
		}
	}
	if _, ok := ghic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := ghic.mutation.Title(); ok {
		if err := githubissue.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if _, ok := ghic.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "body"`)}
	}
	if _, ok := ghic.mutation.HTMLURL(); !ok {
		return &ValidationError{Name: "html_url", err: errors.New(`ent: missing required field "html_url"`)}
	}
	if _, ok := ghic.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "state"`)}
	}
	if v, ok := ghic.mutation.State(); ok {
		if err := githubissue.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "state": %w`, err)}
		}
	}
	if _, ok := ghic.mutation.Locked(); !ok {
		return &ValidationError{Name: "locked", err: errors.New(`ent: missing required field "locked"`)}
	}
	if _, ok := ghic.mutation.ActiveLockReason(); !ok {
		return &ValidationError{Name: "active_lock_reason", err: errors.New(`ent: missing required field "active_lock_reason"`)}
	}
	if _, ok := ghic.mutation.CommentsCount(); !ok {
		return &ValidationError{Name: "comments_count", err: errors.New(`ent: missing required field "comments_count"`)}
	}
	if v, ok := ghic.mutation.CommentsCount(); ok {
		if err := githubissue.CommentsCountValidator(v); err != nil {
			return &ValidationError{Name: "comments_count", err: fmt.Errorf(`ent: validator failed for field "comments_count": %w`, err)}
		}
	}
	if _, ok := ghic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if v, ok := ghic.mutation.CreatedAt(); ok {
		if err := githubissue.CreatedAtValidator(v); err != nil {
			return &ValidationError{Name: "created_at", err: fmt.Errorf(`ent: validator failed for field "created_at": %w`, err)}
		}
	}
	if _, ok := ghic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if v, ok := ghic.mutation.UpdatedAt(); ok {
		if err := githubissue.UpdatedAtValidator(v); err != nil {
			return &ValidationError{Name: "updated_at", err: fmt.Errorf(`ent: validator failed for field "updated_at": %w`, err)}
		}
	}
	if _, ok := ghic.mutation.ClosedAt(); !ok {
		return &ValidationError{Name: "closed_at", err: errors.New(`ent: missing required field "closed_at"`)}
	}
	if v, ok := ghic.mutation.ClosedAt(); ok {
		if err := githubissue.ClosedAtValidator(v); err != nil {
			return &ValidationError{Name: "closed_at", err: fmt.Errorf(`ent: validator failed for field "closed_at": %w`, err)}
		}
	}
	if _, ok := ghic.mutation.AuthorAssociation(); !ok {
		return &ValidationError{Name: "author_association", err: errors.New(`ent: missing required field "author_association"`)}
	}
	if v, ok := ghic.mutation.ID(); ok {
		if err := githubissue.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "id": %w`, err)}
		}
	}
	if _, ok := ghic.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author", err: errors.New("ent: missing required edge \"author\"")}
	}
	return nil
}

func (ghic *GitHubIssueCreate) gremlinSave(ctx context.Context) (*GitHubIssue, error) {
	res := &gremlin.Response{}
	query, bindings := ghic.gremlin().Query()
	if err := ghic.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ghi := &GitHubIssue{config: ghic.config}
	if err := ghi.FromResponse(res); err != nil {
		return nil, err
	}
	return ghi, nil
}

func (ghic *GitHubIssueCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(githubissue.Label)
	if id, ok := ghic.mutation.ID(); ok {
		v.Property(dsl.ID, id)
	}
	if value, ok := ghic.mutation.GithubID(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(githubissue.Label, githubissue.FieldGithubID, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(githubissue.Label, githubissue.FieldGithubID, value)),
		})
		v.Property(dsl.Single, githubissue.FieldGithubID, value)
	}
	if value, ok := ghic.mutation.Number(); ok {
		v.Property(dsl.Single, githubissue.FieldNumber, value)
	}
	if value, ok := ghic.mutation.Title(); ok {
		v.Property(dsl.Single, githubissue.FieldTitle, value)
	}
	if value, ok := ghic.mutation.Body(); ok {
		v.Property(dsl.Single, githubissue.FieldBody, value)
	}
	if value, ok := ghic.mutation.HTMLURL(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(githubissue.Label, githubissue.FieldHTMLURL, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(githubissue.Label, githubissue.FieldHTMLURL, value)),
		})
		v.Property(dsl.Single, githubissue.FieldHTMLURL, value)
	}
	if value, ok := ghic.mutation.State(); ok {
		v.Property(dsl.Single, githubissue.FieldState, value)
	}
	if value, ok := ghic.mutation.Locked(); ok {
		v.Property(dsl.Single, githubissue.FieldLocked, value)
	}
	if value, ok := ghic.mutation.ActiveLockReason(); ok {
		v.Property(dsl.Single, githubissue.FieldActiveLockReason, value)
	}
	if value, ok := ghic.mutation.CommentsCount(); ok {
		v.Property(dsl.Single, githubissue.FieldCommentsCount, value)
	}
	if value, ok := ghic.mutation.CreatedAt(); ok {
		v.Property(dsl.Single, githubissue.FieldCreatedAt, value)
	}
	if value, ok := ghic.mutation.UpdatedAt(); ok {
		v.Property(dsl.Single, githubissue.FieldUpdatedAt, value)
	}
	if value, ok := ghic.mutation.ClosedAt(); ok {
		v.Property(dsl.Single, githubissue.FieldClosedAt, value)
	}
	if value, ok := ghic.mutation.AuthorAssociation(); ok {
		v.Property(dsl.Single, githubissue.FieldAuthorAssociation, value)
	}
	for _, id := range ghic.mutation.AssigneesIDs() {
		v.AddE(githubissue.AssigneesLabel).To(g.V(id)).OutV()
	}
	for _, id := range ghic.mutation.AuthorIDs() {
		v.AddE(githubuser.CreatedIssuesLabel).From(g.V(id)).InV()
	}
	for _, id := range ghic.mutation.ClosedByIDs() {
		v.AddE(githubuser.ClosedIssuesLabel).From(g.V(id)).InV()
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// GitHubIssueCreateBulk is the builder for creating many GitHubIssue entities in bulk.
type GitHubIssueCreateBulk struct {
	config
	builders []*GitHubIssueCreate
}

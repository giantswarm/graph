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
	"github.com/giantswarm/graph/ent/predicate"
)

// GitHubIssueUpdate is the builder for updating GitHubIssue entities.
type GitHubIssueUpdate struct {
	config
	hooks    []Hook
	mutation *GitHubIssueMutation
}

// Where appends a list predicates to the GitHubIssueUpdate builder.
func (ghiu *GitHubIssueUpdate) Where(ps ...predicate.GitHubIssue) *GitHubIssueUpdate {
	ghiu.mutation.Where(ps...)
	return ghiu
}

// SetGithubID sets the "github_id" field.
func (ghiu *GitHubIssueUpdate) SetGithubID(i int) *GitHubIssueUpdate {
	ghiu.mutation.ResetGithubID()
	ghiu.mutation.SetGithubID(i)
	return ghiu
}

// AddGithubID adds i to the "github_id" field.
func (ghiu *GitHubIssueUpdate) AddGithubID(i int) *GitHubIssueUpdate {
	ghiu.mutation.AddGithubID(i)
	return ghiu
}

// SetNumber sets the "number" field.
func (ghiu *GitHubIssueUpdate) SetNumber(i int) *GitHubIssueUpdate {
	ghiu.mutation.ResetNumber()
	ghiu.mutation.SetNumber(i)
	return ghiu
}

// AddNumber adds i to the "number" field.
func (ghiu *GitHubIssueUpdate) AddNumber(i int) *GitHubIssueUpdate {
	ghiu.mutation.AddNumber(i)
	return ghiu
}

// SetTitle sets the "title" field.
func (ghiu *GitHubIssueUpdate) SetTitle(s string) *GitHubIssueUpdate {
	ghiu.mutation.SetTitle(s)
	return ghiu
}

// SetBody sets the "body" field.
func (ghiu *GitHubIssueUpdate) SetBody(s string) *GitHubIssueUpdate {
	ghiu.mutation.SetBody(s)
	return ghiu
}

// SetHTMLURL sets the "html_url" field.
func (ghiu *GitHubIssueUpdate) SetHTMLURL(s string) *GitHubIssueUpdate {
	ghiu.mutation.SetHTMLURL(s)
	return ghiu
}

// SetState sets the "state" field.
func (ghiu *GitHubIssueUpdate) SetState(gi githubissue.State) *GitHubIssueUpdate {
	ghiu.mutation.SetState(gi)
	return ghiu
}

// SetLocked sets the "locked" field.
func (ghiu *GitHubIssueUpdate) SetLocked(b bool) *GitHubIssueUpdate {
	ghiu.mutation.SetLocked(b)
	return ghiu
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (ghiu *GitHubIssueUpdate) SetNillableLocked(b *bool) *GitHubIssueUpdate {
	if b != nil {
		ghiu.SetLocked(*b)
	}
	return ghiu
}

// SetActiveLockReason sets the "active_lock_reason" field.
func (ghiu *GitHubIssueUpdate) SetActiveLockReason(s string) *GitHubIssueUpdate {
	ghiu.mutation.SetActiveLockReason(s)
	return ghiu
}

// SetCommentsCount sets the "comments_count" field.
func (ghiu *GitHubIssueUpdate) SetCommentsCount(i int) *GitHubIssueUpdate {
	ghiu.mutation.ResetCommentsCount()
	ghiu.mutation.SetCommentsCount(i)
	return ghiu
}

// AddCommentsCount adds i to the "comments_count" field.
func (ghiu *GitHubIssueUpdate) AddCommentsCount(i int) *GitHubIssueUpdate {
	ghiu.mutation.AddCommentsCount(i)
	return ghiu
}

// SetCreatedAt sets the "created_at" field.
func (ghiu *GitHubIssueUpdate) SetCreatedAt(s string) *GitHubIssueUpdate {
	ghiu.mutation.SetCreatedAt(s)
	return ghiu
}

// SetUpdatedAt sets the "updated_at" field.
func (ghiu *GitHubIssueUpdate) SetUpdatedAt(s string) *GitHubIssueUpdate {
	ghiu.mutation.SetUpdatedAt(s)
	return ghiu
}

// SetClosedAt sets the "closed_at" field.
func (ghiu *GitHubIssueUpdate) SetClosedAt(s string) *GitHubIssueUpdate {
	ghiu.mutation.SetClosedAt(s)
	return ghiu
}

// SetAuthorAssociation sets the "author_association" field.
func (ghiu *GitHubIssueUpdate) SetAuthorAssociation(s string) *GitHubIssueUpdate {
	ghiu.mutation.SetAuthorAssociation(s)
	return ghiu
}

// AddAssigneeIDs adds the "assignees" edge to the GitHubUser entity by IDs.
func (ghiu *GitHubIssueUpdate) AddAssigneeIDs(ids ...int) *GitHubIssueUpdate {
	ghiu.mutation.AddAssigneeIDs(ids...)
	return ghiu
}

// AddAssignees adds the "assignees" edges to the GitHubUser entity.
func (ghiu *GitHubIssueUpdate) AddAssignees(g ...*GitHubUser) *GitHubIssueUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghiu.AddAssigneeIDs(ids...)
}

// SetAuthorID sets the "author" edge to the GitHubUser entity by ID.
func (ghiu *GitHubIssueUpdate) SetAuthorID(id int) *GitHubIssueUpdate {
	ghiu.mutation.SetAuthorID(id)
	return ghiu
}

// SetAuthor sets the "author" edge to the GitHubUser entity.
func (ghiu *GitHubIssueUpdate) SetAuthor(g *GitHubUser) *GitHubIssueUpdate {
	return ghiu.SetAuthorID(g.ID)
}

// SetClosedByID sets the "closed_by" edge to the GitHubUser entity by ID.
func (ghiu *GitHubIssueUpdate) SetClosedByID(id int) *GitHubIssueUpdate {
	ghiu.mutation.SetClosedByID(id)
	return ghiu
}

// SetNillableClosedByID sets the "closed_by" edge to the GitHubUser entity by ID if the given value is not nil.
func (ghiu *GitHubIssueUpdate) SetNillableClosedByID(id *int) *GitHubIssueUpdate {
	if id != nil {
		ghiu = ghiu.SetClosedByID(*id)
	}
	return ghiu
}

// SetClosedBy sets the "closed_by" edge to the GitHubUser entity.
func (ghiu *GitHubIssueUpdate) SetClosedBy(g *GitHubUser) *GitHubIssueUpdate {
	return ghiu.SetClosedByID(g.ID)
}

// Mutation returns the GitHubIssueMutation object of the builder.
func (ghiu *GitHubIssueUpdate) Mutation() *GitHubIssueMutation {
	return ghiu.mutation
}

// ClearAssignees clears all "assignees" edges to the GitHubUser entity.
func (ghiu *GitHubIssueUpdate) ClearAssignees() *GitHubIssueUpdate {
	ghiu.mutation.ClearAssignees()
	return ghiu
}

// RemoveAssigneeIDs removes the "assignees" edge to GitHubUser entities by IDs.
func (ghiu *GitHubIssueUpdate) RemoveAssigneeIDs(ids ...int) *GitHubIssueUpdate {
	ghiu.mutation.RemoveAssigneeIDs(ids...)
	return ghiu
}

// RemoveAssignees removes "assignees" edges to GitHubUser entities.
func (ghiu *GitHubIssueUpdate) RemoveAssignees(g ...*GitHubUser) *GitHubIssueUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghiu.RemoveAssigneeIDs(ids...)
}

// ClearAuthor clears the "author" edge to the GitHubUser entity.
func (ghiu *GitHubIssueUpdate) ClearAuthor() *GitHubIssueUpdate {
	ghiu.mutation.ClearAuthor()
	return ghiu
}

// ClearClosedBy clears the "closed_by" edge to the GitHubUser entity.
func (ghiu *GitHubIssueUpdate) ClearClosedBy() *GitHubIssueUpdate {
	ghiu.mutation.ClearClosedBy()
	return ghiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ghiu *GitHubIssueUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ghiu.hooks) == 0 {
		if err = ghiu.check(); err != nil {
			return 0, err
		}
		affected, err = ghiu.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GitHubIssueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ghiu.check(); err != nil {
				return 0, err
			}
			ghiu.mutation = mutation
			affected, err = ghiu.gremlinSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ghiu.hooks) - 1; i >= 0; i-- {
			if ghiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ghiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ghiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ghiu *GitHubIssueUpdate) SaveX(ctx context.Context) int {
	affected, err := ghiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ghiu *GitHubIssueUpdate) Exec(ctx context.Context) error {
	_, err := ghiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ghiu *GitHubIssueUpdate) ExecX(ctx context.Context) {
	if err := ghiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ghiu *GitHubIssueUpdate) check() error {
	if v, ok := ghiu.mutation.GithubID(); ok {
		if err := githubissue.GithubIDValidator(v); err != nil {
			return &ValidationError{Name: "github_id", err: fmt.Errorf("ent: validator failed for field \"github_id\": %w", err)}
		}
	}
	if v, ok := ghiu.mutation.Number(); ok {
		if err := githubissue.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf("ent: validator failed for field \"number\": %w", err)}
		}
	}
	if v, ok := ghiu.mutation.Title(); ok {
		if err := githubissue.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := ghiu.mutation.State(); ok {
		if err := githubissue.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	if v, ok := ghiu.mutation.CommentsCount(); ok {
		if err := githubissue.CommentsCountValidator(v); err != nil {
			return &ValidationError{Name: "comments_count", err: fmt.Errorf("ent: validator failed for field \"comments_count\": %w", err)}
		}
	}
	if v, ok := ghiu.mutation.CreatedAt(); ok {
		if err := githubissue.CreatedAtValidator(v); err != nil {
			return &ValidationError{Name: "created_at", err: fmt.Errorf("ent: validator failed for field \"created_at\": %w", err)}
		}
	}
	if v, ok := ghiu.mutation.UpdatedAt(); ok {
		if err := githubissue.UpdatedAtValidator(v); err != nil {
			return &ValidationError{Name: "updated_at", err: fmt.Errorf("ent: validator failed for field \"updated_at\": %w", err)}
		}
	}
	if v, ok := ghiu.mutation.ClosedAt(); ok {
		if err := githubissue.ClosedAtValidator(v); err != nil {
			return &ValidationError{Name: "closed_at", err: fmt.Errorf("ent: validator failed for field \"closed_at\": %w", err)}
		}
	}
	if _, ok := ghiu.mutation.AuthorID(); ghiu.mutation.AuthorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"author\"")
	}
	return nil
}

func (ghiu *GitHubIssueUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ghiu.gremlin().Query()
	if err := ghiu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (ghiu *GitHubIssueUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(githubissue.Label)
	for _, p := range ghiu.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := ghiu.mutation.GithubID(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(githubissue.Label, githubissue.FieldGithubID, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(githubissue.Label, githubissue.FieldGithubID, value)),
		})
		v.Property(dsl.Single, githubissue.FieldGithubID, value)
	}
	if value, ok := ghiu.mutation.AddedGithubID(); ok {
		addValue := rv.Clone().Union(__.Values(githubissue.FieldGithubID), __.Constant(value)).Sum().Next()
		constraints = append(constraints, &constraint{
			pred: g.V().Has(githubissue.Label, githubissue.FieldGithubID, addValue).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(githubissue.Label, githubissue.FieldGithubID, fmt.Sprintf("+= %v", value))),
		})
		v.Property(dsl.Single, githubissue.FieldGithubID, __.Union(__.Values(githubissue.FieldGithubID), __.Constant(value)).Sum())
	}
	if value, ok := ghiu.mutation.Number(); ok {
		v.Property(dsl.Single, githubissue.FieldNumber, value)
	}
	if value, ok := ghiu.mutation.AddedNumber(); ok {
		v.Property(dsl.Single, githubissue.FieldNumber, __.Union(__.Values(githubissue.FieldNumber), __.Constant(value)).Sum())
	}
	if value, ok := ghiu.mutation.Title(); ok {
		v.Property(dsl.Single, githubissue.FieldTitle, value)
	}
	if value, ok := ghiu.mutation.Body(); ok {
		v.Property(dsl.Single, githubissue.FieldBody, value)
	}
	if value, ok := ghiu.mutation.HTMLURL(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(githubissue.Label, githubissue.FieldHTMLURL, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(githubissue.Label, githubissue.FieldHTMLURL, value)),
		})
		v.Property(dsl.Single, githubissue.FieldHTMLURL, value)
	}
	if value, ok := ghiu.mutation.State(); ok {
		v.Property(dsl.Single, githubissue.FieldState, value)
	}
	if value, ok := ghiu.mutation.Locked(); ok {
		v.Property(dsl.Single, githubissue.FieldLocked, value)
	}
	if value, ok := ghiu.mutation.ActiveLockReason(); ok {
		v.Property(dsl.Single, githubissue.FieldActiveLockReason, value)
	}
	if value, ok := ghiu.mutation.CommentsCount(); ok {
		v.Property(dsl.Single, githubissue.FieldCommentsCount, value)
	}
	if value, ok := ghiu.mutation.AddedCommentsCount(); ok {
		v.Property(dsl.Single, githubissue.FieldCommentsCount, __.Union(__.Values(githubissue.FieldCommentsCount), __.Constant(value)).Sum())
	}
	if value, ok := ghiu.mutation.CreatedAt(); ok {
		v.Property(dsl.Single, githubissue.FieldCreatedAt, value)
	}
	if value, ok := ghiu.mutation.UpdatedAt(); ok {
		v.Property(dsl.Single, githubissue.FieldUpdatedAt, value)
	}
	if value, ok := ghiu.mutation.ClosedAt(); ok {
		v.Property(dsl.Single, githubissue.FieldClosedAt, value)
	}
	if value, ok := ghiu.mutation.AuthorAssociation(); ok {
		v.Property(dsl.Single, githubissue.FieldAuthorAssociation, value)
	}
	for _, id := range ghiu.mutation.RemovedAssigneesIDs() {
		tr := rv.Clone().OutE(githubissue.AssigneesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range ghiu.mutation.AssigneesIDs() {
		v.AddE(githubissue.AssigneesLabel).To(g.V(id)).OutV()
	}
	if ghiu.mutation.AuthorCleared() {
		tr := rv.Clone().InE(githubuser.CreatedIssuesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range ghiu.mutation.AuthorIDs() {
		v.AddE(githubuser.CreatedIssuesLabel).From(g.V(id)).InV()
	}
	if ghiu.mutation.ClosedByCleared() {
		tr := rv.Clone().InE(githubuser.ClosedIssuesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range ghiu.mutation.ClosedByIDs() {
		v.AddE(githubuser.ClosedIssuesLabel).From(g.V(id)).InV()
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// GitHubIssueUpdateOne is the builder for updating a single GitHubIssue entity.
type GitHubIssueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GitHubIssueMutation
}

// SetGithubID sets the "github_id" field.
func (ghiuo *GitHubIssueUpdateOne) SetGithubID(i int) *GitHubIssueUpdateOne {
	ghiuo.mutation.ResetGithubID()
	ghiuo.mutation.SetGithubID(i)
	return ghiuo
}

// AddGithubID adds i to the "github_id" field.
func (ghiuo *GitHubIssueUpdateOne) AddGithubID(i int) *GitHubIssueUpdateOne {
	ghiuo.mutation.AddGithubID(i)
	return ghiuo
}

// SetNumber sets the "number" field.
func (ghiuo *GitHubIssueUpdateOne) SetNumber(i int) *GitHubIssueUpdateOne {
	ghiuo.mutation.ResetNumber()
	ghiuo.mutation.SetNumber(i)
	return ghiuo
}

// AddNumber adds i to the "number" field.
func (ghiuo *GitHubIssueUpdateOne) AddNumber(i int) *GitHubIssueUpdateOne {
	ghiuo.mutation.AddNumber(i)
	return ghiuo
}

// SetTitle sets the "title" field.
func (ghiuo *GitHubIssueUpdateOne) SetTitle(s string) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetTitle(s)
	return ghiuo
}

// SetBody sets the "body" field.
func (ghiuo *GitHubIssueUpdateOne) SetBody(s string) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetBody(s)
	return ghiuo
}

// SetHTMLURL sets the "html_url" field.
func (ghiuo *GitHubIssueUpdateOne) SetHTMLURL(s string) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetHTMLURL(s)
	return ghiuo
}

// SetState sets the "state" field.
func (ghiuo *GitHubIssueUpdateOne) SetState(gi githubissue.State) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetState(gi)
	return ghiuo
}

// SetLocked sets the "locked" field.
func (ghiuo *GitHubIssueUpdateOne) SetLocked(b bool) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetLocked(b)
	return ghiuo
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (ghiuo *GitHubIssueUpdateOne) SetNillableLocked(b *bool) *GitHubIssueUpdateOne {
	if b != nil {
		ghiuo.SetLocked(*b)
	}
	return ghiuo
}

// SetActiveLockReason sets the "active_lock_reason" field.
func (ghiuo *GitHubIssueUpdateOne) SetActiveLockReason(s string) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetActiveLockReason(s)
	return ghiuo
}

// SetCommentsCount sets the "comments_count" field.
func (ghiuo *GitHubIssueUpdateOne) SetCommentsCount(i int) *GitHubIssueUpdateOne {
	ghiuo.mutation.ResetCommentsCount()
	ghiuo.mutation.SetCommentsCount(i)
	return ghiuo
}

// AddCommentsCount adds i to the "comments_count" field.
func (ghiuo *GitHubIssueUpdateOne) AddCommentsCount(i int) *GitHubIssueUpdateOne {
	ghiuo.mutation.AddCommentsCount(i)
	return ghiuo
}

// SetCreatedAt sets the "created_at" field.
func (ghiuo *GitHubIssueUpdateOne) SetCreatedAt(s string) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetCreatedAt(s)
	return ghiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ghiuo *GitHubIssueUpdateOne) SetUpdatedAt(s string) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetUpdatedAt(s)
	return ghiuo
}

// SetClosedAt sets the "closed_at" field.
func (ghiuo *GitHubIssueUpdateOne) SetClosedAt(s string) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetClosedAt(s)
	return ghiuo
}

// SetAuthorAssociation sets the "author_association" field.
func (ghiuo *GitHubIssueUpdateOne) SetAuthorAssociation(s string) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetAuthorAssociation(s)
	return ghiuo
}

// AddAssigneeIDs adds the "assignees" edge to the GitHubUser entity by IDs.
func (ghiuo *GitHubIssueUpdateOne) AddAssigneeIDs(ids ...int) *GitHubIssueUpdateOne {
	ghiuo.mutation.AddAssigneeIDs(ids...)
	return ghiuo
}

// AddAssignees adds the "assignees" edges to the GitHubUser entity.
func (ghiuo *GitHubIssueUpdateOne) AddAssignees(g ...*GitHubUser) *GitHubIssueUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghiuo.AddAssigneeIDs(ids...)
}

// SetAuthorID sets the "author" edge to the GitHubUser entity by ID.
func (ghiuo *GitHubIssueUpdateOne) SetAuthorID(id int) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetAuthorID(id)
	return ghiuo
}

// SetAuthor sets the "author" edge to the GitHubUser entity.
func (ghiuo *GitHubIssueUpdateOne) SetAuthor(g *GitHubUser) *GitHubIssueUpdateOne {
	return ghiuo.SetAuthorID(g.ID)
}

// SetClosedByID sets the "closed_by" edge to the GitHubUser entity by ID.
func (ghiuo *GitHubIssueUpdateOne) SetClosedByID(id int) *GitHubIssueUpdateOne {
	ghiuo.mutation.SetClosedByID(id)
	return ghiuo
}

// SetNillableClosedByID sets the "closed_by" edge to the GitHubUser entity by ID if the given value is not nil.
func (ghiuo *GitHubIssueUpdateOne) SetNillableClosedByID(id *int) *GitHubIssueUpdateOne {
	if id != nil {
		ghiuo = ghiuo.SetClosedByID(*id)
	}
	return ghiuo
}

// SetClosedBy sets the "closed_by" edge to the GitHubUser entity.
func (ghiuo *GitHubIssueUpdateOne) SetClosedBy(g *GitHubUser) *GitHubIssueUpdateOne {
	return ghiuo.SetClosedByID(g.ID)
}

// Mutation returns the GitHubIssueMutation object of the builder.
func (ghiuo *GitHubIssueUpdateOne) Mutation() *GitHubIssueMutation {
	return ghiuo.mutation
}

// ClearAssignees clears all "assignees" edges to the GitHubUser entity.
func (ghiuo *GitHubIssueUpdateOne) ClearAssignees() *GitHubIssueUpdateOne {
	ghiuo.mutation.ClearAssignees()
	return ghiuo
}

// RemoveAssigneeIDs removes the "assignees" edge to GitHubUser entities by IDs.
func (ghiuo *GitHubIssueUpdateOne) RemoveAssigneeIDs(ids ...int) *GitHubIssueUpdateOne {
	ghiuo.mutation.RemoveAssigneeIDs(ids...)
	return ghiuo
}

// RemoveAssignees removes "assignees" edges to GitHubUser entities.
func (ghiuo *GitHubIssueUpdateOne) RemoveAssignees(g ...*GitHubUser) *GitHubIssueUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghiuo.RemoveAssigneeIDs(ids...)
}

// ClearAuthor clears the "author" edge to the GitHubUser entity.
func (ghiuo *GitHubIssueUpdateOne) ClearAuthor() *GitHubIssueUpdateOne {
	ghiuo.mutation.ClearAuthor()
	return ghiuo
}

// ClearClosedBy clears the "closed_by" edge to the GitHubUser entity.
func (ghiuo *GitHubIssueUpdateOne) ClearClosedBy() *GitHubIssueUpdateOne {
	ghiuo.mutation.ClearClosedBy()
	return ghiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ghiuo *GitHubIssueUpdateOne) Select(field string, fields ...string) *GitHubIssueUpdateOne {
	ghiuo.fields = append([]string{field}, fields...)
	return ghiuo
}

// Save executes the query and returns the updated GitHubIssue entity.
func (ghiuo *GitHubIssueUpdateOne) Save(ctx context.Context) (*GitHubIssue, error) {
	var (
		err  error
		node *GitHubIssue
	)
	if len(ghiuo.hooks) == 0 {
		if err = ghiuo.check(); err != nil {
			return nil, err
		}
		node, err = ghiuo.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GitHubIssueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ghiuo.check(); err != nil {
				return nil, err
			}
			ghiuo.mutation = mutation
			node, err = ghiuo.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ghiuo.hooks) - 1; i >= 0; i-- {
			if ghiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ghiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ghiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ghiuo *GitHubIssueUpdateOne) SaveX(ctx context.Context) *GitHubIssue {
	node, err := ghiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ghiuo *GitHubIssueUpdateOne) Exec(ctx context.Context) error {
	_, err := ghiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ghiuo *GitHubIssueUpdateOne) ExecX(ctx context.Context) {
	if err := ghiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ghiuo *GitHubIssueUpdateOne) check() error {
	if v, ok := ghiuo.mutation.GithubID(); ok {
		if err := githubissue.GithubIDValidator(v); err != nil {
			return &ValidationError{Name: "github_id", err: fmt.Errorf("ent: validator failed for field \"github_id\": %w", err)}
		}
	}
	if v, ok := ghiuo.mutation.Number(); ok {
		if err := githubissue.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf("ent: validator failed for field \"number\": %w", err)}
		}
	}
	if v, ok := ghiuo.mutation.Title(); ok {
		if err := githubissue.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := ghiuo.mutation.State(); ok {
		if err := githubissue.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	if v, ok := ghiuo.mutation.CommentsCount(); ok {
		if err := githubissue.CommentsCountValidator(v); err != nil {
			return &ValidationError{Name: "comments_count", err: fmt.Errorf("ent: validator failed for field \"comments_count\": %w", err)}
		}
	}
	if v, ok := ghiuo.mutation.CreatedAt(); ok {
		if err := githubissue.CreatedAtValidator(v); err != nil {
			return &ValidationError{Name: "created_at", err: fmt.Errorf("ent: validator failed for field \"created_at\": %w", err)}
		}
	}
	if v, ok := ghiuo.mutation.UpdatedAt(); ok {
		if err := githubissue.UpdatedAtValidator(v); err != nil {
			return &ValidationError{Name: "updated_at", err: fmt.Errorf("ent: validator failed for field \"updated_at\": %w", err)}
		}
	}
	if v, ok := ghiuo.mutation.ClosedAt(); ok {
		if err := githubissue.ClosedAtValidator(v); err != nil {
			return &ValidationError{Name: "closed_at", err: fmt.Errorf("ent: validator failed for field \"closed_at\": %w", err)}
		}
	}
	if _, ok := ghiuo.mutation.AuthorID(); ghiuo.mutation.AuthorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"author\"")
	}
	return nil
}

func (ghiuo *GitHubIssueUpdateOne) gremlinSave(ctx context.Context) (*GitHubIssue, error) {
	res := &gremlin.Response{}
	id, ok := ghiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GitHubIssue.ID for update")}
	}
	query, bindings := ghiuo.gremlin(id).Query()
	if err := ghiuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ghi := &GitHubIssue{config: ghiuo.config}
	if err := ghi.FromResponse(res); err != nil {
		return nil, err
	}
	return ghi, nil
}

func (ghiuo *GitHubIssueUpdateOne) gremlin(id int) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := ghiuo.mutation.GithubID(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(githubissue.Label, githubissue.FieldGithubID, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(githubissue.Label, githubissue.FieldGithubID, value)),
		})
		v.Property(dsl.Single, githubissue.FieldGithubID, value)
	}
	if value, ok := ghiuo.mutation.AddedGithubID(); ok {
		addValue := rv.Clone().Union(__.Values(githubissue.FieldGithubID), __.Constant(value)).Sum().Next()
		constraints = append(constraints, &constraint{
			pred: g.V().Has(githubissue.Label, githubissue.FieldGithubID, addValue).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(githubissue.Label, githubissue.FieldGithubID, fmt.Sprintf("+= %v", value))),
		})
		v.Property(dsl.Single, githubissue.FieldGithubID, __.Union(__.Values(githubissue.FieldGithubID), __.Constant(value)).Sum())
	}
	if value, ok := ghiuo.mutation.Number(); ok {
		v.Property(dsl.Single, githubissue.FieldNumber, value)
	}
	if value, ok := ghiuo.mutation.AddedNumber(); ok {
		v.Property(dsl.Single, githubissue.FieldNumber, __.Union(__.Values(githubissue.FieldNumber), __.Constant(value)).Sum())
	}
	if value, ok := ghiuo.mutation.Title(); ok {
		v.Property(dsl.Single, githubissue.FieldTitle, value)
	}
	if value, ok := ghiuo.mutation.Body(); ok {
		v.Property(dsl.Single, githubissue.FieldBody, value)
	}
	if value, ok := ghiuo.mutation.HTMLURL(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(githubissue.Label, githubissue.FieldHTMLURL, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(githubissue.Label, githubissue.FieldHTMLURL, value)),
		})
		v.Property(dsl.Single, githubissue.FieldHTMLURL, value)
	}
	if value, ok := ghiuo.mutation.State(); ok {
		v.Property(dsl.Single, githubissue.FieldState, value)
	}
	if value, ok := ghiuo.mutation.Locked(); ok {
		v.Property(dsl.Single, githubissue.FieldLocked, value)
	}
	if value, ok := ghiuo.mutation.ActiveLockReason(); ok {
		v.Property(dsl.Single, githubissue.FieldActiveLockReason, value)
	}
	if value, ok := ghiuo.mutation.CommentsCount(); ok {
		v.Property(dsl.Single, githubissue.FieldCommentsCount, value)
	}
	if value, ok := ghiuo.mutation.AddedCommentsCount(); ok {
		v.Property(dsl.Single, githubissue.FieldCommentsCount, __.Union(__.Values(githubissue.FieldCommentsCount), __.Constant(value)).Sum())
	}
	if value, ok := ghiuo.mutation.CreatedAt(); ok {
		v.Property(dsl.Single, githubissue.FieldCreatedAt, value)
	}
	if value, ok := ghiuo.mutation.UpdatedAt(); ok {
		v.Property(dsl.Single, githubissue.FieldUpdatedAt, value)
	}
	if value, ok := ghiuo.mutation.ClosedAt(); ok {
		v.Property(dsl.Single, githubissue.FieldClosedAt, value)
	}
	if value, ok := ghiuo.mutation.AuthorAssociation(); ok {
		v.Property(dsl.Single, githubissue.FieldAuthorAssociation, value)
	}
	for _, id := range ghiuo.mutation.RemovedAssigneesIDs() {
		tr := rv.Clone().OutE(githubissue.AssigneesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range ghiuo.mutation.AssigneesIDs() {
		v.AddE(githubissue.AssigneesLabel).To(g.V(id)).OutV()
	}
	if ghiuo.mutation.AuthorCleared() {
		tr := rv.Clone().InE(githubuser.CreatedIssuesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range ghiuo.mutation.AuthorIDs() {
		v.AddE(githubuser.CreatedIssuesLabel).From(g.V(id)).InV()
	}
	if ghiuo.mutation.ClosedByCleared() {
		tr := rv.Clone().InE(githubuser.ClosedIssuesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range ghiuo.mutation.ClosedByIDs() {
		v.AddE(githubuser.ClosedIssuesLabel).From(g.V(id)).InV()
	}
	if len(ghiuo.fields) > 0 {
		fields := make([]interface{}, 0, len(ghiuo.fields)+1)
		fields = append(fields, true)
		for _, f := range ghiuo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

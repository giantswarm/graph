// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"github.com/giantswarm/graph/ent/githubissue"
	"github.com/giantswarm/graph/ent/githubuser"
	"github.com/giantswarm/graph/ent/predicate"
)

// GitHubIssueQuery is the builder for querying GitHubIssue entities.
type GitHubIssueQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GitHubIssue
	// eager-loading edges.
	withAssignees *GitHubUserQuery
	withAuthor    *GitHubUserQuery
	withClosedBy  *GitHubUserQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the GitHubIssueQuery builder.
func (ghiq *GitHubIssueQuery) Where(ps ...predicate.GitHubIssue) *GitHubIssueQuery {
	ghiq.predicates = append(ghiq.predicates, ps...)
	return ghiq
}

// Limit adds a limit step to the query.
func (ghiq *GitHubIssueQuery) Limit(limit int) *GitHubIssueQuery {
	ghiq.limit = &limit
	return ghiq
}

// Offset adds an offset step to the query.
func (ghiq *GitHubIssueQuery) Offset(offset int) *GitHubIssueQuery {
	ghiq.offset = &offset
	return ghiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ghiq *GitHubIssueQuery) Unique(unique bool) *GitHubIssueQuery {
	ghiq.unique = &unique
	return ghiq
}

// Order adds an order step to the query.
func (ghiq *GitHubIssueQuery) Order(o ...OrderFunc) *GitHubIssueQuery {
	ghiq.order = append(ghiq.order, o...)
	return ghiq
}

// QueryAssignees chains the current query on the "assignees" edge.
func (ghiq *GitHubIssueQuery) QueryAssignees() *GitHubUserQuery {
	query := &GitHubUserQuery{config: ghiq.config}
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := ghiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := ghiq.gremlinQuery(ctx)
		fromU = gremlin.OutE(githubissue.AssigneesLabel).InV()
		return fromU, nil
	}
	return query
}

// QueryAuthor chains the current query on the "author" edge.
func (ghiq *GitHubIssueQuery) QueryAuthor() *GitHubUserQuery {
	query := &GitHubUserQuery{config: ghiq.config}
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := ghiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := ghiq.gremlinQuery(ctx)
		fromU = gremlin.InE(githubuser.CreatedIssuesLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryClosedBy chains the current query on the "closed_by" edge.
func (ghiq *GitHubIssueQuery) QueryClosedBy() *GitHubUserQuery {
	query := &GitHubUserQuery{config: ghiq.config}
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := ghiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := ghiq.gremlinQuery(ctx)
		fromU = gremlin.InE(githubuser.ClosedIssuesLabel).OutV()
		return fromU, nil
	}
	return query
}

// First returns the first GitHubIssue entity from the query.
// Returns a *NotFoundError when no GitHubIssue was found.
func (ghiq *GitHubIssueQuery) First(ctx context.Context) (*GitHubIssue, error) {
	nodes, err := ghiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{githubissue.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ghiq *GitHubIssueQuery) FirstX(ctx context.Context) *GitHubIssue {
	node, err := ghiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GitHubIssue ID from the query.
// Returns a *NotFoundError when no GitHubIssue ID was found.
func (ghiq *GitHubIssueQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ghiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{githubissue.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ghiq *GitHubIssueQuery) FirstIDX(ctx context.Context) int {
	id, err := ghiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GitHubIssue entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GitHubIssue entity is not found.
// Returns a *NotFoundError when no GitHubIssue entities are found.
func (ghiq *GitHubIssueQuery) Only(ctx context.Context) (*GitHubIssue, error) {
	nodes, err := ghiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{githubissue.Label}
	default:
		return nil, &NotSingularError{githubissue.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ghiq *GitHubIssueQuery) OnlyX(ctx context.Context) *GitHubIssue {
	node, err := ghiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GitHubIssue ID in the query.
// Returns a *NotSingularError when exactly one GitHubIssue ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ghiq *GitHubIssueQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ghiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{githubissue.Label}
	default:
		err = &NotSingularError{githubissue.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ghiq *GitHubIssueQuery) OnlyIDX(ctx context.Context) int {
	id, err := ghiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GitHubIssues.
func (ghiq *GitHubIssueQuery) All(ctx context.Context) ([]*GitHubIssue, error) {
	if err := ghiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ghiq.gremlinAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ghiq *GitHubIssueQuery) AllX(ctx context.Context) []*GitHubIssue {
	nodes, err := ghiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GitHubIssue IDs.
func (ghiq *GitHubIssueQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ghiq.Select(githubissue.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ghiq *GitHubIssueQuery) IDsX(ctx context.Context) []int {
	ids, err := ghiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ghiq *GitHubIssueQuery) Count(ctx context.Context) (int, error) {
	if err := ghiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ghiq.gremlinCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ghiq *GitHubIssueQuery) CountX(ctx context.Context) int {
	count, err := ghiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ghiq *GitHubIssueQuery) Exist(ctx context.Context) (bool, error) {
	if err := ghiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ghiq.gremlinExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ghiq *GitHubIssueQuery) ExistX(ctx context.Context) bool {
	exist, err := ghiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GitHubIssueQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ghiq *GitHubIssueQuery) Clone() *GitHubIssueQuery {
	if ghiq == nil {
		return nil
	}
	return &GitHubIssueQuery{
		config:        ghiq.config,
		limit:         ghiq.limit,
		offset:        ghiq.offset,
		order:         append([]OrderFunc{}, ghiq.order...),
		predicates:    append([]predicate.GitHubIssue{}, ghiq.predicates...),
		withAssignees: ghiq.withAssignees.Clone(),
		withAuthor:    ghiq.withAuthor.Clone(),
		withClosedBy:  ghiq.withClosedBy.Clone(),
		// clone intermediate query.
		gremlin: ghiq.gremlin.Clone(),
		path:    ghiq.path,
	}
}

// WithAssignees tells the query-builder to eager-load the nodes that are connected to
// the "assignees" edge. The optional arguments are used to configure the query builder of the edge.
func (ghiq *GitHubIssueQuery) WithAssignees(opts ...func(*GitHubUserQuery)) *GitHubIssueQuery {
	query := &GitHubUserQuery{config: ghiq.config}
	for _, opt := range opts {
		opt(query)
	}
	ghiq.withAssignees = query
	return ghiq
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (ghiq *GitHubIssueQuery) WithAuthor(opts ...func(*GitHubUserQuery)) *GitHubIssueQuery {
	query := &GitHubUserQuery{config: ghiq.config}
	for _, opt := range opts {
		opt(query)
	}
	ghiq.withAuthor = query
	return ghiq
}

// WithClosedBy tells the query-builder to eager-load the nodes that are connected to
// the "closed_by" edge. The optional arguments are used to configure the query builder of the edge.
func (ghiq *GitHubIssueQuery) WithClosedBy(opts ...func(*GitHubUserQuery)) *GitHubIssueQuery {
	query := &GitHubUserQuery{config: ghiq.config}
	for _, opt := range opts {
		opt(query)
	}
	ghiq.withClosedBy = query
	return ghiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GithubID int `json:"github_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GitHubIssue.Query().
//		GroupBy(githubissue.FieldGithubID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ghiq *GitHubIssueQuery) GroupBy(field string, fields ...string) *GitHubIssueGroupBy {
	group := &GitHubIssueGroupBy{config: ghiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *dsl.Traversal, err error) {
		if err := ghiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ghiq.gremlinQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GithubID int `json:"github_id,omitempty"`
//	}
//
//	client.GitHubIssue.Query().
//		Select(githubissue.FieldGithubID).
//		Scan(ctx, &v)
//
func (ghiq *GitHubIssueQuery) Select(fields ...string) *GitHubIssueSelect {
	ghiq.fields = append(ghiq.fields, fields...)
	return &GitHubIssueSelect{GitHubIssueQuery: ghiq}
}

func (ghiq *GitHubIssueQuery) prepareQuery(ctx context.Context) error {
	if ghiq.path != nil {
		prev, err := ghiq.path(ctx)
		if err != nil {
			return err
		}
		ghiq.gremlin = prev
	}
	return nil
}

func (ghiq *GitHubIssueQuery) gremlinAll(ctx context.Context) ([]*GitHubIssue, error) {
	res := &gremlin.Response{}
	traversal := ghiq.gremlinQuery(ctx)
	if len(ghiq.fields) > 0 {
		fields := make([]interface{}, len(ghiq.fields))
		for i, f := range ghiq.fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := ghiq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var ghis GitHubIssues
	if err := ghis.FromResponse(res); err != nil {
		return nil, err
	}
	ghis.config(ghiq.config)
	return ghis, nil
}

func (ghiq *GitHubIssueQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ghiq.gremlinQuery(ctx).Count().Query()
	if err := ghiq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (ghiq *GitHubIssueQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := ghiq.gremlinQuery(ctx).HasNext().Query()
	if err := ghiq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (ghiq *GitHubIssueQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(githubissue.Label)
	if ghiq.gremlin != nil {
		v = ghiq.gremlin.Clone()
	}
	for _, p := range ghiq.predicates {
		p(v)
	}
	if len(ghiq.order) > 0 {
		v.Order()
		for _, p := range ghiq.order {
			p(v)
		}
	}
	switch limit, offset := ghiq.limit, ghiq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := ghiq.unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// GitHubIssueGroupBy is the group-by builder for GitHubIssue entities.
type GitHubIssueGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ghigb *GitHubIssueGroupBy) Aggregate(fns ...AggregateFunc) *GitHubIssueGroupBy {
	ghigb.fns = append(ghigb.fns, fns...)
	return ghigb
}

// Scan applies the group-by query and scans the result into the given value.
func (ghigb *GitHubIssueGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ghigb.path(ctx)
	if err != nil {
		return err
	}
	ghigb.gremlin = query
	return ghigb.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ghigb *GitHubIssueGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ghigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ghigb *GitHubIssueGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ghigb.fields) > 1 {
		return nil, errors.New("ent: GitHubIssueGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ghigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ghigb *GitHubIssueGroupBy) StringsX(ctx context.Context) []string {
	v, err := ghigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ghigb *GitHubIssueGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ghigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githubissue.Label}
	default:
		err = fmt.Errorf("ent: GitHubIssueGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ghigb *GitHubIssueGroupBy) StringX(ctx context.Context) string {
	v, err := ghigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ghigb *GitHubIssueGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ghigb.fields) > 1 {
		return nil, errors.New("ent: GitHubIssueGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ghigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ghigb *GitHubIssueGroupBy) IntsX(ctx context.Context) []int {
	v, err := ghigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ghigb *GitHubIssueGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ghigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githubissue.Label}
	default:
		err = fmt.Errorf("ent: GitHubIssueGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ghigb *GitHubIssueGroupBy) IntX(ctx context.Context) int {
	v, err := ghigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ghigb *GitHubIssueGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ghigb.fields) > 1 {
		return nil, errors.New("ent: GitHubIssueGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ghigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ghigb *GitHubIssueGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ghigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ghigb *GitHubIssueGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ghigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githubissue.Label}
	default:
		err = fmt.Errorf("ent: GitHubIssueGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ghigb *GitHubIssueGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ghigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ghigb *GitHubIssueGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ghigb.fields) > 1 {
		return nil, errors.New("ent: GitHubIssueGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ghigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ghigb *GitHubIssueGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ghigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ghigb *GitHubIssueGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ghigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githubissue.Label}
	default:
		err = fmt.Errorf("ent: GitHubIssueGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ghigb *GitHubIssueGroupBy) BoolX(ctx context.Context) bool {
	v, err := ghigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ghigb *GitHubIssueGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := ghigb.gremlinQuery().Query()
	if err := ghigb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ghigb.fields)+len(ghigb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (ghigb *GitHubIssueGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range ghigb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range ghigb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return ghigb.gremlin.Group().
		By(__.Values(ghigb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}

// GitHubIssueSelect is the builder for selecting fields of GitHubIssue entities.
type GitHubIssueSelect struct {
	*GitHubIssueQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
}

// Scan applies the selector query and scans the result into the given value.
func (ghis *GitHubIssueSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ghis.prepareQuery(ctx); err != nil {
		return err
	}
	ghis.gremlin = ghis.GitHubIssueQuery.gremlinQuery(ctx)
	return ghis.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ghis *GitHubIssueSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ghis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ghis *GitHubIssueSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ghis.fields) > 1 {
		return nil, errors.New("ent: GitHubIssueSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ghis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ghis *GitHubIssueSelect) StringsX(ctx context.Context) []string {
	v, err := ghis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ghis *GitHubIssueSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ghis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githubissue.Label}
	default:
		err = fmt.Errorf("ent: GitHubIssueSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ghis *GitHubIssueSelect) StringX(ctx context.Context) string {
	v, err := ghis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ghis *GitHubIssueSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ghis.fields) > 1 {
		return nil, errors.New("ent: GitHubIssueSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ghis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ghis *GitHubIssueSelect) IntsX(ctx context.Context) []int {
	v, err := ghis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ghis *GitHubIssueSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ghis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githubissue.Label}
	default:
		err = fmt.Errorf("ent: GitHubIssueSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ghis *GitHubIssueSelect) IntX(ctx context.Context) int {
	v, err := ghis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ghis *GitHubIssueSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ghis.fields) > 1 {
		return nil, errors.New("ent: GitHubIssueSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ghis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ghis *GitHubIssueSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ghis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ghis *GitHubIssueSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ghis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githubissue.Label}
	default:
		err = fmt.Errorf("ent: GitHubIssueSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ghis *GitHubIssueSelect) Float64X(ctx context.Context) float64 {
	v, err := ghis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ghis *GitHubIssueSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ghis.fields) > 1 {
		return nil, errors.New("ent: GitHubIssueSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ghis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ghis *GitHubIssueSelect) BoolsX(ctx context.Context) []bool {
	v, err := ghis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ghis *GitHubIssueSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ghis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githubissue.Label}
	default:
		err = fmt.Errorf("ent: GitHubIssueSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ghis *GitHubIssueSelect) BoolX(ctx context.Context) bool {
	v, err := ghis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ghis *GitHubIssueSelect) gremlinScan(ctx context.Context, v interface{}) error {
	var (
		traversal *dsl.Traversal
		res       = &gremlin.Response{}
	)
	if len(ghis.fields) == 1 {
		if ghis.fields[0] != githubissue.FieldID {
			traversal = ghis.gremlin.Values(ghis.fields...)
		} else {
			traversal = ghis.gremlin.ID()
		}
	} else {
		fields := make([]interface{}, len(ghis.fields))
		for i, f := range ghis.fields {
			fields[i] = f
		}
		traversal = ghis.gremlin.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := ghis.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ghis.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

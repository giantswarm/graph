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
	"github.com/giantswarm/graph/ent/person"
)

// PersonCreate is the builder for creating a Person entity.
type PersonCreate struct {
	config
	mutation *PersonMutation
	hooks    []Hook
}

// SetEmail sets the "email" field.
func (pc *PersonCreate) SetEmail(s string) *PersonCreate {
	pc.mutation.SetEmail(s)
	return pc
}

// SetName sets the "name" field.
func (pc *PersonCreate) SetName(s string) *PersonCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetIsGiantSwarmEmployee sets the "isGiantSwarmEmployee" field.
func (pc *PersonCreate) SetIsGiantSwarmEmployee(b bool) *PersonCreate {
	pc.mutation.SetIsGiantSwarmEmployee(b)
	return pc
}

// SetGithubAccountID sets the "github_account" edge to the GitHubUser entity by ID.
func (pc *PersonCreate) SetGithubAccountID(id int) *PersonCreate {
	pc.mutation.SetGithubAccountID(id)
	return pc
}

// SetNillableGithubAccountID sets the "github_account" edge to the GitHubUser entity by ID if the given value is not nil.
func (pc *PersonCreate) SetNillableGithubAccountID(id *int) *PersonCreate {
	if id != nil {
		pc = pc.SetGithubAccountID(*id)
	}
	return pc
}

// SetGithubAccount sets the "github_account" edge to the GitHubUser entity.
func (pc *PersonCreate) SetGithubAccount(g *GitHubUser) *PersonCreate {
	return pc.SetGithubAccountID(g.ID)
}

// Mutation returns the PersonMutation object of the builder.
func (pc *PersonCreate) Mutation() *PersonMutation {
	return pc.mutation
}

// Save creates the Person in the database.
func (pc *PersonCreate) Save(ctx context.Context) (*Person, error) {
	var (
		err  error
		node *Person
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.gremlinSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PersonCreate) SaveX(ctx context.Context) *Person {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PersonCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PersonCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PersonCreate) check() error {
	if _, ok := pc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "email"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := pc.mutation.IsGiantSwarmEmployee(); !ok {
		return &ValidationError{Name: "isGiantSwarmEmployee", err: errors.New(`ent: missing required field "isGiantSwarmEmployee"`)}
	}
	return nil
}

func (pc *PersonCreate) gremlinSave(ctx context.Context) (*Person, error) {
	res := &gremlin.Response{}
	query, bindings := pc.gremlin().Query()
	if err := pc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	pe := &Person{config: pc.config}
	if err := pe.FromResponse(res); err != nil {
		return nil, err
	}
	return pe, nil
}

func (pc *PersonCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		firstQueryPred *dsl.Traversal // constraint predicate.
		pred           *dsl.Traversal // constraint predicate.
		test           *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	createTraversal := func(constraints []*constraint, traversalFuncs []func(*dsl.Traversal)) *dsl.Traversal {
		var v *dsl.Traversal
		if len(constraints) > 0 {
			// We will use coalesce, therefore AddV will be child traversal, so we need __ here
			v = __.New().AddV(person.Label)
		} else {
			v = g.AddV(person.Label)
		}
		for _, tf := range traversalFuncs {
			tf(v)
		}
		return v
	}
	traversalFuncs := []func(*dsl.Traversal){}
	if value, ok := pc.mutation.Email(); ok {
		constraints = append(constraints, &constraint{
			firstQueryPred: g.V().Has(person.Label, person.FieldEmail, value).Count(),
			pred:           __.V().Has(person.Label, person.FieldEmail, value).Count(),
			test:           __.Is(p.NEQ(0)).Constant(NewErrUniqueField(person.Label, person.FieldEmail, value)),
		})
		traversalFuncs = append(traversalFuncs, func(v *dsl.Traversal) {
			v.Property(dsl.Single, person.FieldEmail, value)
		})
	}
	if value, ok := pc.mutation.Name(); ok {
		traversalFuncs = append(traversalFuncs, func(v *dsl.Traversal) {
			v.Property(dsl.Single, person.FieldName, value)
		})
	}
	if value, ok := pc.mutation.IsGiantSwarmEmployee(); ok {
		traversalFuncs = append(traversalFuncs, func(v *dsl.Traversal) {
			v.Property(dsl.Single, person.FieldIsGiantSwarmEmployee, value)
		})
	}
	for _, id := range pc.mutation.GithubAccountIDs() {
		traversalFuncs = append(traversalFuncs, func(v *dsl.Traversal) {
			v.AddE(person.GithubAccountLabel).To(g.V(id)).OutV()
		})
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(person.GithubAccountLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(person.Label, person.GithubAccountLabel, string(id))),
		})
	}
	v := createTraversal(constraints, traversalFuncs)
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	var tr *dsl.Traversal
	if len(constraints) == 1 {
		// use the TraversalSource (g) to start the traversal
		tr = constraints[0].firstQueryPred.Coalesce(constraints[0].test, v.ValueMap(true))
	} else {
		// use the __ class rather than a TraversalSource to construct the child traversal anonymously
		tr = constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	}
	for i, cr := range constraints[1:] {
		if i == len(constraints[1:])-1 && cr.firstQueryPred != nil {
			// use the TraversalSource (g) to start the traversal
			tr = cr.firstQueryPred.Coalesce(cr.test, tr)
		} else {
			// use the __ class rather than a TraversalSource to construct the child traversal anonymously
			tr = cr.pred.Coalesce(cr.test, tr)
		}
	}
	return tr
}

// PersonCreateBulk is the builder for creating many Person entities in bulk.
type PersonCreateBulk struct {
	config
	builders []*PersonCreate
}
